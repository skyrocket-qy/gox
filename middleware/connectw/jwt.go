package connectw

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
)

type AuthInterceptor struct {
	// No direct equivalent to httpx.ErrBinder in connectrpc,
	// errors are returned directly via connect.NewError
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (a *AuthInterceptor) WrapAuth(jwtSecret []byte) connect.Interceptor {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
				token := req.Header().Get("Authorization")
				if token == "" {
					return nil, connect.NewError(
						connect.CodeUnauthenticated,
						erx.New(errcode.ErrMissingAuthorizationHeader),
					)
				}

				const bearerPrefix = "Bearer "

				token = strings.TrimPrefix(token, bearerPrefix)

				calm, err := ParseJWT(token, jwtSecret)
				if err != nil {
					return nil, connect.NewError(
						connect.CodeUnauthenticated,
						erx.W(err).SetCode(errcode.ErrUnauthorized),
					)
				}

				// Store userId in context
				type contextKey string

				ctx = context.WithValue(ctx, contextKey("userId"), calm.Issuer)

				return next(ctx, req)
			},
		)
	})
}

func ParseJWT(tokenString string, secret []byte) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		// Check that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
