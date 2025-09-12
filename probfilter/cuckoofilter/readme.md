# Cuckoo Filter

*   **Space and Time Complexity**:
    *   Space: `O(N * f)` bits, where `N` is capacity and `f` is fingerprint size.
    *   Time (Add, Contains, Delete): Amortized `O(1)` average.

*   **Use Case**: Filtering for recently accessed articles in a CDN cache, supporting dynamic addition and deletion of items.

*   **Pros**:
    *   Supports dynamic addition and deletion of items.
    *   Often higher space efficiency than Bloom filters for low false positive rates.
*   **Cons**:
    *   Insertions can fail if filter is too full.
    *   Slightly more complex implementation.
