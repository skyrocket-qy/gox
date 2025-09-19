package slidewindow_test

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/skyrocket-qy/gox/redisx/slidewindow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockClock implements the Clock interface for testing.
type MockClock struct {
	NowFunc func() time.Time
}

func (m *MockClock) Now() time.Time {
	return m.NowFunc()
}

func TestMovingWindowLimiter_Allow(t *testing.T) {
	// Create a mock Redis client
	client, mock := redismock.NewClientMock()

	// Create a mock clock
	fixedTime := time.Date(2025, time.January, 1, 10, 0, 0, 0, time.UTC)
	mockClock := &MockClock{
		NowFunc: func() time.Time { return fixedTime },
	}

	// Initialize the limiter with the mock clock
	limiter := slidewindow.New(client)
	limiter.SetClock(mockClock) // Assuming a SetClock method is added to MovingWindowLimiter

	key := "test_moving_window_key"
	limit := int64(3)
	window := 1 * time.Second

	ctx := context.Background()

	// Test case 1: All requests allowed within the limit
	for i := range limit {
		currentNow := fixedTime.Add(time.Duration(i) * 10 * time.Millisecond).UnixNano()
		minScore := currentNow - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(key, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
		mock.ExpectZAdd(key, redis.Z{Score: float64(currentNow), Member: currentNow}).SetVal(1)
		mock.ExpectExpire(key, window).SetVal(true)

		mock.ExpectZCard(key).SetVal(int64(i + 1))

		// Update mock clock for each iteration
		mockClock.NowFunc = func() time.Time { return time.Unix(0, currentNow) }

		allowed, err := limiter.Allow(ctx, key, limit, window)
		assert.NoError(t, err)
		assert.True(t, allowed)
	}

	// Test case 2: Request denied when exceeding the limit
	currentNow := fixedTime.Add(time.Duration(limit) * 10 * time.Millisecond).UnixNano()
	minScore := currentNow - window.Nanoseconds()

	mock.ExpectZRemRangeByScore(key, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
	mock.ExpectZAdd(key, redis.Z{Score: float64(currentNow), Member: currentNow}).SetVal(1)
	mock.ExpectExpire(key, window).SetVal(true)

	mock.ExpectZCard(key).SetVal(limit + 1)

	mockClock.NowFunc = func() time.Time { return time.Unix(0, currentNow) }

	allowed, err := limiter.Allow(ctx, key, limit, window)
	assert.NoError(t, err)
	assert.False(t, allowed)

	// Test case 3: Error during Redis pipeline execution (mocking ZRemRangeByScore error)
	currentNow = fixedTime.Add(time.Duration(limit+1) * 10 * time.Millisecond).UnixNano()
	minScore = currentNow - window.Nanoseconds()
	mock.ExpectZRemRangeByScore(key, "0", strconv.FormatInt(minScore, 10)).
		SetErr(errors.New("redis ZRemRangeByScore error"))

	mockClock.NowFunc = func() time.Time { return time.Unix(0, currentNow) }

	allowed, err = limiter.Allow(ctx, key, limit, window)
	require.Error(t, err)
	assert.False(t, allowed)
	assert.Contains(t, err.Error(), "redis ZRemRangeByScore error")

	// Test case 4: Error during ZCard operation
	currentNow = fixedTime.Add(time.Duration(limit+2) * 10 * time.Millisecond).UnixNano()
	minScore = currentNow - window.Nanoseconds()
	mock.ExpectZRemRangeByScore(key, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
	mock.ExpectZAdd(key, redis.Z{Score: float64(currentNow), Member: currentNow}).SetVal(1)
	mock.ExpectExpire(key, window).SetVal(true)

	mock.ExpectZCard(key).SetErr(errors.New("redis zcard error"))

	mockClock.NowFunc = func() time.Time { return time.Unix(0, currentNow) }

	allowed, err = limiter.Allow(ctx, key, limit, window)
	require.Error(t, err)
	assert.False(t, allowed)
	assert.Contains(t, err.Error(), "redis zcard error")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
