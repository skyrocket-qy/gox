package ginw

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestPanicRecover(t *testing.T) {
	// Suppress logging for this test
	originalLogger := log.Logger
	log.Logger = log.Output(io.Discard)

	defer func() {
		log.Logger = originalLogger
	}()

	gin.SetMode(gin.TestMode)

	t.Run("recovers from panic", func(t *testing.T) {
		r := gin.New()
		r.Use(PanicRecover())
		r.GET("/panic", func(c *gin.Context) {
			panic("test panic")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/panic", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("does not interfere with normal request", func(t *testing.T) {
		r := gin.New()
		r.Use(PanicRecover())
		r.GET("/normal", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/normal", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "ok", w.Body.String())
	})
}
