package ginw

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCors(t *testing.T) {
	// Test non-OPTIONS request
	t.Run("Non-OPTIONS request", func(t *testing.T) {
		r := gin.New()
		r.Use(Cors())
		r.GET("/", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(
			t,
			"GET, POST, PUT, DELETE, OPTIONS",
			w.Header().Get("Access-Control-Allow-Methods"),
		)
		assert.Equal(
			t,
			"Origin, Content-Type, Authorization",
			w.Header().Get("Access-Control-Allow-Headers"),
		)
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})

	// Test OPTIONS request
	t.Run("OPTIONS request", func(t *testing.T) {
		r := gin.New()
		r.Use(Cors())
		r.OPTIONS("/", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodOptions, "/", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(
			t,
			"GET, POST, PUT, DELETE, OPTIONS",
			w.Header().Get("Access-Control-Allow-Methods"),
		)
		assert.Equal(
			t,
			"Origin, Content-Type, Authorization",
			w.Header().Get("Access-Control-Allow-Headers"),
		)
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})
}
