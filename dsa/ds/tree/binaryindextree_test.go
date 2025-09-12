package tree

import "testing"

func TestBIT(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	bit := Build(array)
	if sum := bit.Query(0, 4); sum != 15 {
		t.Errorf("expected sum to be 15, but got %d", sum)
	}
	if sum := bit.Query(0, 2); sum != 6 {
		t.Errorf("expected sum to be 6, but got %d", sum)
	}
	bit.Update(2, 3)
	if sum := bit.Query(0, 2); sum != 9 {
		t.Errorf("expected sum to be 9, but got %d", sum)
	}
}
