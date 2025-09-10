# Python program to find bridges in a given undirected graph
#Complexity : O(V+E)

from collections import defaultdict

#This class represents an undirected graph using adjacency list representation
class Graph:

	def __init__(self,vertices):
		self.V= vertices #No. of vertices
		self.graph = defaultdict(list) # default dictionary to store graph
		self.Time = 0

	# function to add an edge to graph
	def addEdge(self,u,v):
		self.graph[u].append(v)
		self.graph[v].append(u)

	'''A recursive function that finds and prints bridges
	using DFS traversal
	u --> The vertex to be visited next
	visited[] --> keeps track of visited vertices
	disc[] --> Stores discovery times of visited vertices
	parent[] --> Stores parent vertices in DFS tree'''
	def bridgeUtil(self,u, visited, parent, low, disc):

		# Mark the current node as visited
		visited[u]= True

		# Initialize discovery time and low value
		disc[u] = self.Time
		low[u] = self.Time
		self.Time += 1

		#Recur for all the vertices adjacent to this vertex
		for v in self.graph[u]:
			# If v is not visited yet, then make it a child of u
			# in DFS tree and recur for it
			if visited[v] == False:
				parent[v] = u
				self.bridgeUtil(v, visited, parent, low, disc)

				# Check if the subtree rooted with v has a connection to
				# one of the ancestors of u
				low[u] = min(low[u], low[v])


				''' If the lowest vertex reachable from subtree
				under v is above u in DFS tree, then u-v is
				a bridge'''
				if low[v] > disc[u]:
					self.bridges.append([u, v])
	
					
			elif v != parent[u]: # Update low value of u for parent function calls.
				low[u] = min(low[u], disc[v])

	# DFS based function to find all bridges. It uses recursive
	# function bridgeUtil()
	def get_bridge(self):
		self.bridges = []
		# Mark all the vertices as not visited and Initialize parent and visited,
		# and ap(articulation point) arrays
		visited = [False] * self.V
		disc = [float("Inf")] * self.V
		low = [float("Inf")] * self.V
		parent = [-1] * self.V

		# Call the recursive helper function to find bridges
		# in DFS tree rooted with vertex 'i'
		for i in range(self.V):
			if visited[i] == False:
				self.bridgeUtil(i, visited, parent, low, disc)

		return self.bridges
		

# Create a graph given in the above diagram
g1 = Graph(5)
g1.addEdge(1, 0)
g1.addEdge(0, 2)
g1.addEdge(2, 1)
g1.addEdge(0, 3)
g1.addEdge(3, 4)


print ("Bridges in first graph ")
print(g1.get_bridge())


#This code is contributed by Neelam Yadav
