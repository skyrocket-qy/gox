package connectw_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/httpx/middleware/connectw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMovingWindowLimiter is a mock implementation of redisx.MovingWindowLimiterInterface for
// testing.
type MockMovingWindowLimiter struct {
	mock.Mock
}

func (m *MockMovingWindowLimiter) Allow(
	ctx context.Context,
	key string,
	limit int64,
	window time.Duration,
) (bool, error) {
	args := m.Called(ctx, key, limit, window)

	return args.Bool(0), args.Error(1)
}

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
	// Create a mock for MovingWindowLimiter
	mockLimiter := new(MockMovingWindowLimiter)

	// Create a mock Redis client (though it won't be used by Throttle directly now)
	_, _ = redismock.NewClientMock() // Keep redismock import for this line

	limit := int64(1)
	window := time.Second * 10
	keyPrefix := "test_throttle_connect"
	testKey := "test_key"

	keyExtractor := func(ctx context.Context) string {
		return testKey
	}

	// Initialize Throttle with the mock limiter
	throttle := connectw.NewThrottle(nil, limit, window, keyPrefix, keyExtractor, mockLimiter)

	interceptor := throttle.UnaryInterceptor()

	mockHandler := connect.UnaryFunc(
		func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return connect.NewResponse(&struct{}{}), nil
		},
	)

	// Test case 1: allows request within limit
	t.Run("allows request within limit", func(t *testing.T) {
		// Expect the Allow method to be called and return true (allowed)
		mockLimiter.On("Allow", mock.Anything, mock.AnythingOfType("string"), limit, window).
			Return(true, nil).
			Once()

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.NoError(t, err)
		mockLimiter.AssertExpectations(t)
	})

	// Test case 2: throttles request over limit
	t.Run("throttles request over limit", func(t *testing.T) {
		// Expect the Allow method to be called and return false (denied)
		mockLimiter.On("Allow", mock.Anything, mock.AnythingOfType("string"), limit, window).
			Return(false, nil).
			Once()

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.Error(t, err)
		connectErr := &connect.Error{}
		assert.ErrorAs(t, err, &connectErr)
		assert.Equal(t, connect.CodeResourceExhausted, connectErr.Code())
		mockLimiter.AssertExpectations(t)
	})

	// Test case 3: error from limiter
	t.Run("returns error from limiter", func(t *testing.T) {
		// Expect the Allow method to be called and return an error
		mockLimiter.On("Allow", mock.Anything, mock.AnythingOfType("string"), limit, window).
			Return(false, errors.New("limiter error")).
			Once()

		_, err := interceptor(mockHandler)(context.Background(), connect.NewRequest(&struct{}{}))
		assert.Error(t, err)
		connectErr := &connect.Error{}
		assert.ErrorAs(t, err, &connectErr)
		assert.Equal(t, connect.CodeInternal, connectErr.Code())
		assert.Contains(t, err.Error(), "limiter error")
		mockLimiter.AssertExpectations(t)
	})
}

func TestNewThrottle_NilLimiter(t *testing.T) {
	db, _ := redismock.NewClientMock()
	throttle := connectw.NewThrottle(
		db,
		1,
		time.Second,
		"test",
		func(ctx context.Context) string { return "" },
		nil,
	)
	assert.NotNil(t, throttle)
}
