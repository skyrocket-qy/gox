package bfs

/* @tags: graph,bfs */

func BfsRecursive(graph map[int][]int, root int) []int {
	visited := map[int]bool{}
	visitSequence := []int{}

	var bfs func(vertices []int, visited map[int]bool, visitSequence *[]int, graph map[int][]int)
	bfs = func(vertices []int, visited map[int]bool, visitSequence *[]int, graph map[int][]int) {
		*visitSequence = append(*visitSequence, vertices...)
		next := []int{}

		for _, u := range vertices {
			visited[u] = true
			for _, v := range graph[u] {
				if _, ok := visited[v]; !ok {
					next = append(next, v)
				}
			}
		}

		if len(next) > 0 {
			bfs(next, visited, visitSequence, graph)
		}
	}

	bfs([]int{root}, visited, &visitSequence, graph)
	return visitSequence
}

func BfsIterative(graph map[int][]int, root int) []int {
	visited := map[int]bool{}
	visitSequence := []int{}

	queue := []int{root}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true
		for _, v := range graph[u] {
			if _, ok := visited[v]; !ok {
				queue = append(queue, v)
			}
		}

		visitSequence = append(visitSequence, u)
	}

	return visitSequence
}
