# Python program to find bridges in a given undirected graph
#Complexity : O(E*(V+E))

from collections import defaultdict

#This class represents an undirected graph using adjacency list representation
class Graph:
    
    def __init__(self,vertices):
        self.V= vertices #No. of vertices
        self.graph = defaultdict(dict) # default dictionary to store graph
        self.Time = 0

	# function to add an edge to graph
    def addEdge(self,u,v):
        self.graph[u][v] = True
        self.graph[v][u] = True
    

    def dfs(self, u, visited):
        for v in self.graph[u]:
            if v not in visited:
                visited.add(v)
                self.dfs(v, visited)


    def compute_group(self):
        visited = set()
        group = 0
        for u in range(self.V):
            if u not in visited:
                group += 1
                self.dfs(u, visited)
        return group

    def get_edges(self):
        edges = set()  # Use a set to avoid duplicate edges
        for u in self.graph:
            for v in self.graph[u]:
                edge = tuple(sorted([u, v]))  # Sort vertices to eliminate duplicates
                edges.add(edge)
        return list(edges)

    def get_bridge(self):
        bridges = []
        # get the disconnect subgraph count
        group = self.compute_group()
        # remove each edge, if subgraph count is increased, that means this edge is bridge
        for edge in self.get_edges():
            u, v = edge
            del self.graph[u][v]
            del self.graph[v][u]
            if self.compute_group() > group:
                bridges.append((u, v))

            self.addEdge(u, v)

        return bridges
    
    
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
