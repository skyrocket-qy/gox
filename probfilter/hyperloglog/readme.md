# HyperLogLog

*   **Space and Time Complexity**:
    *   Space: `O(m)` registers (kilobytes for billions of items).
    *   Time (Add): `O(1)` average.
    *   Time (Estimate, Merge): `O(m)`.

*   **Use Case**:
    *   **Counting unique visitors**: For a popular website with high traffic, estimating unique elements in large datasets with low memory.
    *   **Distributed Unique Event Counting**: In a distributed logging or event processing system, counting unique events across multiple servers or data streams.
    *   **Ad Impression/Click Uniqueness**: For an ad serving platform, estimating the number of unique ad impressions or clicks to prevent fraud and provide accurate billing.
    *   **Network Intrusion Detection**: Estimating the number of unique source/destination IP pairs or unique attack signatures observed in network traffic for large-scale intrusion detection systems.
    *   **Database Cardinality Estimation**: In large-scale data warehouses or analytical databases, estimating the cardinality of columns or join keys to optimize query execution plans.

*   **Pros**:
    *   Extremely space-efficient.
    *   Fast, constant-time insertions.
    *   Supports union operation for distributed counting.
*   **Cons**:
    *   Result is an approximation.
    *   Cannot retrieve actual items or support deletions.
    *   Bias for small cardinalities.
