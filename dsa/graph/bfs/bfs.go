package bfs

/* @tags: graph,bfs */

func BfsRecursive(graph map[int][]int, root int) []int {
	visited := map[int]bool{root: true}
	visitSequence := []int{}

	var bfs func(vertices []int, graph map[int][]int)

	bfs = func(vertices []int, graph map[int][]int) {
		if len(vertices) == 0 {
			return
		}
		visitSequence = append(visitSequence, vertices...)
		next := []int{}

		for _, u := range vertices {
			for _, v := range graph[u] {
				if _, ok := visited[v]; !ok {
					visited[v] = true
					next = append(next, v)
				}
			}
		}
		bfs(next, graph)
	}

	bfs([]int{root}, graph)

	return visitSequence
}

func BfsIterative(graph map[int][]int, root int) []int {
	visited := map[int]bool{root: true}
	visitSequence := []int{}

	queue := []int{root}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		visitSequence = append(visitSequence, u)

		for _, v := range graph[u] {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return visitSequence
}
