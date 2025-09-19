package unionfind

/* @tags: union find,group */

// UnionFind is a data structure that keeps track of a set of elements partitioned into a number of
// disjoint (non-overlapping) subsets.
type UnionFind[T comparable] struct {
	parents map[T]T
	sizes   map[T]int // use to flatten
}

// NewUnionFind creates and returns a new UnionFind instance.
func New[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parents: make(map[T]T),
		sizes:   make(map[T]int),
	}
}

// Find returns the representative (root) of the set that x belongs to.
func (uf *UnionFind[T]) Find(x T) T {
	// If x is not in the map, it's a new element, so set itself as the root.
	if _, ok := uf.parents[x]; !ok {
		uf.parents[x] = x
		uf.sizes[x] = 1
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
		if uf.sizes[rootX] < uf.sizes[rootY] {
			rootX, rootY = rootY, rootX
		}

		uf.parents[rootX] = rootY
		uf.sizes[rootX] += uf.sizes[rootY]
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
