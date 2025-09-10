package singlesourceshortestpath

import "math"

/* @tags: graph,shortest path */

func DijkstraAlgorithm(graph [][]int, start, end int) int {
	n := len(graph)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt
	}
	parent := make([]int, n)
	visited := make([]bool, n)

	dis[start] = 0
	parent[start] = start

	for i := 0; i < n; i++ {
		a := -1
		min := math.MaxInt
		for j := 0; j < n; j++ {
			if !visited[j] && dis[j] < min {
				a = i
				min = dis[i]
			}
		}

		if a == -1 {
			break
		}

		visited[a] = true
		for b := 0; b < n; b++ {
			if !visited[b] && graph[a][b] != -1 && dis[a]+graph[a][b] < dis[b] {
				dis[b] = dis[a] + graph[a][b]
				parent[b] = a
			}
		}
	}

	return dis[end]
}
