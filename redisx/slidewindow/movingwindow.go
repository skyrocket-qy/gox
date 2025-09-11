package slidewindow

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
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

// MovingWindowLimiterInterface defines the interface for a moving window rate limiter.
type MovingWindowLimiterInterface interface {
	Allow(ctx context.Context, key string, limit int64, window time.Duration) (bool, error)
}

// MovingWindowLimiter provides a general sliding window rate limiter using Redis.
type MovingWindowLimiter struct {
	redisClient *redis.Client
	clock       Clock
}

// NewMovingWindowLimiter creates a new MovingWindowLimiter.
func New(redisClient *redis.Client) *MovingWindowLimiter {
	return &MovingWindowLimiter{
		redisClient: redisClient,
		clock:       realClock{},
	}
}

// SetClock sets the clock for the MovingWindowLimiter. Useful for testing.
func (l *MovingWindowLimiter) SetClock(clock Clock) {
	l.clock = clock
}

// Allow checks if a request is allowed based on the sliding window algorithm.
// It returns true if the request is allowed, false otherwise, and an error if any Redis operation fails.
func (l *MovingWindowLimiter) Allow(ctx context.Context, key string, limit int64, window time.Duration) (bool, error) {
	now := l.clock.Now().UnixNano()
	minScore := now - window.Nanoseconds()

	pipe := l.redisClient.Pipeline()
	pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(minScore, 10))
	pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
	pipe.Expire(ctx, key, window)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	count, err := l.redisClient.ZCard(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count <= limit, nil
}
