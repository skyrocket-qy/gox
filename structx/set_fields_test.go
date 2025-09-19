package structx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name            string
	Age             int
	Address         string
	unexportedField bool
}

func TestSetFields(t *testing.T) {
	t.Run("Successful case", func(t *testing.T) {
		s := &testStruct{}
		values := map[string]any{
			"Name":    "John",
			"Age":     30,
			"Address": "123 Main St",
		}
		err := SetFields(s, values)
		assert.NoError(t, err)
		assert.Equal(t, "John", s.Name)
		assert.Equal(t, 30, s.Age)
		assert.Equal(t, "123 Main St", s.Address)
	})

	t.Run("v is nil", func(t *testing.T) {
		err := SetFields(nil, nil)
		require.Error(t, err)
	})

	t.Run("v is not a pointer to a struct", func(t *testing.T) {
		s := testStruct{}
		err := SetFields(s, nil)
		require.Error(t, err)
	})

	t.Run("Field not found", func(t *testing.T) {
		s := &testStruct{}
		values := map[string]any{"NonExistent": "value"}
		err := SetFields(s, values)
		require.Error(t, err)
	})

	t.Run("Field cannot be set", func(t *testing.T) {
		s := &testStruct{}
		values := map[string]any{"unexportedField": true}
		err := SetFields(s, values)
		require.Error(t, err)
	})

	t.Run("Type mismatch", func(t *testing.T) {
		s := &testStruct{}
		values := map[string]any{"Age": "not an int"}
		err := SetFields(s, values)
		require.Error(t, err)
	})

	t.Run("Type conversion works", func(t *testing.T) {
		s := &testStruct{}
		values := map[string]any{"Age": int64(30)}
		err := SetFields(s, values)
		assert.NoError(t, err)
		assert.Equal(t, 30, s.Age)
	})
}
