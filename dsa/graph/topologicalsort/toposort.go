package topologicalsort

/* @tags: sort,graph,schedular */

/*
type node int
type edge int

The representation of graph: map[node][]edge
*/

// T: O(V + E)
// S: O(V).
func TopoSortRemoveVertix(graph map[int][]int) (seq []int) {
	inDeg := make(map[int]int)
	for node := range graph {
		inDeg[node] = 0
	}
	for _, edges := range graph {
		for _, edge := range edges {
			inDeg[edge]++
		}
	}

	queue := []int{}
	for node, degree := range inDeg {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for _, edge := range graph[node] {
			inDeg[edge]--
			if inDeg[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}

	if len(queue) < len(graph) {
		return nil
	}

	return queue
}

const (
	unvisited = 0
	visiting  = 1
	visited   = 2
)

func TopoSortDfs(graph map[int][]int) []int {
	states := make(map[int]int)
	for node := range graph {
		states[node] = unvisited
	}

	var result []int
	var hasCycle bool

	var dfs func(node int)
	dfs = func(node int) {
		states[node] = visiting
		for _, neighbor := range graph[node] {
			if states[neighbor] == visiting {
				hasCycle = true
				return
			}
			if states[neighbor] == unvisited {
				dfs(neighbor)
			}
			if hasCycle {
				return
			}
		}
		states[node] = visited
		result = append([]int{node}, result...)
	}

	for node := range graph {
		if states[node] == unvisited {
			dfs(node)
		}
		if hasCycle {
			return nil
		}
	}

	return result
}
