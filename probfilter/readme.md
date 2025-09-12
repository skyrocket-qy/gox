# Probabilistic Models for Pre-Database Filtering

This repository provides an overview and practical scenarios for various probabilistic data structures, useful for optimizing database access, filtering, and cardinality estimation in software engineering.

Each model is detailed in its own dedicated folder, providing explanations, use-case scenarios, and comparisons.

## Choosing the Right Probabilistic Model

To help you select the most suitable probabilistic model for your needs, consider the following categories and their primary applications:

### 1. Membership Testing (Is an item in a set?)

*   **[Bloom Filter](bloom-filter/README.md)**: Basic membership testing; space-efficient, no false negatives, but false positives possible.
*   **[Counting Bloom Filter](counting-bloom-filter/README.md)**: Supports deletion; more space than standard Bloom Filter.
*   **[Cuckoo Filter](cuckoo-filter/README.md)**: Dynamic additions/deletions; often more space-efficient for low false positive rates.
*   **[Quotient Filter](quotient-filter/README.md)**: Mergeable and resizable; often more space-efficient.
*   **[XOR Filter](xor-filter/README.md)**: Extremely fast and space-efficient for static sets; no false positives for items in set.

### 2. Cardinality Estimation (How many unique items are there?)

*   **[HyperLogLog](hyperloglog/README.md)**: Estimates unique items in large datasets with low memory.
*   **[HyperLogLog++](hyperloglog-plus-plus/README.md)**: Enhanced HyperLogLog with improved accuracy for smaller cardinalities.

### 3. Frequency Estimation & Top Items (How often does an item appear? What are the most frequent items?)

*   **[Count-Min Sketch](count-min-sketch/README.md)**: Estimates item frequencies; may overestimate, never underestimates.
*   **[Top-K](top-k/README.md)**: Identifies the most frequent items (top K) in a data stream.

### 4. Quantile & Percentile Estimation (What is the value at a certain percentile?)

*   **[t-digest](t-digest/README.md)**: Estimates percentiles from large datasets/streams; accurate at tails.

### 5. Similarity Estimation (How similar are two sets or documents?)

*   **[MinHash / Locality Sensitive Hashing (LSH)](minhash-lsh/README.md)**: Estimates Jaccard similarity; enables fast approximate nearest neighbor search.

### 6. Sorted Data Structures (Maintaining sorted order with probabilistic guarantees)

*   **[Skip List](skip-list/README.md)**: Probabilistic data structure for sorted data; similar performance to balanced trees with simpler implementation.

## Use Case Matrix: Choosing the Right Probabilistic Data Structure

| Feature / Use Case             | Bloom Filter                               | Counting Bloom Filter                          | Cuckoo Filter                                  | Quotient Filter                                | XOR Filter                                     | HyperLogLog                                    | HyperLogLog++                                  | Count-Min Sketch                               | t-digest                                       | Top-K                                          | Skip List                                      | MinHash / LSH                                  |
| :----------------------------- | :----------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- | :--------------------------------------------- |
| **Primary Goal**               | Membership Testing                         | Membership Testing                             | Membership Testing                             | Membership Testing                             | Membership Testing                             | Cardinality Estimation                         | Cardinality Estimation                         | Frequency Estimation                           | Quantile/Percentile Estimation                 | Top-K Items                                    | Sorted Data / Fast Operations                  | Similarity Estimation                          |
| **Supports Additions**         | Yes                                        | Yes                                            | Yes                                            | Yes                                            | No (Static)                                    | Yes                                            | Yes                                            | Yes                                            | Yes                                            | Yes                                            | Yes                                            | No (Static)                                    |
| **Supports Deletions**         | No                                         | Yes (Approximate)                              | Yes (Exact)                                    | Yes (Exact, Complex)                           | No (Static)                                    | No                                             | No                                             | No                                             | No                                             | No (for underlying counts)                     | Yes (Exact)                                    | No (Static)                                    |
| **False Positives**            | Yes                                        | Yes                                            | Yes                                            | Yes                                            | No (for items in set)                          | N/A (Approximation)                            | N/A (Approximation)                            | Yes (Overestimation)                           | N/A (Approximation)                            | Yes (for underlying counts)                    | No                                             | Yes (Approximation)                            |
| **False Negatives**            | No                                         | No                                             | No                                             | No                                             | No                                             | N/A                                            | N/A                                            | No                                             | N/A                                            | No                                             | No                                             | N/A                                            |
| **Memory Efficiency**          | Very High                                  | High (more than Bloom)                         | High (often better than Bloom for low FPR)     | Very High (often best)                         | Extremely High (best for static sets)          | Extremely High                                 | Extremely High (better for small counts)       | High                                           | High                                           | High (depends on K)                            | Moderate (more than list, less than tree)      | High                                           |
| **Lookup Speed**               | Very Fast (O(k))                           | Very Fast (O(k))                               | Very Fast (O(1))                               | Very Fast (O(1))                               | Extremely Fast (O(1))                          | N/A                                            | N/A                                            | Very Fast (O(d))                               | Fast (O(log C))                                | Fast (O(d + log K))                            | Fast (O(log N))                                | Fast (O(num_permutations))                     |
| **Implementation Complexity**  | Low                                        | Low-Medium                                     | Medium-High                                    | High                                           | Very High (Construction)                       | Medium                                         | Medium-High                                    | Medium                                         | Medium-High                                    | Medium-High                                    | Medium                                         | Medium-High                                    |
| **Mergeable**                  | Yes                                        | Yes                                            | No                                             | Yes                                            | No                                             | Yes                                            | Yes                                            | Yes                                            | Yes                                            | No                                             | No                                             | No                                             |
| **Use Cases**                  | Caching, Set Membership, Deduplication     | Dynamic Caching, Blacklisting                  | Dynamic Caching, Set Membership                | Distributed Systems, Set Reconciliation        | Static Dictionaries, Whitelisting              | Unique Visitor Counting, Analytics             | Accurate Unique Counts (small & large)         | Trending Topics, Anomaly Detection             | Latency Monitoring, Percentile Aggregation     | Real-time Leaderboards, Frequent Item Mining   | Database Indexing, Concurrent Data Structures  | Near-Duplicate Detection, Clustering, Plagiarism |

---
For more detailed information on each model, including mathematical foundations, implementation considerations, and code examples, navigate to their respective directories.
You can also find a list of further reading and resources in [RESOURCES.md](RESOURCES.md).