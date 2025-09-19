package connectw_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/skyrocket-qy/gox/httpx/middleware/connectw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthInterceptor_WrapAuth(t *testing.T) {
	secret := []byte("test_secret")

	// Create a new AuthInterceptor
	authInterceptor := connectw.NewAuthInterceptor()

	// Create a dummy next function for the interceptor chain
	next := connect.UnaryFunc(
		func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			return connect.NewResponse(&struct{}{}), nil
		},
	)

	// Get the interceptor
	interceptor := authInterceptor.WrapAuth(secret)

	// Test Case 1: Missing Authorization header
	t.Run("Missing Authorization Header", func(t *testing.T) {
		req := connect.NewRequest(&struct{}{})
		_, err := interceptor.WrapUnary(next)(context.Background(), req)
		require.Error(t, err)
		connectErr := &connect.Error{}
		ok := errors.As(err, &connectErr)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeUnauthenticated, connectErr.Code())
		assert.Contains(t, connectErr.Message(), "401.0002")
	})

	// Test Case 2: Invalid JWT token
	t.Run("Invalid JWT Token", func(t *testing.T) {
		req := connect.NewRequest(&struct{}{})
		req.Header().Set("Authorization", "Bearer invalid_token")
		_, err := interceptor.WrapUnary(next)(context.Background(), req)
		require.Error(t, err)
		connectErr := &connect.Error{}
		ok := errors.As(err, &connectErr)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeUnauthenticated, connectErr.Code())
		assert.Contains(t, connectErr.Message(), "401.0000")
	})

	// Test Case 3: Valid JWT token
	t.Run("Valid JWT Token", func(t *testing.T) {
		// Create a valid token
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user_id",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secret)

		req := connect.NewRequest(&struct{}{})
		req.Header().Set("Authorization", "Bearer "+tokenString)

		// Custom next function to check context value
		customNext := connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				userID, ok := connectw.UserIDFromContext(ctx)
				assert.True(t, ok)
				assert.Equal(t, "test_user_id", userID)

				return connect.NewResponse(&struct{}{}), nil
			},
		)

		_, err := interceptor.WrapUnary(customNext)(context.Background(), req)
		assert.NoError(t, err)
	})

	// Test Case 4: Expired JWT token
	t.Run("Expired JWT Token", func(t *testing.T) {
		// Create an expired token
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user_id",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // Expired an hour ago
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secret)

		req := connect.NewRequest(&struct{}{})
		req.Header().Set("Authorization", "Bearer "+tokenString)

		_, err := interceptor.WrapUnary(next)(context.Background(), req)
		require.Error(t, err)
		connectErr := &connect.Error{}
		ok := errors.As(err, &connectErr)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeUnauthenticated, connectErr.Code())
		assert.Contains(t, connectErr.Message(), "401.0000")
	})

	// Test Case 5: Token with wrong signing method
	t.Run("Wrong Signing Method", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user_id",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodNone, claims) // Use a different signing method
		tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

		req := connect.NewRequest(&struct{}{})
		req.Header().Set("Authorization", "Bearer "+tokenString)

		_, err := interceptor.WrapUnary(next)(context.Background(), req)
		require.Error(t, err)
		connectErr := &connect.Error{}
		ok := errors.As(err, &connectErr)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeUnauthenticated, connectErr.Code())
		assert.Contains(t, connectErr.Message(), "401.0000")
	})

	// Test Case 6: Token with wrong secret
	t.Run("Wrong Secret", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user_id",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte("wrong_secret"))

		req := connect.NewRequest(&struct{}{})
		req.Header().Set("Authorization", "Bearer "+tokenString)

		_, err := interceptor.WrapUnary(next)(context.Background(), req)
		require.Error(t, err)
		connectErr := &connect.Error{}
		ok := errors.As(err, &connectErr)
		assert.True(t, ok)
		assert.Equal(t, connect.CodeUnauthenticated, connectErr.Code())
		assert.Contains(t, connectErr.Message(), "401.0000")
	})
}

func TestParseJWT(t *testing.T) {
	secret := []byte("test_secret")

	// Test Case 1: Valid token
	t.Run("Valid Token", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secret)

		parsedClaims, err := connectw.ParseJWT(tokenString, secret)
		assert.NoError(t, err)
		assert.Equal(t, "test_user", parsedClaims.Issuer)
	})

	// Test Case 2: Invalid token string
	t.Run("Invalid Token String", func(t *testing.T) {
		_, err := connectw.ParseJWT("invalid_token_string", secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "token is malformed")
	})

	// Test Case 3: Expired token
	t.Run("Expired Token", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secret)

		_, err := connectw.ParseJWT(tokenString, secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "token is expired")
	})

	// Test Case 4: Token with wrong secret
	t.Run("Wrong Secret", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte("wrong_secret"))

		_, err := connectw.ParseJWT(tokenString, secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "signature is invalid")
	})

	// Test Case 5: Token with unexpected signing method
	t.Run("Unexpected Signing Method", func(t *testing.T) {
		claims := &jwt.RegisteredClaims{
			Issuer:    "test_user",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
		tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

		_, err := connectw.ParseJWT(tokenString, secret)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected signing method")
	})
}
