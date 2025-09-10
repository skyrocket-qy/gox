package ginw

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
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

func (a *InterAuthMid) CheckAuth(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			a.errBinder.Bind(c, erx.New(errcode.ErrMissingAuthorizationHeader))

			return
		}

		const bearerPrefix = "Bearer "

		token = strings.TrimPrefix(token, bearerPrefix)

		calm, err := ParseJWT(token, jwtSecret)
		if err != nil {
			a.errBinder.Bind(c, erx.W(err).SetCode(errcode.ErrUnauthorized))

			return
		}

		c.Set("userId", calm.Issuer)

		c.Next()
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
