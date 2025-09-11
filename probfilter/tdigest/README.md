# t-digest

*   **Space and Time Complexity**:
    *   Space: `O(compression)` or `O(number_of_centroids)`.
    *   Time (Add): `O(log(number_of_centroids))` average.
    *   Time (Quantile Estimation): `O(number_of_centroids)` or `O(log(number_of_centroids))`.
    *   Time (Merge): `O(number_of_centroids)`.

*   **Use Case**: Monitoring API response times to calculate percentiles (e.g., 95th, 99th) in a space-efficient way.

*   **Pros**:
    *   Very space-efficient for estimating percentiles.
    *   Can be merged for distributed calculations.
    *   More accurate at the tails of the distribution.
*   **Cons**:
    *   Approximation, not exact.
    *   Accuracy depends on digest size.
    *   No deletions.
