# Quotient Filter

*   **Space and Time Complexity**:
    *   Space: `O(N * f)` bits, where `N` is capacity and `f` is fingerprint size.
    *   Time (Add, Contains, Delete): Amortized `O(1)` average, `O(N)` worst-case.

*   **Use Case**:
    *   **Synchronizing data between distributed databases**: By efficiently merging filters to determine missing keys.
    *   **Distributed Caching Systems**: Maintaining a distributed cache of frequently accessed items, allowing efficient synchronization of cache contents across nodes.
    *   **Network Flow Deduplication**: Deduplicating flow records across multiple collectors in network monitoring.
    *   **Blockchain Light Clients**: Representing a compact summary of transactions or states for efficient verification without downloading the entire blockchain.
    *   **Approximate Set Reconciliation**: For backend systems where two or more parties need to reconcile their sets of data efficiently without revealing all elements.

*   **Pros**:
    *   Often more space-efficient than Bloom and Cuckoo filters.
    *   Mergeable and resizable without rehashing original items.
    *   Good data locality for faster queries.
*   **Cons**:
    *   More complex to implement.
    *   Performance degrades near capacity.
