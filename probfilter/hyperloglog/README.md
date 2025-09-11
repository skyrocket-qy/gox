# HyperLogLog

*   **Space and Time Complexity**:
    *   Space: `O(m)` registers (kilobytes for billions of items).
    *   Time (Add): `O(1)` average.
    *   Time (Estimate, Merge): `O(m)`.

*   **Use Case**: Counting unique visitors for a popular website with high traffic, estimating unique elements in large datasets with low memory.

*   **Pros**:
    *   Extremely space-efficient.
    *   Fast, constant-time insertions.
    *   Supports union operation for distributed counting.
*   **Cons**:
    *   Result is an approximation.
    *   Cannot retrieve actual items or support deletions.
    *   Bias for small cardinalities.
