package bitmap

// Bitmap is a bit array data structure.
type Bitmap struct {
	bits []uint64
	size uint
}

// NewBitmap creates a new Bitmap with the given size (number of bits).
func NewBitmap(size uint) *Bitmap {
	return &Bitmap{
		bits: make([]uint64, (size+63)/64), // Each uint64 holds 64 bits
		size: size,
	}
}

// Set sets the bit at the given position to 1.
func (bm *Bitmap) Set(pos uint) {
	if pos >= bm.size {
		return // Out of bounds
	}
	word := pos / 64
	bit := pos % 64
	bm.bits[word] |= (1 << bit)
}

// Clear sets the bit at the given position to 0.
func (bm *Bitmap) Clear(pos uint) {
	if pos >= bm.size {
		return // Out of bounds
	}
	word := pos / 64
	bit := pos % 64
	bm.bits[word] &^= (1 << bit)
}

// Test returns true if the bit at the given position is 1, false otherwise.
func (bm *Bitmap) Test(pos uint) bool {
	if pos >= bm.size {
		return false // Out of bounds
	}
	word := pos / 64
	bit := pos % 64
	return (bm.bits[word] & (1 << bit)) != 0
}

// Count returns the number of set bits in the bitmap.
func (bm *Bitmap) Count() uint {
	var count uint
	for _, word := range bm.bits {
		count += uint(popcount(word))
	}
	return count
}

// Size returns the total number of bits in the bitmap.
func (bm *Bitmap) Size() uint {
	return bm.size
}

// popcount counts the number of set bits (1s) in a uint64.
// This is a common algorithm, often implemented with hardware instructions.
func popcount(n uint64) int {
	count := 0
	for n > 0 {
		n &= (n - 1) // Clear the least significant set bit
		count++
	}
	return count
}