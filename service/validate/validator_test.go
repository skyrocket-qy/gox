package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestValidator(t *testing.T) {
	// Initialize the validator
	New()

	t.Run("Valid struct", func(t *testing.T) {
		s := TestStruct{
			Name:  "Test User",
			Email: "test@example.com",
		}
		err := Struct(s)
		assert.NoError(t, err)
	})

	t.Run("Invalid struct - missing required field", func(t *testing.T) {
		s := TestStruct{
			Email: "test@example.com",
		}
		err := Struct(s)
		assert.Error(t, err)
	})

	t.Run("Invalid struct - invalid email", func(t *testing.T) {
		s := TestStruct{
			Name:  "Test User",
			Email: "not-an-email",
		}
		err := Struct(s)
		assert.Error(t, err)
	})
}
