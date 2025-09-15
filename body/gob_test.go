package body_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/body"

	"github.com/stretchr/testify/assert"
)

type GobTestStruct struct {
	Name    string
	Value   int
	private bool // gob should ignore private fields
}

func TestGobEncoding(t *testing.T) {
	t.Run("should successfully encode and decode a struct", func(t *testing.T) {
		// Arrange
		original := GobTestStruct{
			Name:    "test-struct",
			Value:   12345,
			private: true,
		}

		var decoded GobTestStruct

		// Act
		encodedBytes, err := body.GobEncode(original)

		// Assert Encode
		assert.NoError(t, err)
		assert.NotEmpty(t, encodedBytes)

		// Act Decode
		err = body.GobDecode(encodedBytes, &decoded)

		// Assert Decode
		assert.NoError(t, err)

		// We expect the decoded struct to match the original, but the private field won't be
		// encoded.
		// So we create an expected struct without the private field set.
		expected := GobTestStruct{
			Name:  "test-struct",
			Value: 12345,
		}
		assert.Equal(t, expected, decoded)
	})

	t.Run("should return an error when decoding invalid data", func(t *testing.T) {
		// Arrange
		invalidBytes := []byte("this is not gob data")

		var decoded GobTestStruct

		// Act
		err := body.GobDecode(invalidBytes, &decoded)

		// Assert
		assert.Error(t, err)
	})

	t.Run("should return an error when encoding an invalid type", func(t *testing.T) {
		// Arrange
		// Channels are not encodable by gob.
		ch := make(chan int)

		// Act
		_, err := body.GobEncode(ch)

		// Assert
		assert.Error(t, err)
	})

	t.Run("should return an error when decoding into a non-pointer", func(t *testing.T) {
		// Arrange
		original := GobTestStruct{Name: "test"}
		encodedBytes, _ := body.GobEncode(original)

		var decoded GobTestStruct

		// Act
		// Pass the struct by value, not by pointer.
		err := body.GobDecode(encodedBytes, decoded)

		// Assert
		assert.Error(t, err)
	})
}
