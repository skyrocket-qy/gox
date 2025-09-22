package longestcommonsubsequence_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/longestcommonsubsequence"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
		{
			name: "Case 4",
			args: args{
				text1: "",
				text2: "a",
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				text1: "a",
				text2: "",
			},
			want: 0,
		},
		{
			name: "Case 6",
			args: args{
				text1: "",
				text2: "",
			},
			want: 0,
		},
		{
			name: "Case 7",
			args: args{
				text1: "bsbininm",
				text2: "jmjkbkjkv",
			},
			want: 1,
		},
		{
			name: "Case 8",
			args: args{
				text1: "oxcpqrsvwf",
				text2: "shmtulqrypy",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestcommonsubsequence.LongestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
