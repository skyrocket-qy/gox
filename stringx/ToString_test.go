package stringx_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/stringx"
)

func TestToString(t *testing.T) {
	type testStruct struct {
		s string
	}

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"string", "hello", "hello"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"int", 123, "123"},
		{"int8", int8(12), "12"},
		{"int16", int16(123), "123"},
		{"int32", int32(123), "123"},
		{"int64", int64(123), "123"},
		{"uint", uint(123), "123"},
		{"uint8", uint8(12), "12"},
		{"uint16", uint16(123), "123"},
		{"uint32", uint32(123), "123"},
		{"uint64", uint64(123), "123"},
		{"float32", float32(1.23), "1.23"},
		{"float64", float64(1.23), "1.23"},
		{"struct", testStruct{s: "test"}, "{test}"},
		{"nil", nil, "<nil>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.ToString(tt.input); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
