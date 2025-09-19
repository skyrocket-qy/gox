package unionfind_test

import (
	"cmp"
	"reflect"
	"sort"
	"testing"

	"github.com/skyrocket-qy/gox/dsa/alg/unionfind"
)

// sortAndCompareGroups is a helper function to sort and compare two slices of groups.
// It sorts the groups themselves and the elements within each group for consistent comparison.
func sortAndCompareGroups[T cmp.Ordered](t *testing.T, actual, expected [][]T) bool {
	t.Helper()
	// Sort elements within each group
	for _, group := range actual {
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})
	}

	for _, group := range expected {
		sort.Slice(group, func(i, j int) bool {
			return group[i] < group[j]
		})
	}

	// Sort the groups themselves
	sort.Slice(actual, func(i, j int) bool {
		// Compare groups based on their first element after internal sorting
		if len(actual[i]) == 0 || len(actual[j]) == 0 {
			return len(actual[i]) < len(actual[j])
		}

		return actual[i][0] < actual[j][0]
	})
	sort.Slice(expected, func(i, j int) bool {
		if len(expected[i]) == 0 || len(expected[j]) == 0 {
			return len(expected[i]) < len(expected[j])
		}

		return expected[i][0] < expected[j][0]
	})

	return reflect.DeepEqual(actual, expected)
}

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
		t.Errorf(
			"Expected uf.Find(\"a\") to be equal to uf.Find(\"b\") after uf.Union(\"a\", \"b\")",
		)
	}

	uf.Union("b", "c")

	if uf.Find("a") != uf.Find("c") {
		t.Errorf(
			"Expected uf.Find(\"a\") to be equal to uf.Find(\"c\") after uf.Union(\"b\", \"c\")",
		)
	}

	uf.Union("d", "e")

	if uf.Find("a") == uf.Find("d") {
		t.Errorf("Expected uf.Find(\"a\") to not be equal to uf.Find(\"d\")")
	}

	uf.Union("a", "d")

	if uf.Find("a") != uf.Find("e") {
		t.Errorf(
			"Expected uf.Find(\"a\") to be equal to uf.Find(\"e\") after uf.Union(\"a\", \"d\")",
		)
	}
}

func TestGroups(t *testing.T) {
	tests := []struct {
		name           string
		setup          func(uf *unionfind.UnionFind[int])
		expectedGroups [][]int
	}{
		{
			name:           "Empty UnionFind",
			setup:          func(uf *unionfind.UnionFind[int]) {},
			expectedGroups: [][]int{},
		},
		{
			name: "Single elements, no unions",
			setup: func(uf *unionfind.UnionFind[int]) {
				uf.Find(1)
				uf.Find(2)
				uf.Find(3)
			},
			expectedGroups: [][]int{{1}, {2}, {3}},
		},
		{
			name: "Two distinct groups",
			setup: func(uf *unionfind.UnionFind[int]) {
				uf.Union(1, 2)
				uf.Union(3, 4)
				uf.Find(5) // Element 5 is in its own group
			},
			expectedGroups: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "All elements in one group",
			setup: func(uf *unionfind.UnionFind[int]) {
				uf.Union(1, 2)
				uf.Union(2, 3)
				uf.Union(3, 4)
			},
			expectedGroups: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "Complex groups",
			setup: func(uf *unionfind.UnionFind[int]) {
				uf.Union(1, 2)
				uf.Union(3, 4)
				uf.Union(1, 3)
				uf.Union(5, 6)
				uf.Find(7)
			},
			expectedGroups: [][]int{{1, 2, 3, 4}, {5, 6}, {7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := unionfind.New[int]()
			tt.setup(uf)
			actualGroups := uf.Groups()

			if !sortAndCompareGroups(t, actualGroups, tt.expectedGroups) {
				t.Errorf(
					"Test %s failed: Expected %v, got %v",
					tt.name,
					tt.expectedGroups,
					actualGroups,
				)
			}
		})
	}

	// Test with string type
	stringTests := []struct {
		name           string
		setup          func(uf *unionfind.UnionFind[string])
		expectedGroups [][]string
	}{
		{
			name:           "String - Empty UnionFind",
			setup:          func(uf *unionfind.UnionFind[string]) {},
			expectedGroups: [][]string{},
		},
		{
			name: "String - Single elements, no unions",
			setup: func(uf *unionfind.UnionFind[string]) {
				uf.Find("a")
				uf.Find("b")
				uf.Find("c")
			},
			expectedGroups: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name: "String - Two distinct groups",
			setup: func(uf *unionfind.UnionFind[string]) {
				uf.Union("a", "b")
				uf.Union("c", "d")
				uf.Find("e")
			},
			expectedGroups: [][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			uf := unionfind.New[string]()
			tt.setup(uf)
			actualGroups := uf.Groups()

			if !sortAndCompareGroups(t, actualGroups, tt.expectedGroups) {
				t.Errorf(
					"Test %s failed: Expected %v, got %v",
					tt.name,
					tt.expectedGroups,
					actualGroups,
				)
			}
		})
	}
}
