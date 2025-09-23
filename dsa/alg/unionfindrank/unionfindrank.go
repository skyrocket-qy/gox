package unionfindrank

import "cmp"

type UnionFind[T cmp.Ordered] struct {
	parents  map[T]T
	isParent func(T, T) bool
}

// NewUnionFind creates and returns a new UnionFind instance.
func New[T cmp.Ordered](isParent func(T, T) bool) *UnionFind[T] {
	return &UnionFind[T]{
		parents:  make(map[T]T),
		isParent: isParent,
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
func (uf *UnionFind[T]) Union(x, y T) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// If they are not already in the same set, merge them.
	if rootX != rootY {
		if !uf.isParent(rootX, rootY) {
			rootX, rootY = rootY, rootX
		}

		uf.parents[rootY] = rootX

		return true
	}

	return false
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
