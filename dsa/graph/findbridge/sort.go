package findbridge

import (
	"sort"
)

func SortBridges(bridges [][]int) {
	for _, bridge := range bridges {
		sort.Ints(bridge)
	}

	sort.Slice(bridges, func(i, j int) bool {
		if bridges[i][0] != bridges[j][0] {
			return bridges[i][0] < bridges[j][0]
		}

		return bridges[i][1] < bridges[j][1]
	})
}
