package unionfindwithrank

/* @tags: union find,group */

var parents map[int]int
var ranks map[int]int

func initRanks(eles []int) {
	for _, ele := range eles {
		ranks[ele] = 0
	}
}

func find(x int) int {
	if _, ok := parents[x]; !ok {
		parents[x] = x
	}

	if x != parents[x] {
		parents[x] = find(parents[x])
	}
	return parents[x]
}

func union(x, y int) {
	rootX := find(x)
	rootY := find(y)

	if ranks[rootX] < ranks[rootY] {
		parents[rootX] = rootY
	} else {
		parents[rootY] = rootX
		if ranks[rootX] == ranks[rootY] {
			ranks[rootX]++
		}
	}
}
