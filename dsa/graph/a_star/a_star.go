package astar

/* @tags: a star,graph,heuristic */

import (
	"container/heap"
	"math"
)

// Point represents a coordinate in the grid.
type Point struct {
	X, Y int
}

// Node represents a node in the graph for A* search.
type Node struct {
	Point             // Coordinates of the node
	Cost      float64 // Cost from start to this node (g-score)
	Heuristic float64 // Estimated cost from this node to goal (h-score)
	Parent    *Node   // Parent node to reconstruct path
	// The Index field is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// FScore returns the total estimated cost (f-score = g + h).
func (n *Node) FScore() float64 {
	return n.Cost + n.Heuristic
}

// PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest FScore
	return pq[i].FScore() < pq[j].FScore()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	old := *pq
	n := len(old)
	node := x.(*Node)
	node.Index = n
	*pq = append(old, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.Index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// AStar finds the shortest path between a start and goal node using the A* algorithm.
//
// grid: A 2D array representing the graph, where 0 is traversable and 1 is an obstacle.
// startX, startY: Coordinates of the starting node.
// goalX, goalY: Coordinates of the goal node.
// heuristicFn: A function that calculates the heuristic (estimated cost to goal).
//
// Returns the path as a slice of Nodes (from start to goal) and the total cost, or nil and 0 if no path is found.
func AStar(grid [][]int, startX, startY, goalX, goalY int, heuristicFn func(x1, y1, x2, y2 int) float64) ([]*Node, float64) {
	rows := len(grid)
	if rows == 0 {
		return nil, 0
	}
	cols := len(grid[0])

	// Check if start and goal are within bounds and not obstacles
	if startX < 0 || startX >= rows || startY < 0 || startY >= cols || grid[startX][startY] == 1 {
		return nil, 0
	}
	if goalX < 0 || goalX >= rows || goalY < 0 || goalY >= cols || grid[goalX][goalY] == 1 {
		return nil, 0
	}

	startPoint := Point{X: startX, Y: startY}
	goalPoint := Point{X: goalX, Y: goalY}

	startNode := &Node{Point: startPoint, Cost: 0, Heuristic: heuristicFn(startX, startY, goalX, goalY)}

	openSet := make(PriorityQueue, 0)
	heap.Push(&openSet, startNode)

	openSetNodes := make(map[Point]*Node) // Map to quickly check if a node is in openSet
	openSetNodes[startPoint] = startNode

	// cameFrom maps a node's Point to its parent Node.
	cameFrom := make(map[Point]*Node)

	// gScore maps a node's Point to the cost of the cheapest path from start to that node found so far.
	gScore := make(map[Point]float64)
	gScore[startPoint] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Node)
		delete(openSetNodes, current.Point)

		if current.Point == goalPoint {
			// Reconstruct path
			path := make([]*Node, 0)
			for n := current; n != nil; n = cameFrom[n.Point] {
				path = append([]*Node{n}, path...)
			}
			return path, current.Cost
		}

		// Define possible movements (8-directional movement)
		dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
		dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

		for i := 0; i < len(dx); i++ {
			neighborX, neighborY := current.X+dx[i], current.Y+dy[i]
			neighborPoint := Point{X: neighborX, Y: neighborY}

			if neighborX >= 0 && neighborX < rows && neighborY >= 0 && neighborY < cols && grid[neighborX][neighborY] == 0 {
				// Calculate tentative gScore for neighbor
				moveCost := 1.0
				if dx[i] != 0 && dy[i] != 0 { // Diagonal move
					moveCost = math.Sqrt(2)
				}
				tentativeGScore := current.Cost + moveCost

				if existingGScore, ok := gScore[neighborPoint]; !ok || tentativeGScore < existingGScore {
					cameFrom[neighborPoint] = current
					gScore[neighborPoint] = tentativeGScore

					neighborNode, foundInOpenSet := openSetNodes[neighborPoint]
					if foundInOpenSet {
						neighborNode.Cost = tentativeGScore
						neighborNode.Heuristic = heuristicFn(neighborX, neighborY, goalX, goalY)
						heap.Fix(&openSet, neighborNode.Index)
					} else {
						newNode := &Node{
							Point:     neighborPoint,
							Cost:      tentativeGScore,
							Heuristic: heuristicFn(neighborX, neighborY, goalX, goalY),
							Parent:    current, // Set parent here for path reconstruction
						}
						heap.Push(&openSet, newNode)
						openSetNodes[neighborPoint] = newNode
					}
				}
			}
		}
	}

	return nil, 0 // No path found
}

// ManhattanDistance is a common heuristic for grid-based paths (4-directional movement).
func ManhattanDistance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
}

// EuclideanDistance is a common heuristic for grid-based paths (8-directional movement).
func EuclideanDistance(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}
