# Top-K

*   **Space and Time Complexity**:
    *   Space: `O(d * w + k)`, where `d * w` is for frequency estimator and `k` for top items.
    *   Time (Add): `O(d + log k)`.
    *   Time (GetTopK): `O(k log k)` or `O(k)`.

*   **Use Case**: Real-time "Top 10" list for a video streaming service to show most-watched videos.

*   **Pros**:
    *   Very space-efficient.
    *   Provides real-time estimates.
*   **Cons**:
    *   Approximation; may not find true top K items.
    *   Accuracy depends on structure size and algorithm.
