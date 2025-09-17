package bitx_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/bitx"
	"github.com/stretchr/testify/assert"
)

func TestMultiple2(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 10, 20},
		{"zero", 0, 0},
		{"negative", -5, -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.Multiple2(tt.input))
		})
	}
}

func TestMultiple4(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 10, 40},
		{"zero", 0, 0},
		{"negative", -5, -20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.Multiple4(tt.input))
		})
	}
}

func TestDivide4(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"positive", 40, 10},
		{"positive with remainder", 43, 10},
		{"zero", 0, 0},
		{"negative", -40, -10},
		{"negative with remainder", -39, -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.Divide4(tt.input))
		})
	}
}

func TestCountBit1Builtin(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  uint32
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"all ones", 0xffffffff, 32},
		{"some ones", 0x12345678, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.CountBit1Builtin(tt.input))
		})
	}
}

func TestCountBit1(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"seven", 7, 3},
		{"max int32", 0x7fffffff, 31},
		{"negative one", -1, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.CountBit1(tt.input))
		})
	}
}

func TestCountBit1ChangeIn(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"seven", 7, 3},
		{"max int32", 0x7fffffff, 31},
		{"negative one", -1, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.CountBit1ChangeIn(tt.input))
		})
	}
}

func TestCountBit0(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 32},
		{"one", 1, 31},
		{"seven", 7, 29},
		{"max int32", 0x7fffffff, 1},
		{"negative one", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.CountBit0(tt.input))
		})
	}
}

func TestCountBit0ChangeIn(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 32},
		{"one", 1, 31},
		{"seven", 7, 29},
		{"max int32", 0x7fffffff, 1},
		{"negative one", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.CountBit0ChangeIn(tt.input))
		})
	}
}

func TestMarkNThBitTo1(t *testing.T) {
	tests := []struct {
		name string
		in   int
		n    int
		want int
	}{
		{"mark 3rd bit of 0", 0, 3, 4},
		{"mark 1st bit of 8", 8, 1, 9},
		{"mark bit that is already 1", 5, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.MarkNThBitTo1(tt.in, tt.n))
		})
	}
}

func TestReverseNthBit(t *testing.T) {
	tests := []struct {
		name string
		in   int
		n    int
		want int
	}{
		{"reverse 3rd bit of 0", 0, 3, 4},
		{"reverse 1st bit of 9", 9, 1, 8},
		{"reverse 1st bit of 8", 8, 1, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.ReverseNthBit(tt.in, tt.n))
		})
	}
}

func TestClearNthBit(t *testing.T) {
	tests := []struct {
		name string
		in   int
		n    int
		want int
	}{
		{"clear 3rd bit of 4", 4, 3, 0},
		{"clear 1st bit of 9", 9, 1, 8},
		{"clear bit that is already 0", 8, 1, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.ClearNthBit(tt.in, tt.n))
		})
	}
}

func TestReverseBitBuiltin(t *testing.T) {
	tests := []struct {
		name  string
		input uint32
		want  uint32
	}{
		{"zero", 0, 0},
		{"one", 1, 0x80000000},
		{"all ones", 0xffffffff, 0xffffffff},
		{"some value", 0x12345678, 0x1e6a2c48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.ReverseBitBuiltin(tt.input))
		})
	}
}

func TestLeastSignificantBit1(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 2},
		{"three", 3, 1},
		{"four", 4, 4},
		{"six", 6, 2},
		{"negative one", -1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.LeastSignificantBit1(tt.input))
		})
	}
}

func TestIsPowerOf2(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  bool
	}{
		{"zero", 0, false},
		{"one", 1, true},
		{"two", 2, true},
		{"three", 3, false},
		{"four", 4, true},
		{"large power of 2", 1024, true},
		{"large non power of 2", 1023, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.IsPowerOf2(tt.input))
		})
	}
}

func TestFindLackNum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"missing 3", []int{0, 1, 2, 4}, 3},
		{"missing 0", []int{1, 2, 3}, 0},
		{"missing last", []int{0, 1, 2}, 3},
		{"empty", []int{}, 0},
		{"one element missing 0", []int{1}, 0},
		{"one element missing 1", []int{0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.FindLackNum(tt.input))
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  byte
	}{
		{"uppercase A", 'A', 'a'},
		{"uppercase Z", 'Z', 'z'},
		{"lowercase a", 'a', 'a'},
		{"not a letter", '5', '5'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.ToLower(tt.input))
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  byte
	}{
		{"lowercase a", 'a', 'A'},
		{"lowercase z", 'z', 'Z'},
		{"uppercase A", 'A', 'A'},
		{"not a letter", '5', '5'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bitx.ToUpper(tt.input))
		})
	}
}
