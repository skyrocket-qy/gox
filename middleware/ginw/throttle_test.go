package ginw

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func TestThrottle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, mock := redismock.NewClientMock()

	limit := int64(2)
	window := time.Second * 10
	keyPrefix := "test_throttle"
	clientIP := "127.0.0.1"
	expectedKey := keyPrefix + ":" + clientIP

	keyExtractor := func(c *gin.Context) string {
		return c.ClientIP()
	}

	throttle := NewThrottle(db, limit, window, keyPrefix, keyExtractor)
	mockClock := &mockClock{currentTime: time.Now()}
	throttle.clock = mockClock

	r := gin.New()
	r.Use(throttle.Middleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// First request - should be allowed
	t.Run("first request", func(t *testing.T) {
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedKey, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
		mock.ExpectZAdd(expectedKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedKey, window).SetVal(true)
		mock.ExpectZCard(expectedKey).SetVal(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = clientIP + ":1234"
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Second request - should be allowed
	t.Run("second request", func(t *testing.T) {
		mockClock.Advance(time.Second)
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedKey, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
		mock.ExpectZAdd(expectedKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedKey, window).SetVal(true)
		mock.ExpectZCard(expectedKey).SetVal(2)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = clientIP + ":1234"
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Third request - should be throttled
	t.Run("third request throttled", func(t *testing.T) {
		mockClock.Advance(time.Second)
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedKey, "0", strconv.FormatInt(minScore, 10)).SetVal(0)
		mock.ExpectZAdd(expectedKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedKey, window).SetVal(true)
		mock.ExpectZCard(expectedKey).SetVal(3)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = clientIP + ":1234"
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})

	// Request after window - should be allowed
	t.Run("request after window", func(t *testing.T) {
		mockClock.Advance(window)
		now := mockClock.Now().UnixNano()
		minScore := now - window.Nanoseconds()

		mock.ExpectZRemRangeByScore(expectedKey, "0", strconv.FormatInt(minScore, 10)).
			SetVal(2)
			// 2 old requests removed
		mock.ExpectZAdd(expectedKey, redis.Z{Score: float64(now), Member: now}).SetVal(1)
		mock.ExpectExpire(expectedKey, window).SetVal(true)
		mock.ExpectZCard(expectedKey).SetVal(1) // Back to 1 request in window

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = clientIP + ":1234"
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
