# XOR Filter

### Explanation

An XOR filter is a type of probabilistic data structure used for set membership testing. They are faster and more memory-efficient than Bloom and Cuckoo filters. XOR filters are a type of perfect hash function, meaning they have no false positives for items in the set they were built with. However, they are static; once constructed, items cannot be added or removed.

### Scenario: Serving static assets from a CDN

A Content Delivery Network (CDN) needs to determine quickly if an asset (e.g., an image or a JavaScript file) is stored at an edge location. The set of assets is large but changes infrequently.

An XOR filter can be built containing the fingerprints of all assets. When a request for an asset comes in, the CDN can use the XOR filter to very quickly check if the asset is supposed to be in the cache. Because XOR filters are extremely fast and small, this check is highly efficient. When the set of assets is updated, a new filter must be constructed and distributed to the edge locations.

### Comparison

*   **Pros**:
    *   Uses less space than Bloom or Cuckoo filters (e.g., ~1.23 bits per item).
    *   Faster lookups than Bloom or Cuckoo filters.
    *   No false positives for items that are in the set.
*   **Cons**:
    *   It is a static data structure; it cannot be modified after it is built. Adding or removing items requires a complete rebuild of the filter.
    *   The construction process is more complex than for a Bloom filter.

### Mathematical Foundations

XOR filters are a type of perfect hash function, meaning they guarantee no false positives for items that were part of the set used to construct the filter. They achieve this by mapping each item to three positions in an array and storing values such that the XOR sum of the values at these three positions equals a fingerprint of the item. The construction process involves solving a system of linear equations over GF(2) (Galois Field of 2 elements), often using a peeling algorithm on a bipartite graph.

The space efficiency of XOR filters is remarkable, typically requiring only about 1.23 bits per item, making them more compact than Bloom or Cuckoo filters for many use cases.

### Implementation Considerations

*   **Static Nature**: The most significant consideration is that XOR filters are static. Once constructed, items cannot be added or removed. Any change to the set of items requires a complete rebuild of the filter, which can be computationally intensive.
*   **Construction Algorithm**: The construction algorithm is the most complex part of an XOR filter. It typically involves:
    1.  Mapping each item to three candidate positions in the filter array using three hash functions.
    2.  Building a bipartite graph where one set of nodes represents items and the other represents array positions.
    3.  Using a "peeling" algorithm to find a unique assignment for each item, ensuring that each item can be uniquely identified by XORing the values at its three positions.
    4.  Populating the filter array with values such that the XOR sum property holds.
*   **Hash Functions**: Three independent hash functions are crucial for the construction and lookup processes.
*   **Memory Efficiency**: Despite the complex construction, XOR filters offer superior memory efficiency and faster lookups compared to other probabilistic filters once built.

### Performance Analysis

*   **Space Complexity**: `O(N)` bits, typically around 1.23 bits per item, making them extremely space-efficient.
*   **Time Complexity**:
    *   **Construction**: `O(N)` on average, but can be complex and computationally intensive, especially for very large `N`.
    *   **Contains**: `O(1)` (constant time) with very few memory accesses, making them extremely fast.
    *   **Add/Delete**: Not supported after construction.
*   **Practical Performance**: Once constructed, XOR filters offer unparalleled lookup speed and memory efficiency for static sets.

### Trade-offs

*   **Static Nature**: The most significant limitation is that XOR filters are static. They cannot be updated (additions or deletions) after construction. Any change to the dataset requires rebuilding the entire filter, which can be costly.
*   **Construction Complexity**: The construction algorithm is more complex than for Bloom or Cuckoo filters, involving graph algorithms and solving systems of equations.
*   **No False Positives (for items in the set)**: A major advantage is that for items that were part of the original set, there are no false positives. This makes them "perfect" for membership testing of the original set.
*   **Space and Speed**: They offer the best combination of space efficiency and lookup speed among probabilistic membership filters for static sets.

## Code Example

A basic Go implementation of the XOR Filter can be found [here](code/xor_filter.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd xor-filter/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go xor_filter.go
    ```