package singlesourceshortestpath

import "math"

/* @tags: graph,shortest path */

func DijkstraAlgorithm(graph [][]int, start, end int) int {
	n := len(graph)

	dis := make([]int, n)
	for i := range n {
		dis[i] = math.MaxInt
	}

	parent := make([]int, n)
	visited := make([]bool, n)

	dis[start] = 0
	parent[start] = start

	for range n {
		a := -1

		minDistance := math.MaxInt
		for j := range n {
			if !visited[j] && dis[j] < minDistance {
				a = j
				minDistance = dis[j]
			}
		}

		if a == -1 {
			break
		}

		visited[a] = true
		for b := range n {
			if !visited[b] && graph[a][b] != -1 && dis[a]+graph[a][b] < dis[b] {
				dis[b] = dis[a] + graph[a][b]
				parent[b] = a
			}
		}
	}

	return dis[end]
}
