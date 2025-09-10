package unionfind

/* @tags: union find,group */

// UF is a hash map where you can find the root of a group of elements giving an element.
// A key in UF is a element, UF[x] is x's parent.
// If UF[x] == x meaning x is the root of its group.
var parents map[int]int

// Given an element, find the root of the group to which this element belongs.
func find(x int) int {
	// this may be the first time we see x, so set itself as the root.
	if _, ok := parents[x]; !ok {
		parents[x] = x
	}
	// If x == UF[x], meaning x is the root of this group.
	// If x != UF[x], we use the find function again on x's parent UF[x]
	// until we find the root and set it as the parent (value) of x in UF.
	if x != parents[x] {
		parents[x] = find(parents[x])
	}
	return parents[x]
}

// Given two elements x and y, we know that x and y should be in the same group,
// this means the group that contains x and the group that contains y
// should be merged together if they are currently separate groups.
// So we first find the root of x and the root of y using the find function.
// We then set the root of y (rootY) as the root of the root of x (rootX).
func union(x, y int) {
	rootX := find(x)
	rootY := find(y)
	// set the root of y (rootY) as the root of the root of x (rootX)
	parents[rootX] = rootY
}
