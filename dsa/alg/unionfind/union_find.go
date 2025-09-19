package unionfind

/* @tags: union find,group */

// UnionFind is a data structure that keeps track of a set of elements partitioned into a number of disjoint (non-overlapping) subsets.
type UnionFind[T comparable] struct {
	parents map[T]T
}

// NewUnionFind creates and returns a new UnionFind instance.
func New[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parents: make(map[T]T),
	}
}

// Find returns the representative (root) of the set that x belongs to.
func (uf *UnionFind[T]) Find(x T) T {
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
func (uf *UnionFind[T]) Union(x, y T) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// If they are not already in the same set, merge them.
	if rootX != rootY {
		uf.parents[rootX] = rootY
	}
}

func (uf *UnionFind[T]) Groups() [][]T {
	groups := make(map[T][]T)

	for elem := range uf.parents {
		root := uf.Find(elem)
		groups[root] = append(groups[root], elem)
	}

	result := make([][]T, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
