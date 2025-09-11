package connectw

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"
	"github.com/redis/go-redis/v9"
	"github.com/skyrocket-qy/gox/redisx/slidewindow"
)

// Clock is an interface for getting the current time.
type Clock interface {
	Now() time.Time
}

// realClock implements the Clock interface using the real time.
type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

// Throttle represents a sliding window rate limiter for connectrpc.
type Throttle struct {
	limiter      slidewindow.MovingWindowLimiterInterface
	limit        int64
	window       time.Duration
	keyPrefix    string
	KeyExtractor func(ctx context.Context) string
	clock        Clock
}

// NewThrottle creates a new Throttle interceptor for connectrpc.
func NewThrottle(redisClient *redis.Client, limit int64, window time.Duration, keyPrefix string,
	keyExtractor func(ctx context.Context) string,
	limiter slidewindow.MovingWindowLimiterInterface,
) *Throttle {
	if limiter == nil {
		limiter = slidewindow.New(redisClient)
	}
	return &Throttle{
		limiter:      limiter,
		limit:        limit,
		window:       window,
		keyPrefix:    keyPrefix,
		KeyExtractor: keyExtractor,
		clock:        realClock{},
	}
}

// UnaryInterceptor returns a connect.UnaryInterceptorFunc that applies rate limiting.
func (t *Throttle) UnaryInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				key := t.keyPrefix + ":" + t.KeyExtractor(ctx)

				allowed, err := t.limiter.Allow(ctx, key, t.limit, t.window)
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, err)
				}

				if !allowed {
					return nil, connect.NewError(
						connect.CodeResourceExhausted,
						errors.New("too many requests"),
					)
				}

				return next(ctx, req)
			},
		)
	}
}
