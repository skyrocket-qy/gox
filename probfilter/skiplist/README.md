# Skip List

*   **Space and Time Complexity**:
    *   Space: `O(N)` average.
    *   Time (Search, Insert, Delete): `O(log N)` average.

*   **Use Case**:
    *   **High-performance database indexing**: For fast lookups, insertions, and deletions, especially in concurrent environments.
    *   **In-memory Sorted Sets/Leaderboards**: For backend services that need to maintain sorted lists with frequent updates and range queries.
    *   **Concurrent Priority Queues**: Implementing high-performance, concurrent priority queues in backend systems.
    *   **Distributed System Coordination**: Managing ordered lists of nodes or tasks in distributed systems with dynamic membership and concurrent access.
    *   **Range Queries in Key-Value Stores**: Efficiently supporting range queries on sorted data in backend key-value stores.

*   **Pros**:
    *   Simpler to implement than balanced trees.
    *   Good average performance for search, insertion, and deletion.
    *   Well-suited for concurrent applications.
*   **Cons**:
    *   Uses more memory.
    *   Performance is probabilistic, not guaranteed.
