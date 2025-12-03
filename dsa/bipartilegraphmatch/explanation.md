# Hopcroft-Karp Algorithm Explanation

The Hopcroft-Karp algorithm is an efficient method for finding the maximum matching in a bipartite graph. It runs in $O(E \sqrt{V})$ time, which is faster than the $O(VE)$ time of simple DFS/BFS based approaches (like Kuhn's algorithm) or general max flow algorithms (like Ford-Fulkerson).

## Core Concepts

The algorithm works in **phases**. Each phase consists of two main steps:

1.  **BFS (Breadth-First Search) - Layering**:
    *   It searches for the *shortest* augmenting paths from free (unmatched) workers to free jobs.
    *   It partitions the graph into "layers" based on distance from free workers.
    *   It finds the length of the shortest augmenting path. If no such path exists, the algorithm terminates.

2.  **DFS (Depth-First Search) - Augmenting**:
    *   It looks for vertex-disjoint augmenting paths of the length found by BFS.
    *   It updates the matching by flipping edges along these paths (free edges become matched, matched edges become free).
    *   It ensures that we find a *maximal set* of shortest augmenting paths in this phase.

## Visualizing a Phase

Imagine a graph with Workers (U) on the left and Jobs (V) on the right.

### Step 1: BFS (Build Layers)
We start BFS from all **unmatched workers**.
*   **Layer 0**: Unmatched Workers.
*   **Layer 1**: Jobs connected to Layer 0.
*   **Layer 2**: Workers matched to Jobs in Layer 1.
*   **Layer 3**: Jobs connected to Layer 2.
*   ...
*   We stop when we reach a layer containing **unmatched jobs**. Let's say this is Layer $k$.
*   We only care about paths of length $k$.

### Step 2: DFS (Find Paths)
We run DFS from unmatched workers to find paths to unmatched jobs, strictly following the layers ($L_0 \to L_1 \to L_2 \dots$).
*   If we find a path, we "augment" it (swap matched/unmatched edges) and remove the vertices involved from consideration for this phase.
*   We repeat until no more paths of length $k$ exist.

## Why is it fast?
By finding *multiple* shortest paths in one phase and augmenting them simultaneously, the length of the shortest augmenting path strictly increases with each phase. It can be shown that there are at most $2\sqrt{V}$ phases.

## Example Trace

**Graph**:
*   Workers: A, B
*   Jobs: 1, 2
*   Edges: (A, 1), (A, 2), (B, 1)

**Initial State**: Matching = {}

**Phase 1**:
1.  **BFS**:
    *   Start at free workers {A, B} (Dist 0).
    *   Neighbors of A: {1, 2}. Neighbors of B: {1}.
    *   Jobs 1 and 2 are free. We found shortest paths of length 1.
    *   Distances: A=0, B=0, 1=1, 2=1.
2.  **DFS**:
    *   Try A: A -> 1. 1 is free. **Match (A, 1)**.
    *   Try B: B -> 1. 1 is now matched (to A).
        *   Can we go deeper? 1 is matched to A. Next layer would be A.
        *   But A is already visited/used in this DFS phase?
        *   Actually, in HK, we look for *vertex-disjoint* paths. Since 1 is used by A, B cannot use 1.
        *   Does B have other neighbors? No.
    *   End of Phase 1. Matching: {(A, 1)}.

**Phase 2**:
1.  **BFS**:
    *   Free workers: {B}.
    *   Start BFS from B (Dist 0).
    *   Neighbors of B: {1}.
    *   1 is matched to A. So we go to A (Dist 2).
    *   Neighbors of A: {1, 2}.
    *   1 is visited. 2 is free!
    *   Path found: B -> 1 -> A -> 2. Length 3.
2.  **DFS**:
    *   Try B: B -> 1 -> A -> 2.
    *   2 is free. **Augment**:
        *   Add (B, 1), Remove (A, 1), Add (A, 2).
    *   New Matching: {(B, 1), (A, 2)}.

**Phase 3**:
*   BFS finds no augmenting paths. Done.
*   Max Matching: 2.
