# Concurrent DSA

Here are the implementations under dsa/ that would be valuable to implement in a concurrent version:

  Algorithms (`dsa/alg/`)
   * getallsubsequence: Generating subsequences can be parallelized.
   * kthselect: Parallel quickselect approaches can be used.
   * longestcommonsubsequence: Dynamic programming table filling can be parallelized.
   * permutation: Generating permutations can be parallelized.
   * sort: Many sorting algorithms have parallel counterparts (e.g., parallel merge sort, parallel quicksort).
   * stringmatch: Can be parallelized for multiple patterns or large texts.
   * unionfind / unionfindwithrank: Concurrent versions of Union-Find structures exist.

  Data Structures (`dsa/ds/`)
   * binary-search-tree: Concurrent BSTs (e.g., lock-free or fine-grained locking) are a common and valuable area.
   * circular-queue, queue, stack: Concurrent queues and stacks are very common and useful (e.g., lock-free queues, concurrent
     stacks).
   * heap, priority-queue: Concurrent heaps/priority queues are valuable for parallel processing.
   * lfu-cache, lru-cache: Caches are frequently accessed concurrently, making concurrent implementations highly valuable.
   * ordered-set: Concurrent ordered sets (e.g., using skip lists or concurrent trees) are beneficial.
   * prefix-tree (Trie), radix-tree: Concurrent tries/radix trees can improve performance for search and update operations.
   * randomized-set: Concurrent versions are possible, though potentially complex.

  Graphs (`dsa/graph/`)
   * a_star, bfs, dfs: Graph traversal algorithms can often be parallelized for large graphs.
   * find-bridge: Can be parallelized.
   * mst (Minimum Spanning Tree): Parallel MST algorithms exist (e.g., Boruvka's algorithm).
   * single-source-shortest-path: Parallel SSSP algorithms exist (e.g., parallel Bellman-Ford, parallel Dijkstra for certain
     graph types).
   * topologicalsort: Can be parallelized.