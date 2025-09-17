package unionfind_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/unionfind"
)

func TestFind(t *testing.T) {
	uf := unionfind.New[int]()

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
	uf := unionfind.New[int]()

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

func TestUnionFindString(t *testing.T) {
	uf := unionfind.New[string]()

	uf.Union("a", "b")
	if uf.Find("a") != uf.Find("b") {
		t.Errorf("Expected uf.Find(\"a\") to be equal to uf.Find(\"b\") after uf.Union(\"a\", \"b\")")
	}

	uf.Union("b", "c")
	if uf.Find("a") != uf.Find("c") {
		t.Errorf("Expected uf.Find(\"a\") to be equal to uf.Find(\"c\") after uf.Union(\"b\", \"c\")")
	}

	uf.Union("d", "e")
	if uf.Find("a") == uf.Find("d") {
		t.Errorf("Expected uf.Find(\"a\") to not be equal to uf.Find(\"d\")")
	}

	uf.Union("a", "d")
	if uf.Find("a") != uf.Find("e") {
		t.Errorf("Expected uf.Find(\"a\") to be equal to uf.Find(\"e\") after uf.Union(\"a\", \"d\")")
	}
}