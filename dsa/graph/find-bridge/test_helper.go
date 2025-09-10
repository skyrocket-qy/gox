package findbridge

import (
	"sort"
)

// Helper function to sort bridges for consistent comparison.
func sortBridges(bridges [][]int) {
	sort.Slice(bridges, func(i, j int) bool {
		if bridges[i][0] != bridges[j][0] {
			return bridges[i][0] < bridges[j][0]
		}

		return bridges[i][1] < bridges[j][1]
	})

	for i := range bridges {
		if bridges[i][0] > bridges[i][1] {
			bridges[i][0], bridges[i][1] = bridges[i][1], bridges[i][0]
		}
	}
}
