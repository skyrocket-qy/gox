# Bloom Filter

### Explanation

A Bloom filter is a space-efficient probabilistic data structure used to test whether an element is a member of a set. It consists of a bit array and multiple hash functions. False positive matches are possible, but false negatives are not. This means the filter can tell you if an element *might* be in the set or is *definitely not* in the set.

### Scenario: Checking for existing usernames

Imagine a service like Gmail or Twitter, where usernames must be unique. Every time a new user signs up, the service has to check if the desired username is already taken. This would require a database query for every attempt.

Instead, we can use a Bloom filter that stores all existing usernames. When a user tries to register a new username:
1. The username is checked against the Bloom filter.
2. If the filter returns "definitely not in the set," we know the username is available without needing to query the database.
3. If the filter returns "might be in the set" (a positive match), we then perform a database query to confirm. Since there's a chance of a false positive, the database is the source of truth.

This approach significantly reduces the load on the database, as most invalid username attempts are filtered out before a query is ever made.

### Comparison

*   **Pros**:
    *   Very space-efficient compared to storing all items in a hash set.
    *   Fast, constant-time insertions and lookups (O(k), where k is the number of hash functions).
    *   No false negatives.
*   **Cons**:
    *   False positives are possible. The rate can be tuned by adjusting the size of the filter and the number of hash functions.
    *   Cannot delete elements from a standard Bloom filter (though variations like Counting Bloom Filters exist).
    *   The size of the filter must be decided in advance based on the expected number of items.

### Mathematical Foundations

The core of a Bloom filter's efficiency lies in its probabilistic nature. The false positive rate (FPR), denoted as `p`, is a critical parameter. The optimal number of bits `m` and hash functions `k` for a given number of expected elements `n` and desired false positive rate `p` can be calculated as follows:

*   **Optimal number of bits (m)**: `m = -(n * ln(p)) / (ln(2)^2)`
*   **Optimal number of hash functions (k)**: `k = (m / n) * ln(2)`

These formulas ensure that for a given `n` and `p`, the filter uses the minimum possible space while maintaining the desired error rate. The probability of a false positive increases with the number of elements added and decreases with the size of the bit array and the number of hash functions.

### Implementation Considerations

*   **Hash Functions**: The quality of hash functions is crucial. They should be independent and uniformly distribute elements across the bit array. Using multiple independent hash functions can be achieved by combining two universal hash functions (e.g., `h(x) = (h1(x) + i * h2(x)) mod m` for `i` from 0 to `k-1`).
*   **Bit Array Management**: Efficiently managing the bit array (e.g., using a `[]byte` slice and bitwise operations in Go) is important for memory efficiency and performance.
*   **Capacity Planning**: The size of the Bloom filter (`m`) and the number of hash functions (`k`) must be determined upfront based on the expected number of elements (`n`) and the acceptable false positive rate (`p`). If the number of elements significantly exceeds `n`, the false positive rate will increase rapidly.
*   **No Deletions**: Standard Bloom filters do not support element deletion. If deletions are required, a Counting Bloom Filter or Cuckoo Filter should be considered.

### Performance Analysis

*   **Space Complexity**: `O(m)` bits, where `m` is the size of the bit array. This is highly space-efficient, often requiring only a few bits per element.
*   **Time Complexity**:
    *   **Add**: `O(k)` operations, where `k` is the number of hash functions. Each operation involves hashing and setting a bit.
    *   **Contains**: `O(k)` operations. Each operation involves hashing and checking a bit.
    *   **Delete**: Not supported in standard Bloom filters.
*   **Practical Performance**: Bloom filters are extremely fast for both insertions and lookups due to their constant-time hash operations and direct bit array access. The speed is largely independent of the number of elements stored, depending only on `k`.

### Trade-offs

*   **Space vs. False Positive Rate**: There's a direct trade-off between the memory allocated (`m`) and the acceptable false positive rate (`p`). To reduce `p` for a given number of elements `n`, `m` must increase.
*   **Speed vs. False Positive Rate**: Increasing the number of hash functions `k` can reduce the false positive rate (up to an optimal `k`), but it also increases the time taken for Add and Contains operations.
*   **Static Capacity**: The filter's size is fixed at creation. If the number of elements `n` significantly exceeds the planned capacity, the false positive rate will degrade rapidly. Resizing requires rebuilding the entire filter.
*   **No Deletions**: This is a fundamental limitation. If elements need to be removed, a Counting Bloom Filter or Cuckoo Filter is a better choice, but they come with their own trade-offs (e.g., increased memory, more complex implementation).

## Code Example

A basic Go implementation of the Bloom Filter can be found [here](code/bloom_filter.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd bloom-filter/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go bloom_filter.go
    ```