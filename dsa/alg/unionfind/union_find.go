package unionfind

/* @tags: union find,group */

// UnionFind is a data structure that keeps track of a set of elements partitioned into a number of disjoint (non-overlapping) subsets.
type UnionFind struct {
	parents map[int]int
}

// NewUnionFind creates and returns a new UnionFind instance.
func New() *UnionFind {
	return &UnionFind{
		parents: make(map[int]int),
	}
}

// Find returns the representative (root) of the set that x belongs to.
func (uf *UnionFind) Find(x int) int {
	// If x is not in the map, it's a new element, so set itself as the root.
	if _, ok := uf.parents[x]; !ok {
		uf.parents[x] = x
	}

	// Path compression: If x is not the root, recursively find its root and set it as x's parent.
	if x != uf.parents[x] {
		uf.parents[x] = uf.Find(uf.parents[x])
	}

	return uf.parents[x]
}

// Union merges the sets containing x and y.
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// If they are not already in the same set, merge them.
	if rootX != rootY {
		uf.parents[rootX] = rootY
	}
}
