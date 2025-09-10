package ginw

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testCorsRequest(t *testing.T, method string, handler gin.HandlerFunc) {
	r := gin.New()
	r.Use(Cors())

	switch method {
	case http.MethodGet:
		r.GET("/", handler)
	case http.MethodOptions:
		r.OPTIONS("/", handler)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), method, "/", nil)
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
}

func TestCors(t *testing.T) {
	// Test non-OPTIONS request
	t.Run("Non-OPTIONS request", func(t *testing.T) {
		testCorsRequest(t, http.MethodGet, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
	})

	// Test OPTIONS request
	t.Run("OPTIONS request", func(t *testing.T) {
		testCorsRequest(t, http.MethodOptions, func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
	})
}
