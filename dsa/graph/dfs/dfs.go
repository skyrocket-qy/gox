package dfs

/* @tags: graph,dfs */

func DfsRecursive(graph map[int][]int, root int) []int {
	visitSequence := []int{}
	visited := map[int]bool{}
	var dfs func(graph map[int][]int, visited map[int]bool, root int, visitSequence *[]int)
	dfs = func(graph map[int][]int, visited map[int]bool, root int, visitSequence *[]int) {
		*visitSequence = append(*visitSequence, root)
		visited[root] = true
		for _, v := range graph[root] {
			if _, ok := visited[v]; !ok {
				dfs(graph, visited, v, visitSequence)
			}
		}
	}

	dfs(graph, visited, root, &visitSequence)
	return visitSequence
}

func DfsIterative(graph map[int][]int, root int) []int {
	visitSequence := []int{}
	visited := map[int]bool{root: true}
	stk := []int{root}
	for len(stk) > 0 {
		cur := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		visitSequence = append(visitSequence, cur)
		visited[cur] = true
		for _, v := range graph[cur] {
			if _, ok := visited[v]; !ok {
				stk = append(stk, v)
				visited[v] = true
			}
		}
	}

	return visitSequence
}
