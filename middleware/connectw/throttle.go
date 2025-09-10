package connectw

import (
	"context"
	"errors"
	"strconv"
	"time"

	"connectrpc.com/connect"
	"github.com/redis/go-redis/v9"
)

// Throttle represents a sliding window rate limiter for connectrpc.
type Throttle struct {
	redisClient  *redis.Client
	limit        int64
	window       time.Duration
	keyPrefix    string
	KeyExtractor func(ctx context.Context) string // Now takes http.Request
}

// NewThrottle creates a new Throttle interceptor for connectrpc.
func NewThrottle(redisClient *redis.Client, limit int64, window time.Duration, keyPrefix string,
	keyExtractor func(ctx context.Context) string,
) *Throttle {
	return &Throttle{
		redisClient:  redisClient,
		limit:        limit,
		window:       window,
		keyPrefix:    keyPrefix,
		KeyExtractor: keyExtractor,
	}
}

// UnaryInterceptor returns a connect.UnaryInterceptorFunc that applies rate limiting.
func (t *Throttle) UnaryInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			key := t.keyPrefix + ":" + t.KeyExtractor(ctx)

			now := time.Now().UnixNano()
			minScore := now - t.window.Nanoseconds()

			pipe := t.redisClient.Pipeline()
			pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(minScore, 10))
			pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
			pipe.Expire(ctx, key, t.window) // Set expiration for the key
			_, err := pipe.Exec(ctx)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, err)
			}

			count, err := t.redisClient.ZCard(ctx, key).Result()
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, err)
			}

			if count > t.limit {
				return nil, connect.NewError(connect.CodeResourceExhausted, errors.New("too many requests"))
			}

			return next(ctx, req)
		})
	}
}
