# Cuckoo Filter

### Explanation

A Cuckoo filter is another probabilistic data structure for set membership testing, similar to a Bloom filter. It is based on Cuckoo Hashing and, most notably, supports dynamic addition and deletion of items. This makes it a powerful alternative to Bloom filters in scenarios where the set of items changes frequently.

### Scenario: Filtering for recently accessed articles

Consider a content delivery network (CDN) that caches recently accessed articles. To avoid a slow lookup in the main storage (e.g., a database or object storage), the CDN can use a Cuckoo filter to keep track of which articles are currently in its cache.

1.  When an article is added to the cache, its ID is added to the Cuckoo filter.
2.  When a request for an article comes in, the CDN first checks the filter. If the filter says the article is not in the cache, the request is forwarded to the main storage.
3.  If the filter indicates the article *might* be in the cache, the CDN attempts to retrieve it from the cache.
4.  Crucially, when an article is removed from the cache (e.g., due to an eviction policy), its ID is also removed from the Cuckoo filter. This keeps the filter accurate over time.

### Comparison

*   **Pros**:
    *   Supports dynamic addition and **deletion** of items.
    *   Often has higher space efficiency (fewer bits per item) than Bloom filters for low false positive rates (e.g., < 3%).
*   **Cons**:
    *   Insertions can fail if the filter is too full, requiring the filter to be rebuilt with more space.
    *   Implementation is slightly more complex than a standard Bloom filter.

### Mathematical Foundations

Cuckoo filters are based on Cuckoo Hashing, where each item has a few possible locations (buckets) in the hash table. When an item is inserted, it tries to occupy one of its designated locations. If all locations are occupied, it "kicks out" an existing item, which then tries to find a new home in one of its alternative locations. This process can cascade, leading to a chain of displacements. If a cycle is detected or a maximum number of displacements is reached, the insertion fails, and the filter might need to be resized.

The false positive rate of a Cuckoo filter is primarily determined by the size of the fingerprints and the number of entries per bucket. Larger fingerprints and more entries per bucket generally lead to lower false positive rates.

### Implementation Considerations

*   **Cuckoo Hashing**: The core challenge is implementing the cuckoo hashing mechanism, including handling displacements and potential insertion failures. This often involves a loop that attempts to re-insert kicked-out items.
*   **Fingerprints**: Instead of storing the full item, Cuckoo filters store a small fingerprint of the item. The size of this fingerprint directly impacts the false positive rate and memory usage.
*   **Bucket Size**: The number of entries per bucket (e.g., 2, 4, 8) affects both the load factor (how full the filter can get before insertions become difficult) and the false positive rate. Larger bucket sizes generally allow for higher load factors.
*   **Hash Functions**: Two hash functions are typically used: one to determine the primary bucket and another to determine the alternate bucket based on the fingerprint. These hash functions need to be carefully chosen to ensure good distribution.
*   **Resizing**: If an insertion fails, the Cuckoo filter needs to be resized and all existing items re-inserted. This can be a costly operation and should be managed to avoid frequent occurrences.
*   **Deletion**: Cuckoo filters support deletion by simply removing the fingerprint from its location. This is a significant advantage over standard Bloom filters.

### Performance Analysis

*   **Space Complexity**: `O(N * f)` bits, where `N` is the capacity (number of items) and `f` is the fingerprint size. Often more space-efficient than Bloom filters for low false positive rates.
*   **Time Complexity**:
    *   **Add**: Amortized `O(1)` on average, but can be `O(max_displacements)` in the worst case if many items are kicked out.
    *   **Contains**: `O(1)` on average, as it only checks a few fixed locations.
    *   **Delete**: `O(1)` on average.
*   **Practical Performance**: Lookups are very fast and deterministic. Insertions are fast on average but can be slow in rare cases if many displacements occur.

### Trade-offs

*   **Insertion Failure**: Unlike Bloom filters, Cuckoo filters can become "full" and fail to insert new items, requiring a costly rebuild/resize operation. This is a significant drawback for applications that cannot tolerate insertion failures.
*   **Complexity**: More complex to implement than Bloom filters due to the cuckoo hashing mechanism and displacement handling.
*   **Space vs. Load Factor**: There's a trade-off between space efficiency and the maximum load factor (how full the filter can get). Higher load factors (e.g., >95%) can lead to more frequent insertion failures.
*   **Deletion Support**: A major advantage over standard Bloom filters, as it supports exact deletions.
*   **False Positive Rate**: Can achieve lower false positive rates than Bloom filters for the same memory footprint in certain scenarios.

## Code Example

A basic Go implementation of the Cuckoo Filter can be found [here](code/cuckoo_filter.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd cuckoo-filter/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go cuckoo_filter.go
    ```