package connectw

import (
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// Throttle represents a sliding window rate limiter.
type Throttle struct {
	redisClient  *redis.Client
	limit        int64
	window       time.Duration
	keyPrefix    string
	KeyExtractor func(r *http.Request) string
}

// NewThrottle creates a new Throttle middleware.
func NewThrottle(redisClient *redis.Client, limit int64, window time.Duration, keyPrefix string, keyExtractor func(r *http.Request) string) *Throttle {
	return &Throttle{
		redisClient:  redisClient,
		limit:        limit,
		window:       window,
		keyPrefix:    keyPrefix,
		KeyExtractor: keyExtractor,
	}
}

// Handle is the middleware function for connectw.
func (t *Throttle) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		key := t.keyPrefix + ":" + t.KeyExtractor(r) // Use the extractor

		now := time.Now().UnixNano()
		minScore := now - t.window.Nanoseconds()

		// Remove old requests and add current request in a single transaction
		pipe := t.redisClient.Pipeline()
		pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(minScore, 10))
		pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
		pipe.Expire(ctx, key, t.window) // Set expiration for the key
		_, err := pipe.Exec(ctx)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Get current count
		count, err := t.redisClient.ZCard(ctx, key).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if count > t.limit {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too Many Requests"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
