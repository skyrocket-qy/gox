package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    any
		wantErr bool
	}{
		// String conversions
		{"string_valid", "hello", "hello", false},
		{"string_empty", "", "", false},

		// Int conversions
		{"int_valid", "123", 123, false},
		{"int_negative", "-45", -45, false},
		{"int_invalid", "abc", 0, true},
		{"int_large", "9223372036854775807", int(9223372036854775807), false},
		{"int_overflow", "9223372036854775808", 0, true}, // Max int64 + 1

		// Float conversions
		{"float_valid", "1.23", 1.23, false},
		{"float_negative", "-4.56", -4.56, false},
		{"float_invalid", "xyz", 0.0, true},

		// Bool conversions
		{"bool_true", "true", true, false},
		{"bool_false", "false", false, false},
		{"bool_invalid", "other", false, true},

		// Unsupported type (will be tested with a specific type parameter)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.want.(type) {
			case string:
				got, err := StrToType[string](tt.input)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want.(string), got)
				}
			case int:
				got, err := StrToType[int](tt.input)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want.(int), got)
				}
			case float64:
				got, err := StrToType[float64](tt.input)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want.(float64), got)
				}
			case bool:
				got, err := StrToType[bool](tt.input)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.want.(bool), got)
				}
			}
		})
	}

	// Test unsupported type
	t.Run("unsupported_type", func(t *testing.T) {
		_, err := StrToType[complex64]("1+2i")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
	})
}
