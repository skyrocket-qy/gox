package ginw

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Throttle represents a sliding window rate limiter.
type Throttle struct {
	redisClient  *redis.Client
	limit        int64
	window       time.Duration
	keyPrefix    string
	KeyExtractor func(c *gin.Context) string
}

// NewThrottle creates a new Throttle middleware.
func NewThrottle(redisClient *redis.Client, limit int64, window time.Duration, keyPrefix string,
	keyExtractor func(c *gin.Context) string,
) *Throttle {
	return &Throttle{
		redisClient:  redisClient,
		limit:        limit,
		window:       window,
		keyPrefix:    keyPrefix,
		KeyExtractor: keyExtractor,
	}
}

// Middleware is the middleware function for ginw.
func (t *Throttle) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		key := t.keyPrefix + ":" + t.KeyExtractor(c)

		now := time.Now().UnixNano()
		minScore := now - t.window.Nanoseconds()

		// Remove old requests and add current request in a single transaction
		pipe := t.redisClient.Pipeline()
		pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(minScore, 10))
		pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
		pipe.Expire(ctx, key, t.window) // Set expiration for the key
		_, err := pipe.Exec(ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Get current count
		count, err := t.redisClient.ZCard(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if count > t.limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next()
	}
}
