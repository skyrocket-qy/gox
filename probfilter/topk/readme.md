# Top-K

*   **Space and Time Complexity**:
    *   Space: `O(d * w + k)`, where `d * w` is for frequency estimator and `k` for top items.
    *   Time (Add): `O(d + log k)`.
    *   Time (GetTopK): `O(k log k)` or `O(k)`.

*   **Use Case**:
    *   **Real-time "Top 10" list**: For a video streaming service to show most-watched videos.
    *   **Network Traffic Analysis**: Identifying the top N busiest IP addresses, ports, or protocols in real-time from network flow data.
    *   **Log Analysis**: Finding the most frequent error messages, user agents, or API endpoints being hit in a large stream of logs for operational monitoring and debugging.
    *   **Trending Hashtags/Keywords**: In a social media backend, identifying the most frequently used hashtags or keywords in real-time.
    *   **Fraud Detection**: Identifying the top N most frequent transaction patterns, credit card numbers, or user accounts that might indicate fraudulent activity in a stream of financial transactions.

*   **Pros**:
    *   Very space-efficient.
    *   Provides real-time estimates.
*   **Cons**:
    *   Approximation; may not find true top K items.
    *   Accuracy depends on structure size and algorithm.
