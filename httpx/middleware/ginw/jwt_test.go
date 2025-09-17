package ginw_test

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
	"github.com/skyrocket-qy/gox/httpx"
	"github.com/skyrocket-qy/gox/httpx/middleware/ginw"
	"github.com/stretchr/testify/assert"
)

const testJWTSecret = "test-secret"

func generateTestJWT(issuer string, secret []byte, expiresAt time.Time) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func TestInterAuthMid_CheckAuth(t *testing.T) {
	// Suppress logging for this test
	originalLogger := log.Logger
	log.Logger = log.Output(io.Discard)

	defer func() {
		log.Logger = originalLogger
	}()

	gin.SetMode(gin.TestMode)

	authMid := ginw.NewInterAuthMid()
	authMid.ErrBinder = httpx.NewErrBinder(map[erx.Code]int{
		errcode.ErrUnauthorized:               http.StatusUnauthorized,
		errcode.ErrMissingAuthorizationHeader: http.StatusUnauthorized,
	})

	// Create a router with the middleware
	r := gin.New()
	r.Use(authMid.CheckAuth([]byte(testJWTSecret)))
	r.GET("/protected", func(c *gin.Context) {
		userId, exists := c.Get("userId")
		if !exists {
			c.Status(http.StatusUnauthorized)

			return
		}

		if userIdStr, ok := userId.(string); ok {
			c.String(http.StatusOK, "welcome "+userIdStr)
		} else {
			c.Status(http.StatusInternalServerError) // Or log an error, this should not happen in a correct flow
		}
	})

	t.Run("Valid token", func(t *testing.T) {
		token, err := generateTestJWT("test-user", []byte(testJWTSecret), time.Now().Add(time.Hour))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/protected",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "welcome test-user", w.Body.String())
	})

	t.Run("Invalid token - bad signature", func(t *testing.T) {
		token, err := generateTestJWT(
			"test-user",
			[]byte("wrong-secret"),
			time.Now().Add(time.Hour),
		)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/protected",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Expired token", func(t *testing.T) {
		token, err := generateTestJWT(
			"test-user",
			[]byte(testJWTSecret),
			time.Now().Add(-time.Hour),
		)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/protected",
			nil,
		)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Missing Authorization header", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/protected",
			nil,
		)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Malformed Authorization header", func(t *testing.T) {
		token, err := generateTestJWT("test-user", []byte(testJWTSecret), time.Now().Add(time.Hour))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			"/protected",
			nil,
		)
		// No "Bearer " prefix
		req.Header.Set("Authorization", token)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestParseJWT(t *testing.T) {
	t.Run("Unexpected signing method", func(t *testing.T) {
		// Create a token with a different signing method (e.g., ES256)
		claims := &jwt.RegisteredClaims{
			Issuer: "test-user",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

		// We can't sign it with a private key, so we'll just take the unsigned token string
		unsignedToken, err := token.SigningString()
		assert.NoError(t, err)

		// Add a bogus but validly encoded signature
		signature := base64.RawURLEncoding.EncodeToString([]byte("bogussignature"))
		tokenString := unsignedToken + "." + signature

		_, err = ginw.ParseJWT(tokenString, []byte(testJWTSecret))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected signing method")
	})
}
