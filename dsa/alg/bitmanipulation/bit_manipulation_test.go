package bitmanipulation

import "testing"

func TestBitManipulation(t *testing.T) {
	t.Run("Multiple2", func(t *testing.T) {
		if got := Multiple2(5); got != 10 {
			t.Errorf("Multiple2() = %v, want %v", got, 10)
		}
	})

	t.Run("Multiple4", func(t *testing.T) {
		if got := Multiple4(5); got != 20 {
			t.Errorf("Multiple4() = %v, want %v", got, 20)
		}
	})

	t.Run("Divide4", func(t *testing.T) {
		if got := Divide4(20); got != 5 {
			t.Errorf("Divide4() = %v, want %v", got, 5)
		}
	})

	t.Run("CountBit1Builtin", func(t *testing.T) {
		if got := CountBit1Builtin(7); got != 3 {
			t.Errorf("CountBit1Builtin() = %v, want %v", got, 3)
		}
	})

	t.Run("CountBit1", func(t *testing.T) {
		if got := CountBit1(7); got != 3 {
			t.Errorf("CountBit1() = %v, want %v", got, 3)
		}
	})

	t.Run("CountBit1ChangeIn", func(t *testing.T) {
		if got := CountBit1ChangeIn(7); got != 3 {
			t.Errorf("CountBit1ChangeIn() = %v, want %v", got, 3)
		}
	})

	t.Run("CountBit0", func(t *testing.T) {
		// Assuming 32-bit integers, 7 is 0...0111, so it should have 32-3=29 zero bits.
		if got := CountBit0(7); got != 29 {
			t.Errorf("CountBit0() = %v, want %v", got, 29)
		}
	})

	t.Run("CountBit0ChangeIn", func(t *testing.T) {
		if got := CountBit0ChangeIn(7); got != 29 {
			t.Errorf("CountBit0ChangeIn() = %v, want %v", got, 29)
		}
	})

	t.Run("MarkNThBitTo1", func(t *testing.T) {
		// 5 is 0101, mark 3rd bit (from right, 1-based) to 1 -> 0101 | 0100 = 0101 = 5 (already 1)
		// 8 is 1000, mark 3rd bit to 1 -> 1000 | 0100 = 1100 = 12
		if got := MarkNThBitTo1(8, 3); got != 12 {
			t.Errorf("MarkNThBitTo1() = %v, want %v", got, 12)
		}
		if got := MarkNThBitTo1(5, 3); got != 5 {
			t.Errorf("MarkNThBitTo1() = %v, want %v", got, 5)
		}
	})

	t.Run("ReverseNthBit", func(t *testing.T) {
		// 5 is 0101, reverse 3rd bit -> 0101 ^ 0100 = 0001 = 1
		if got := ReverseNthBit(5, 3); got != 1 {
			t.Errorf("ReverseNthBit() = %v, want %v", got, 1)
		}
		// 8 is 1000, reverse 3rd bit -> 1000 ^ 0100 = 1100 = 12
		if got := ReverseNthBit(8, 3); got != 12 {
			t.Errorf("ReverseNthBit() = %v, want %v", got, 12)
		}
	})

	t.Run("ClearNthBit", func(t *testing.T) {
		// 5 is 0101, clear 3rd bit -> 0101 & ~(0100) = 0101 & 1011 = 0001 = 1
		if got := ClearNthBit(5, 3); got != 1 {
			t.Errorf("ClearNthBit() = %v, want %v", got, 1)
		}
		// 8 is 1000, clear 3rd bit -> 1000 & ~(0100) = 1000 & 1011 = 1000 = 8
		if got := ClearNthBit(8, 3); got != 8 {
			t.Errorf("ClearNthBit() = %v, want %v", got, 8)
		}
	})

	t.Run("ReverseBitBuiltin", func(t *testing.T) {
		var x uint32 = 25
		// 25 = 0...011001
		// reverse should be 100110...0 = 2,550,136,832
		var expected uint32 = 2550136832
		if got := ReverseBitBuiltin(x); got != expected {
			t.Errorf("ReverseBitBuiltin() = %v, want %v", got, expected)
		}
	})

	t.Run("LeastSignificantBit1", func(t *testing.T) {
		// 12 is 1100, LSB is 100 = 4
		if got := LeastSignificantBit1(12); got != 4 {
			t.Errorf("LeastSignificantBit1() = %v, want %v", got, 4)
		}
		// 10 is 1010, LSB is 10 = 2
		if got := LeastSignificantBit1(10); got != 2 {
			t.Errorf("LeastSignificantBit1() = %v, want %v", got, 2)
		}
	})

	t.Run("IsPowerOf2", func(t *testing.T) {
		if !IsPowerOf2(8) {
			t.Errorf("IsPowerOf2(8) = false, want true")
		}
		if IsPowerOf2(9) {
			t.Errorf("IsPowerOf2(9) = true, want false")
		}
		if IsPowerOf2(0) {
			t.Errorf("IsPowerOf2(0) = true, want false")
		}
	})

	t.Run("FindLackNum", func(t *testing.T) {
		in := []int{0, 1, 2, 4}
		if got := FindLackNum(in); got != 3 {
			t.Errorf("FindLackNum() = %v, want %v", got, 3)
		}
	})

	t.Run("ToLower", func(t *testing.T) {
		if got := ToLower('A'); got != 'a' {
			t.Errorf("ToLower('A') = %c, want %c", got, 'a')
		}
		if got := ToLower('z'); got != 'z' {
			t.Errorf("ToLower('z') = %c, want %c", got, 'z')
		}
	})

	t.Run("ToUpper", func(t *testing.T) {
		if got := ToUpper('a'); got != 'A' {
			t.Errorf("ToUpper('a') = %c, want %c", got, 'A')
		}
		if got := ToUpper('Z'); got != 'Z' {
			t.Errorf("ToUpper('Z') = %c, want %c", got, 'Z')
		}
	})
}
