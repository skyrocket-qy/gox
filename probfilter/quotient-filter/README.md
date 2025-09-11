# Quotient Filter

### Explanation

A Quotient filter is a space-efficient probabilistic data structure for set membership testing that is often more space-efficient than Bloom and Cuckoo filters. It stores a "fingerprint" of the item and has the advantage of being mergeable and resizable without requiring access to the original items, making it highly flexible.

### Scenario: Synchronizing data between distributed databases

Imagine a distributed database system where different nodes need to synchronize their data. Each node can create a Quotient filter to represent its set of keys. These filters, being very compact, can be efficiently sent over the network.

When a node receives a Quotient filter from another node, it can merge it with its own. This allows the node to quickly determine which keys it is missing without having to send the entire set of keys back and forth. This process, known as set reconciliation, is much more efficient with Quotient filters.

### Comparison

*   **Pros**:
    *   Often more space-efficient than Bloom and Cuckoo filters.
    *   Can be merged and resized without needing to rehash all the original items.
    *   Exhibits good data locality, which can lead to faster queries due to better cache performance.
*   **Cons**:
    *   More complex to implement than Bloom filters.
    *   Performance can degrade as the filter approaches its capacity.

### Mathematical Foundations

A Quotient Filter is based on a technique called "quotienting," where a hash value is split into a "quotient" and a "remainder" (fingerprint). The quotient determines the canonical bucket where an item *should* reside, and the remainder is stored as the fingerprint. Unlike Cuckoo filters, which use multiple hash functions to find alternative locations, Quotient filters use a single hash function and resolve collisions by shifting elements within a contiguous block of memory called a "run."

The false positive rate is determined by the size of the fingerprint.

### Implementation Considerations

*   **Bit Packing**: Quotient filters are highly optimized for space by packing fingerprints and metadata bits (is_occupied, is_continuation, is_shifted) into a compact bit array. This requires careful bit manipulation.
*   **Run Management**: The core complexity lies in managing "runs" of elements. When an item is inserted, it might need to shift existing items to maintain the sorted order of fingerprints within its run. Similarly, lookups involve navigating these runs.
*   **Metadata Bits**: Three metadata bits per slot are typically used:
    *   `is_occupied`: Indicates if the slot is occupied by an item.
    *   `is_continuation`: Indicates if the item in this slot is a continuation of a run that started earlier.
    *   `is_shifted`: Indicates if the item in this slot has been shifted from its canonical position.
*   **Hashing**: A single hash function is used to derive both the quotient and the remainder. The quality of this hash function is critical.
*   **Resizing and Merging**: Quotient filters have the advantage of being mergeable and resizable without requiring access to the original items, which is beneficial for distributed systems. However, these operations still require careful implementation.
*   **Deletion**: Deletion in Quotient filters is possible but complex, as it requires correctly updating the metadata bits and potentially shifting elements to maintain the integrity of runs.

### Performance Analysis

*   **Space Complexity**: `O(N * f)` bits, where `N` is the capacity and `f` is the fingerprint size. Often more space-efficient than Bloom and Cuckoo filters.
*   **Time Complexity**:
    *   **Add**: Amortized `O(1)` on average, but can be `O(N)` in worst-case scenarios (e.g., when many elements need to be shifted).
    *   **Contains**: `O(1)` on average, as it involves a few hash computations and bit checks.
    *   **Delete**: `O(1)` on average, but can be `O(N)` in worst-case scenarios.
*   **Practical Performance**: Generally very fast for lookups due to good cache locality. Insertion and deletion performance can degrade as the filter approaches full capacity.

### Trade-offs

*   **Complexity**: Significantly more complex to implement than Bloom or Cuckoo filters due to the intricate bit packing and run management.
*   **Performance Degradation at High Load**: Performance can degrade as the filter approaches its capacity, as insertions and deletions might require more shifts.
*   **Mergeability and Resizability**: A key advantage is the ability to merge two Quotient filters and resize them without needing access to the original items, which is highly beneficial for distributed systems and dynamic environments.
*   **Deletion Support**: Supports exact deletions, which is an advantage over standard Bloom filters.
*   **Space Efficiency**: Often achieves better space efficiency than Bloom and Cuckoo filters, especially for higher load factors.

## Code Example

A basic Go implementation of the Quotient Filter can be found [here](code/quotient_filter.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd quotient-filter/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go quotient_filter.go
    ```