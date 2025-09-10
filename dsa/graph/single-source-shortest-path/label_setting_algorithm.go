package singlesourceshortestpath

/* @tags: graph,shortest path */

import "math"

func LabelSettingAlgorithm(graph [][]int, start, end int) int {
	n := len(graph)

	dis := make([]int, n)
	for i := range n {
		dis[i] = math.MaxInt
	}

	parent := make([]int, n)
	visited := make([]bool, n)

	dis[start] = 0
	// use to find path
	parent[start] = start
	visited[start] = true

	// Add all vertices
	for range n - 1 {
		a, b := -1, -1
		min := math.MaxInt

		for j := range n {
			if visited[j] {
				for k := range n {
					if !visited[k] && graph[j][k] != -1 {
						if dis[j]+graph[j][k] < min {
							a, b = j, k
							min = dis[j] + graph[j][k]
						}
					}
				}
			}
		}

		// all connected vertices visited
		if a == -1 || b == -1 {
			break
		}

		dis[b] = min
		parent[b] = a
		visited[b] = true
	}

	return dis[end]
}
