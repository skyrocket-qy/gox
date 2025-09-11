# HyperLogLog

### Explanation

HyperLogLog is a probabilistic algorithm used for the count-distinct problem, which is the process of finding the number of unique elements in a dataset (also known as the cardinality). It can estimate the cardinality of very large datasets with a high degree of accuracy while using a very small and fixed amount of memory. It does not store the items themselves.

### Scenario: Counting unique visitors

A popular news website wants to count the number of unique visitors it receives each day. Storing every visitor's IP address or user ID in a set to count the unique entries would be memory-intensive, especially for a high-traffic site.

With HyperLogLog, the website can process a stream of visitor IDs and add each one to a HyperLogLog structure. At any time, the website can get a highly accurate estimate of the number of unique visitors by querying the structure, which might only be a few kilobytes in size. This is far more efficient than storing millions of unique identifiers.

### Comparison

*   **Pros**:
    *   Extremely space-efficient. It can estimate the cardinality of a set of billions of items with a typical error rate of around 2% using only 1.5 kB of memory.
    *   Fast, constant-time insertions.
    *   The union of two HyperLogLog structures can be computed, which allows for distributed or parallel counting.
*   **Cons**:
    *   The result is an approximation, not an exact count.
    *   It cannot retrieve the actual items that were added, only the estimated count of unique items.
    *   It does not support the deletion of items.

### Mathematical Foundations

HyperLogLog (HLL) estimates cardinality by observing the patterns of leading zeros in the hash values of the elements. The core idea is that the maximum number of leading zeros observed across all hash values provides an estimate of the logarithm of the number of unique elements.

The algorithm divides the hash space into `m` (a power of 2) registers. Each element is hashed, and its hash value is split into two parts: a prefix that determines which register to update, and a suffix whose leading zeros are counted. The maximum count of leading zeros for each register is stored. The final estimate is derived from the harmonic mean of these maximum leading zero counts, with a bias correction factor (`alpha`).

The standard error of the estimate is approximately `1.04 / sqrt(m)`. This means that increasing `m` (and thus memory usage) improves accuracy.

### Implementation Considerations

*   **Hash Function**: A high-quality, uniformly distributing hash function (e.g., MurmurHash, FNV-1a) is essential to ensure that hash values are random and the leading zero counts are representative.
*   **Register Size**: Each register typically stores a small integer (e.g., 5-6 bits) representing the maximum number of leading zeros. This contributes to HLL's extreme memory efficiency.
*   **Bias Correction**: The raw estimate from the harmonic mean often has a bias, especially for small cardinalities. HLL implementations include empirical bias correction mechanisms to improve accuracy.
*   **Small Range Correction**: For very small cardinalities, the HLL estimate can be inaccurate. Implementations often switch to a linear counting approach (counting zero registers) for better accuracy in this range.
*   **Large Range Correction**: For extremely large cardinalities, a correction factor is applied to account for hash collisions.
*   **Merging**: HLL structures can be easily merged by taking the maximum value for each corresponding register. This makes HLL suitable for distributed and parallel counting.
*   **No Deletions**: HLL does not support the deletion of elements.

### Performance Analysis

*   **Space Complexity**: `O(m)` registers, where `m` is the number of registers. Each register typically uses a few bits (e.g., 5-6 bits). This is extremely space-efficient, allowing estimation of billions of unique items with kilobytes of memory.
*   **Time Complexity**:
    *   **Add**: `O(1)` (constant time) on average, involving a single hash computation and a register update.
    *   **Estimate**: `O(m)` to iterate through all registers and compute the harmonic mean.
    *   **Merge**: `O(m)` to iterate through and take the maximum of corresponding registers.
*   **Practical Performance**: Very fast for additions. Estimation and merging are also efficient, scaling linearly with the number of registers, not the number of items.

### Trade-offs

*   **Approximation**: The primary trade-off is that HLL provides an approximation, not an exact count. The accuracy is tunable by adjusting the number of registers (`m`).
*   **No Item Retrieval**: HLL only estimates cardinality; it does not store the actual items, so you cannot retrieve them or check for individual membership.
*   **No Deletions**: HLL does not support the deletion of elements.
*   **Fixed Memory**: Memory usage is fixed once `m` is chosen, regardless of the number of items added.
*   **Bias for Small Cardinalities**: Standard HLL can have a noticeable bias for very small cardinalities, which is often addressed by using a small range correction or by switching to a hybrid approach (like HyperLogLog++).

## Code Example

A basic Go implementation of the HyperLogLog can be found [here](code/hyperloglog.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd hyperloglog/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go hyperloglog.go
    ```