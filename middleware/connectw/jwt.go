package connectw

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
	"github.com/skyrocket-qy/gox/httpx"
)

type InterAuthMid struct {
	errBinder *httpx.ErrBinder
}

func NewInterAuthMid() *InterAuthMid {
	return &InterAuthMid{}
}

func (a *InterAuthMid) CheckAuth(jwtSecret []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				a.errBinder.Bind(w, r, erx.New(errcode.ErrMissingAuthorizationHeader))
				return
			}

			const bearerPrefix = "Bearer "
			token = strings.TrimPrefix(token, bearerPrefix)

			calm, err := ParseJWT(token, jwtSecret)
			if err != nil {
				a.errBinder.Bind(w, r, erx.W(err).SetCode(errcode.ErrUnauthorized))
				return
			}

			ctx := context.WithValue(r.Context(), "userId", calm.Issuer)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ParseJWT(tokenString string, secret []byte) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
