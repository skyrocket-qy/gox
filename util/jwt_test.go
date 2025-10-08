package util

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenRefreshToken(t *testing.T) {
	t.Run("should generate a valid UUID", func(t *testing.T) {
		token1 := GenRefreshToken()
		_, err := uuid.Parse(token1)
		assert.NoError(t, err)
	})

	t.Run("should generate different tokens on subsequent calls", func(t *testing.T) {
		token1 := GenRefreshToken()
		token2 := GenRefreshToken()
		assert.NotEqual(t, token1, token2)
	})
}
