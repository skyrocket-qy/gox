package stringmatch_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/stringmatch"
	"github.com/stretchr/testify/assert"
)

func TestBoyerMoore(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		pattern  string
		expected []int
	}{
		{
			name:     "Simple case",
			text:     "ABAAABCD",
			pattern:  "ABC",
			expected: []int{4},
		},
		{
			name:     "Multiple occurrences",
			text:     "ABAAABCDABCD",
			pattern:  "ABCD",
			expected: []int{4, 8},
		},
		{
			name:     "Pattern at the beginning",
			text:     "ABCDABAA",
			pattern:  "ABCD",
			expected: []int{0},
		},
		{
			name:     "Pattern at the end",
			text:     "ABAAABCD",
			pattern:  "ABCD",
			expected: []int{4},
		},
		{
			name:     "No occurrence",
			text:     "ABAAABCD",
			pattern:  "XYZ",
			expected: []int{},
		},
		{
			name:     "Empty text",
			text:     "",
			pattern:  "ABC",
			expected: []int{},
		},
		{
			name:     "Empty pattern",
			text:     "ABAAABCD",
			pattern:  "",
			expected: []int{},
		},
		{
			name:     "Pattern longer than text",
			text:     "ABC",
			pattern:  "ABCD",
			expected: []int{},
		},
		{
			name:     "Complex case",
			text:     "HERE IS A SIMPLE EXAMPLE, EXAMPLE",
			pattern:  "EXAMPLE",
			expected: []int{17, 26},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, stringmatch.BoyerMoore(tc.text, tc.pattern))
		})
	}
}
