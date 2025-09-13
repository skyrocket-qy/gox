package columnname

import "testing"

func TestToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello_world",
			want:  "HelloWorld",
		},
		{
			input: "foo_bar_baz",
			want:  "FooBarBaz",
		},
		{
			input: "single",
			want:  "Single",
		},
		{
			input: "",
			want:  "",
		},
		{
			input: "a_b_c",
			want:  "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToCamel(tt.input); got != tt.want {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
