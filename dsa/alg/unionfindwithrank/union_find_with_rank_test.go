package unionfindwithrank

import (
	"testing"
)

func setup() {
	parents = make(map[int]int)
	ranks = make(map[int]int)
}

func TestFindWithRank(t *testing.T) {
	setup()

	// Test find on a new element
	if find(1) != 1 {
		t.Errorf("Expected find(1) to be 1, but got %d", find(1))
	}

	// Test find after a union
	parents[1] = 2
	if find(1) != 2 {
		t.Errorf("Expected find(1) to be 2, but got %d", find(1))
	}
}

func TestUnionWithRank(t *testing.T) {
	setup()

	initRanks([]int{1, 2, 3, 4, 5})

	union(1, 2)
	if find(1) != find(2) {
		t.Errorf("Expected find(1) to be equal to find(2) after union(1, 2)")
	}

	union(2, 3)
	if find(1) != find(3) {
		t.Errorf("Expected find(1) to be equal to find(3) after union(2, 3)")
	}

	union(4, 5)
	if find(1) == find(4) {
		t.Errorf("Expected find(1) to not be equal to find(4)")
	}

	union(1, 4)
	if find(1) != find(5) {
		t.Errorf("Expected find(1) to be equal to find(5) after union(1, 4)")
	}

	// Test rank updates
	setup()
	initRanks([]int{1, 2, 3, 4})
	union(1, 2) // rank of root of {1, 2} is 1
	union(3, 4) // rank of root of {3, 4} is 1
	union(1, 3) // rank of root of {1, 2, 3, 4} is 2

	root := find(1)
	if ranks[root] != 2 {
		t.Errorf("Expected rank of the new root to be 2, but got %d", ranks[root])
	}
}
