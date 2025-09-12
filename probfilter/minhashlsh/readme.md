# MinHash / Locality Sensitive Hashing (LSH)

*   **Space and Time Complexity**:
    *   Space (MinHash): `O(num_permutations)` per signature.
    *   Space (LSH): `O(N * bands)`.
    *   Time (MinHash Signature Generation): `O(num_items_in_set * num_permutations)`.
    *   Time (Jaccard Similarity Estimation, LSH Bucketing): `O(num_permutations)`.
    *   Time (Approximate Nearest Neighbor Search): Highly efficient, avoids `O(N^2)`.

*   **Use Case**: Detecting near-duplicate documents in large corpora to avoid indexing redundant content or detect plagiarism.

*   **Pros**:
    *   Efficiently estimates Jaccard similarity.
    *   Scales well to very large datasets.
    *   LSH allows fast approximate nearest neighbor search.
*   **Cons**:
    *   Approximation; accuracy depends on signature size and LSH parameters.
    *   Can be sensitive to hash function choice.
    *   Not suitable for exact similarity matching.
