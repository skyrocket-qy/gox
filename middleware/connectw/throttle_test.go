package connectw

import (
	"context"
	"strconv"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// mockClock is a mock implementation of the Clock interface for testing.
type mockClock struct {
	currentTime time.Time
}

func (m *mockClock) Now() time.Time {
	return m.currentTime
}

func (m *mockClock) Advance(d time.Duration) {
	m.currentTime = m.currentTime.Add(d)
}

func TestThrottle_UnaryInterceptor(t *testing.T) {
	db, mock := redismock.NewClientMock()

	limit := int64(1)
	window := time.Second * 10
	keyPrefix := "test_throttle_connect"
	testKey := "test_key"
	expectedRedisKey := keyPrefix + ":" + testKey

	keyExtractor := func(ctx context.Context) string {
		return testKey
	}

	throttle := NewThrottle(db, limit, window, keyPrefix, keyExtractor)
	mockClock := &mockClock{currentTime: time.Now()}
	throttle.clock = mockClock

	interceptor := throttle.UnaryInterceptor()

	mockHandler := connect.UnaryFunc(
		func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return connect.NewResponse(&struct{}{}), nil
		},
	)

	// First request - should be allowed
	t.Run("allows request within limit", func(t *testing.T) {
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedRedisKey, "0", strconv.FormatInt(minScore, 10)).
			SetVal(0)
		mock.ExpectZAdd(expectedRedisKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedRedisKey, window).SetVal(true)
		mock.ExpectZCard(expectedRedisKey).SetVal(1)

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.NoError(t, err)
	})

	// Second request - should be throttled
	t.Run("throttles request over limit", func(t *testing.T) {
		mockClock.Advance(time.Second)
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedRedisKey, "0", strconv.FormatInt(minScore, 10)).
			SetVal(0)
		mock.ExpectZAdd(expectedRedisKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedRedisKey, window).SetVal(true)
		mock.ExpectZCard(expectedRedisKey).SetVal(2)

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.Error(t, err)
		connectErr := &connect.Error{}
		assert.ErrorAs(t, err, &connectErr)
		assert.Equal(t, connect.CodeResourceExhausted, connectErr.Code())
	})

	// Request after window - should be allowed
	t.Run("allows request after window", func(t *testing.T) {
		mockClock.Advance(window)
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedRedisKey, "0", strconv.FormatInt(minScore, 10)).
			SetVal(1)
		mock.ExpectZAdd(expectedRedisKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedRedisKey, window).SetVal(true)
		mock.ExpectZCard(expectedRedisKey).SetVal(1)

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.NoError(t, err)
	})
}
