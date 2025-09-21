package excel_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/excel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			switch gotType := tt.want.(type) {
			case string:
				got, err := excel.StrToType[string](tt.input)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, gotType, got)
				}
			case int:
				got, err := excel.StrToType[int](tt.input)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, gotType, got)
				}
			case float64:
				got, err := excel.StrToType[float64](tt.input)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.InDelta(t, gotType, got, 0.000001)
				}
			case bool:
				got, err := excel.StrToType[bool](tt.input)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, gotType, got)
				}
			}
		})
	}

	// Test unsupported type
	t.Run("unsupported_type", func(t *testing.T) {
		_, err := excel.StrToType[complex64]("1+2i")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
	})
}
