# Counting Bloom Filter

*   **Space and Time Complexity**:
    *   Space: `O(m * c)` bits, where `m` is number of counters and `c` is bits per counter.
    *   Time (Add, Contains, Remove): `O(k)` operations.

*   **Use Case**: Managing a list of malicious URLs where URLs need to be added and removed dynamically.

*   **Pros**:
    *   Supports deletion of elements.
    *   Relatively space-efficient.
*   **Cons**:
    *   Requires more space than standard Bloom filter.
    *   Counters can overflow.
    *   "Soft" deletions can introduce errors.
