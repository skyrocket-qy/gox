# MinHash / Locality Sensitive Hashing (LSH)

### Explanation

MinHash is a technique for quickly estimating how similar two sets are. It works by converting sets into "signatures" (small fixed-size arrays of hash values) such that the similarity of the signatures (Jaccard similarity) is a good estimate of the Jaccard similarity of the original sets. Locality Sensitive Hashing (LSH) is often used in conjunction with MinHash to group similar signatures together efficiently, allowing for fast approximate nearest neighbor searches in large datasets.

### Scenario: Detecting Near-Duplicate Documents

Imagine a search engine or a content management system that needs to identify near-duplicate web pages or documents to avoid indexing redundant content or to detect plagiarism. Storing and comparing every document pair is computationally infeasible for large corpora.

MinHash can be used to generate a small signature for each document. These signatures are then grouped using LSH. Documents falling into the same LSH bucket are likely to be near-duplicates. A more precise comparison can then be performed only on these candidate pairs, significantly reducing the computational load.

### Comparison

*   **Pros**:
    *   Efficiently estimates Jaccard similarity between sets.
    *   Scales well to very large datasets.
    *   LSH allows for fast approximate nearest neighbor search.
*   **Cons**:
    *   It's an approximation; accuracy depends on the size of the MinHash signature and LSH parameters.
    *   Can be sensitive to the choice of hash functions.
    *   Not suitable for exact similarity matching.

### Mathematical Foundations

**MinHash**: The fundamental principle behind MinHash is that the Jaccard similarity of two sets (size of intersection divided by size of union) can be estimated by the probability that their MinHash signatures are identical. Specifically, `J(A, B) ≈ P(minhash(A) == minhash(B))`. This is achieved by applying a set of random permutations to the elements of the sets and recording the minimum hash value for each permutation.

**Locality Sensitive Hashing (LSH)**: LSH is a technique to group similar items into the same "buckets" with high probability. For MinHash, LSH works by dividing the MinHash signature into `b` bands, each containing `r` rows. Each band is hashed to a bucket. If two signatures are similar, it's highly probable that at least one of their bands will hash to the same bucket, leading to a candidate pair. The probability of a candidate pair being generated is `1 - (1 - s^r)^b`, where `s` is the Jaccard similarity.

### Implementation Considerations

*   **Hash Functions**: A large number of independent hash functions are needed for MinHash. These are often simulated by using a single hash function with different seeds or by applying different mathematical transformations.
*   **Signature Size**: The number of permutations (`numPermutations`) directly impacts the accuracy of the Jaccard similarity estimate and the size of the MinHash signature. More permutations lead to better accuracy but higher computational cost and memory usage.
*   **LSH Parameters (bands and rows)**: The choice of `b` (bands) and `r` (rows per band) in LSH is crucial for tuning the trade-off between false positives (dissimilar items grouped together) and false negatives (similar items not grouped together). The product `b * r` must equal the `numPermutations` of the MinHash signature.
*   **Bucket Management**: LSH requires efficient hash tables (buckets) to store and retrieve candidate pairs.
*   **Shingling**: For text documents, items are typically converted into "shingles" (k-grams) before MinHashing. The choice of `k` for shingles affects the definition of "similarity."
*   **Scalability**: MinHash and LSH are designed for scalability, allowing similarity estimation and approximate nearest neighbor search on very large datasets where pairwise comparisons are infeasible.

### Performance Analysis

*   **Space Complexity**:
    *   **MinHash**: `O(num_permutations)` per signature.
    *   **LSH**: `O(N * bands)` where `N` is the number of items, for storing signatures in buckets.
*   **Time Complexity**:
    *   **MinHash Signature Generation**: `O(num_items_in_set * num_permutations)`.
    *   **Jaccard Similarity Estimation**: `O(num_permutations)`.
    *   **LSH Bucketing**: `O(num_permutations)` per signature.
    *   **Approximate Nearest Neighbor Search**: Highly efficient, as it avoids `O(N^2)` pairwise comparisons.
*   **Practical Performance**: MinHash and LSH are highly scalable for large datasets, enabling similarity search that would otherwise be computationally prohibitive.

### Trade-offs

*   **Approximation**: Both MinHash and LSH provide approximate results. The accuracy of similarity estimation and the recall/precision of LSH depend on the chosen parameters (`num_permutations`, `bands`, `rows`).
*   **Parameter Tuning**: Selecting optimal `num_permutations`, `bands`, and `rows` is crucial and often requires experimentation based on the desired accuracy and computational budget.
*   **Hash Function Quality**: The quality of the hash functions directly impacts the accuracy of MinHash signatures.
*   **No Exact Similarity**: MinHash provides an estimate of Jaccard similarity, not the exact value.
*   **Scalability vs. Accuracy**: There's a trade-off between the scalability achieved by LSH and the accuracy of the nearest neighbor search. More aggressive bucketing (fewer bands, more rows) can lead to higher recall but also more false positives.

## Code Example

A basic Go implementation of MinHash and LSH can be found [here](code/minhash_lsh.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd minhash-lsh/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go minhash_lsh.go
    ```