# XOR Filter

*   **Space and Time Complexity**:
    *   Space: `O(N)` bits, ~1.23 bits per item.
    *   Time (Construction): `O(N)` average.
    *   Time (Contains): `O(1)`.
    *   Add/Delete: Not supported.

*   **Use Case**: Serving static assets from a CDN where the set of assets is large but changes infrequently.

*   **Pros**:
    *   Less space than Bloom or Cuckoo filters.
    *   Faster lookups than Bloom or Cuckoo filters.
    *   No false positives for items in the set.
*   **Cons**:
    *   Static data structure; requires complete rebuild for modifications.
    *   More complex construction.
