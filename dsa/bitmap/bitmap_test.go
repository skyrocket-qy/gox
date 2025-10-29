package bitmap

import (
	"testing"
)

func TestNewBitmap(t *testing.T) {
	bm := NewBitmap(100)
	if bm.size != 100 {
		t.Errorf("Expected size 100, got %d", bm.size)
	}
	if len(bm.bits) != (100+63)/64 {
		t.Errorf("Expected bits slice length %d, got %d", (100+63)/64, len(bm.bits))
	}

	bm = NewBitmap(0)
	if bm.size != 0 {
		t.Errorf("Expected size 0, got %d", bm.size)
	}
	if len(bm.bits) != 0 {
		t.Errorf("Expected bits slice length 0, got %d", len(bm.bits))
	}
}

func TestSetAndTest(t *testing.T) {
	bm := NewBitmap(128)

	bm.Set(0)
	if !bm.Test(0) {
		t.Errorf("Bit 0 should be set")
	}

	bm.Set(63)
	if !bm.Test(63) {
		t.Errorf("Bit 63 should be set")
	}

	bm.Set(64)
	if !bm.Test(64) {
		t.Errorf("Bit 64 should be set")
	}

	bm.Set(127)
	if !bm.Test(127) {
		t.Errorf("Bit 127 should be set")
	}

	// Test unset bits
	if bm.Test(1) {
		t.Errorf("Bit 1 should not be set")
	}

	// Test out of bounds
	if bm.Test(128) {
		t.Errorf("Bit 128 is out of bounds and should not be testable as true")
	}
	bm.Set(128) // Should not panic or change anything
}

func TestClear(t *testing.T) {
	bm := NewBitmap(64)
	bm.Set(10)
	bm.Set(20)

	if !bm.Test(10) {
		t.Errorf("Bit 10 should be set initially")
	}
	bm.Clear(10)
	if bm.Test(10) {
		t.Errorf("Bit 10 should be cleared")
	}

	// Clear an already clear bit
	bm.Clear(5)
	if bm.Test(5) {
		t.Errorf("Bit 5 should remain cleared")
	}

	// Clear out of bounds
	bm.Clear(100) // Should not panic
}

func TestCount(t *testing.T) {
	bm := NewBitmap(200)

	if bm.Count() != 0 {
		t.Errorf("Expected 0 set bits, got %d", bm.Count())
	}

	bm.Set(0)
	bm.Set(1)
	bm.Set(63)
	bm.Set(64)
	bm.Set(199)

	if bm.Count() != 5 {
		t.Errorf("Expected 5 set bits, got %d", bm.Count())
	}

	bm.Clear(0)
	if bm.Count() != 4 {
		t.Errorf("Expected 4 set bits after clearing, got %d", bm.Count())
	}

	bm.Set(0) // Set it back
	bm.Set(0) // Set it again, count should not change
	if bm.Count() != 5 {
		t.Errorf("Expected 5 set bits after setting again, got %d", bm.Count())
	}
}

func TestSize(t *testing.T) {
	bm := NewBitmap(123)
	if bm.Size() != 123 {
		t.Errorf("Expected size 123, got %d", bm.Size())
	}
}
