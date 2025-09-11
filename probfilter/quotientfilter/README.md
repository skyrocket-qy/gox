# Quotient Filter

*   **Space and Time Complexity**:
    *   Space: `O(N * f)` bits, where `N` is capacity and `f` is fingerprint size.
    *   Time (Add, Contains, Delete): Amortized `O(1)` average, `O(N)` worst-case.

*   **Use Case**: Synchronizing data between distributed databases by efficiently merging filters to determine missing keys.

*   **Pros**:
    *   Often more space-efficient than Bloom and Cuckoo filters.
    *   Mergeable and resizable without rehashing original items.
    *   Good data locality for faster queries.
*   **Cons**:
    *   More complex to implement.
    *   Performance degrades near capacity.
