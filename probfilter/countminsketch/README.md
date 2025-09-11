# Count-Min Sketch

*   **Space and Time Complexity**:
    *   Space: `O(d * w)` counters.
    *   Time (Add, Estimate): `O(d)` operations.

*   **Use Case**:
    *   **Identifying trending topics or hashtags**: In real-time on a social media platform by estimating item frequencies.
    *   **Network Anomaly Detection**: Identifying unusually frequent source/destination IP pairs, port scans, or specific packet types in network traffic.
    *   **Database Query Frequency**: Estimating the frequency of specific queries or query patterns hitting a database backend to identify hot spots and optimize indexing or caching strategies.
    *   **API Usage Monitoring**: Tracking the frequency of API calls per user, per endpoint, or per application to identify heavy users, potential abuse, or popular features.
    *   **Fraud Detection in Financial Transactions**: Estimating the frequency of certain transaction types, amounts, or originating accounts to flag potentially fraudulent activities.
    *   **Content Popularity Tracking**: In a content delivery system, estimating the popularity of articles, videos, or images to inform caching decisions or content recommendations.

*   **Pros**:
    *   Very space-efficient for estimating frequencies.
    *   Fast, constant-time updates and queries.
    *   Can be used for finding "heavy hitters".
*   **Cons**:
    *   Always overestimates frequency, never underestimates.
    *   Cannot delete items (in standard version).
    *   Does not store items themselves.
