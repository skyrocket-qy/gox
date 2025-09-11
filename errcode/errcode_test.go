package errcode

import "testing"

func TestErr_Str(t *testing.T) {
	tests := []struct {
		name string
		c    err
		want string
	}{
		{
			name: "ErrBadRequest",
			c:    ErrBadRequest,
			want: "400.0000",
		},
		{
			name: "ErrNotFound",
			c:    ErrNotFound,
			want: "404.0000",
		},
		{
			name: "ErrUnknown",
			c:    ErrUnknown,
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
