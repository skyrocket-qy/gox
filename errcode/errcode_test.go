package errcode_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/errcode"
)

func TestErr_Str(t *testing.T) {
	tests := []struct {
		name string
		c    errcode.Err
		want string
	}{
		{
			name: "ErrBadRequest",
			c:    errcode.ErrBadRequest,
			want: "400.0000",
		},
		{
			name: "ErrNotFound",
			c:    errcode.ErrNotFound,
			want: "404.0000",
		},
		{
			name: "ErrUnknown",
			c:    errcode.ErrUnknown,
			want: "500.0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Str(); got != tt.want {
				t.Errorf("err.Str() = %v, want %v", got, tt.want)
			}
		})
	}
}
