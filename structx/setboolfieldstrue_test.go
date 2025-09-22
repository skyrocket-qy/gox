package structx_test

import (
	"reflect"
	"testing"

	"github.com/skyrocket-qy/gox/structx"
)

type Example struct {
	Field1 bool
	Field2 Nested
}

type Nested struct {
	Field3 bool
	Field4 bool
}

func Test_SetBoolFieldsTrue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name: "Set boolean fields to true",
			input: &Example{
				Field1: false,
				Field2: Nested{
					Field3: false,
					Field4: false,
				},
			},
			expected: &Example{
				Field1: true,
				Field2: Nested{
					Field3: true,
					Field4: true,
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := structx.SetBoolFieldsTrue(tt.input)
			if err != nil {
				t.Errorf("SetBoolFieldsTrue() error = %v", err)

				return
			}

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("SetBoolFieldsTrue() = %v, want %v", tt.input, tt.expected)
			}
		})
	}

	t.Run("nil input", func(t *testing.T) {
		t.Parallel()

		err := structx.SetBoolFieldsTrue(nil)
		if err == nil {
			t.Error("SetBoolFieldsTrue() error = nil, want error")
		}
	})

	t.Run("not a pointer to a struct", func(t *testing.T) {
		t.Parallel()

		err := structx.SetBoolFieldsTrue(Example{})
		if err == nil {
			t.Error("SetBoolFieldsTrue() error = nil, want error")
		}
	})
}
