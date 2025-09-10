from collections import defaultdict


def topoSortBfs(graph: dict) -> list[int]:
    # add indegree
    inDeg = { node: 0 for node in graph}
    for node in graph:
        for adj in graph[node]:
            inDeg[adj] += 1
    
    # bfs
    q = [node for node, d in inDeg.items() if d == 0]
    for u in q:
        for v in graph[u]:
            inDeg[v] -= 1
            if inDeg[v] == 0:
                q.append(v)

    if len(q) < len(graph):
        return None
    return q