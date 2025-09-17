package unionfind_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/unionfind"
)

func TestFind(t *testing.T) {
	uf := unionfind.New()

	// Test find on a new element
	if uf.Find(1) != 1 {
		t.Errorf("Expected uf.Find(1) to be 1, but got %d", uf.Find(1))
	}

	// Test find after a union
	uf.Union(1, 2)

	if uf.Find(1) != uf.Find(2) {
		t.Errorf("Expected uf.Find(1) to be equal to uf.Find(2) after uf.Union(1, 2)")
	}

	// Test path compression indirectly
	uf.Union(2, 3)
	// Now 1, 2, 3 should be in the same set. uf.Find(1) should eventually point to the root of 3.
	if uf.Find(1) != uf.Find(3) {
		t.Errorf("Expected uf.Find(1) to be equal to uf.Find(3) after uf.Union(2, 3)")
	}
}

func TestUnion(t *testing.T) {
	uf := unionfind.New()

	uf.Union(1, 2)

	if uf.Find(1) != uf.Find(2) {
		t.Errorf("Expected uf.Find(1) to be equal to uf.Find(2) after uf.Union(1, 2)")
	}

	uf.Union(2, 3)

	if uf.Find(1) != uf.Find(3) {
		t.Errorf("Expected uf.Find(1) to be equal to uf.Find(3) after uf.Union(2, 3)")
	}

	uf.Union(4, 5)

	if uf.Find(1) == uf.Find(4) {
		t.Errorf("Expected uf.Find(1) to not be equal to uf.Find(4)")
	}

	uf.Union(1, 4)

	if uf.Find(1) != uf.Find(5) {
		t.Errorf("Expected uf.Find(1) to be equal to uf.Find(5) after uf.Union(1, 4)")
	}
}
