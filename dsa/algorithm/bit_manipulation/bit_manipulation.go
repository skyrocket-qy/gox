package bitmanipulation

/* @tags: bit manipulation */

func Multiple2(in int) int {
	return in << 1
}

func Multiple4(in int) int {
	return in << 2
}

func Divide4(in int) int {
	return in >> 2
}

func CountBit1Builtin(x uint32) uint32 {
	x = (x & 0x55555555) + ((x & 0xaaaaaaaa) >> 1)
	x = (x & 0x33333333) + ((x & 0xcccccccc) >> 2)
	x = (x & 0x0f0f0f0f) + ((x & 0xf0f0f0f0) >> 4)
	x = (x & 0x00ff00ff) + ((x & 0xff00ff00) >> 8)
	x = (x & 0x0000ffff) + ((x & 0xffff0000) >> 16)
	return x
}

func CountBit1(in int) int {
	out := in & 1
	for i := 1; i < 32; i++ {
		if (in & (1 << i)) != 0 {
			out++
		}
	}
	return out
}

func CountBit1ChangeIn(in int) int {
	out := 0
	for ; in != 0; in >>= 1 {
		if (in & 1) != 0 {
			out++
		}
	}
	return out
}

func CountBit0(in int) int {
	out := 0
	for i := 0; i < 32; i++ {
		if (in & (1 << i)) == 0 {
			out++
		}
	}
	return out
}

func CountBit0ChangeIn(in int) int {
	out := 0
	for i := 0; i < 32; i++ {
		if (in & 1) == 0 {
			out++
		}
		in >>= 1
	}
	return out
}

func MarkNThBitTo1(in, n int) int {
	return in | (1 << (n - 1))
}

func ReverseNthBit(in, n int) int {
	return in ^ (1 << (n - 1))
}

func ClearNthBit(in, n int) int {
	return in & ^(1 << (n - 1))
}

func ReverseBitBuiltin(x uint32) uint32 {
	x = ((x >> 1) & 0x55555555) | ((x << 1) & 0xaaaaaaaa)
	x = ((x >> 2) & 0x33333333) | ((x << 2) & 0xcccccccc)
	x = ((x >> 4) & 0x0f0f0f0f) | ((x << 4) & 0xf0f0f0f0)
	x = ((x >> 8) & 0x00ff00ff) | ((x << 8) & 0xff00ff00)
	x = ((x >> 16) & 0x0000ffff) | ((x << 16) & 0xffff0000)
	return x
}

func LeastSignificantBit1(in int) int {
	return in & -in
}

func IsPowerOf2(in uint) bool {
	return (in & (in - 1)) == 0
}

// in = [0,1,2,4], out = 3
func FindLackNum(in []int) int {
	n := 0
	for i, v := range in {
		n = n ^ i ^ v
	}
	return n ^ len(in)
}

func ToLower(in byte) byte {
	if in >= 'A' && in <= 'Z' {
		return in ^ 32
	}
	return in
}

func ToUpper(in byte) byte {
	if in >= 'a' && in <= 'z' {
		return in ^ 32
	}
	return in
}
