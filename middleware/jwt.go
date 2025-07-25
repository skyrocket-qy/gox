package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/skyrocket-qy/erx"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			pkg.Bind(c, erx.New(pkg.ErrMissingAuthorizationHeader))
			c.Abort()
			return
		}

		tokenString := authHeader[7:]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return cfg.Cfg.Jwt.Secret, nil
		})
		if err != nil || !token.Valid {
			pkg.Bind(c, erx.New(pkg.ErrUnauthorized))
			c.Abort()
			return
		}

		// Token is valid, store the claims in the context
		if claims, ok := token.Claims.(jwt.RegisteredClaims); ok {
			c.Set("userID", claims.Subject)
		}

		// Continue to the next handler
		c.Next()
	}
}
