package findbridge

import (
	"math"
)

// Graph represents an undirected graph using adjacency list representation
type Graph struct {
	V       int           // No. of vertices
	Adj     [][]int       // Adjacency list
	Time    int           // Global time variable for discovery times
	Bridges [][]int       // Stores the found bridges
}

// NewGraph creates a new Graph instance
func NewGraph(vertices int) *Graph {
	adj := make([][]int, vertices)
	for i := range adj {
		adj[i] = []int{}
	}
	return &Graph{
		V:   vertices,
		Adj: adj,
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u)
}

// bridgeUtil is a recursive function that finds and prints bridges
// using DFS traversal
// u --> The vertex to be visited next
// visited[] --> keeps track of visited vertices
// disc[] --> Stores discovery times of visited vertices
// parent[] --> Stores parent vertices in DFS tree
// low[] --> Stores the lowest discovery time reachable from subtree rooted with current vertex
func (g *Graph) bridgeUtil(u int, visited []bool, parent []int, low []int, disc []int) {

	// Mark the current node as visited
	visited[u] = true

	// Initialize discovery time and low value
	disc[u] = g.Time
	low[u] = g.Time
	g.Time++

	// Recur for all the vertices adjacent to this vertex
	for _, v := range g.Adj[u] {
		// If v is not visited yet, then make it a child of u
		// in DFS tree and recur for it
		if !visited[v] {
			parent[v] = u
			g.bridgeUtil(v, visited, parent, low, disc)

			// Check if the subtree rooted with v has a connection to
			// one of the ancestors of u
			low[u] = int(math.Min(float64(low[u]), float64(low[v])))

			// If the lowest vertex reachable from subtree
			// under v is above u in DFS tree, then u-v is a bridge
			if low[v] > disc[u] {
				g.Bridges = append(g.Bridges, []int{u, v})
			}
		} else if v != parent[u] { // Update low value of u for parent function calls.
			low[u] = int(math.Min(float64(low[u]), float64(disc[v])))
		}
	}
}

// FindBridges DFS based function to find all bridges.
func (g *Graph) FindBridges() [][]int {
	g.Bridges = [][]int{}
	// Mark all the vertices as not visited and Initialize parent and visited,
	// and ap(articulation point) arrays
	visited := make([]bool, g.V)
	disc := make([]int, g.V)
	low := make([]int, g.V)
	parent := make([]int, g.V)

	for i := 0; i < g.V; i++ {
		disc[i] = math.MaxInt32
		low[i] = math.MaxInt32
		parent[i] = -1
	}

	// Call the recursive helper function to find bridges
	// in DFS tree rooted with vertex 'i'
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.bridgeUtil(i, visited, parent, low, disc)
		}
	}

	return g.Bridges
}
