# t-digest

*   **Space and Time Complexity**:
    *   Space: `O(compression)` or `O(number_of_centroids)`.
    *   Time (Add): `O(log(number_of_centroids))` average.
    *   Time (Quantile Estimation): `O(number_of_centroids)` or `O(log(number_of_centroids))`.
    *   Time (Merge): `O(number_of_centroids)`.

*   **Use Case**:
    *   **Monitoring API response times**: To calculate percentiles (e.g., 95th, 99th) in a space-efficient way.
    *   **Backend Service Latency Monitoring**: Collecting and summarizing latency distributions for various microservices or API endpoints to identify performance bottlenecks.
    *   **Database Query Performance**: Tracking the distribution of query execution times in a database backend to understand typical performance and identify outlier queries.
    *   **User Behavior Analytics**: Analyzing the distribution of user session durations, page load times, or time spent on specific features in a web application backend.
    *   **Resource Utilization Monitoring**: Summarizing the distribution of CPU, memory, or network usage across a fleet of servers or containers to identify resource hogs and optimize allocation.

*   **Pros**:
    *   Very space-efficient for estimating percentiles.
    *   Can be merged for distributed calculations.
    *   More accurate at the tails of the distribution.
*   **Cons**:
    *   Approximation, not exact.
    *   Accuracy depends on digest size.
    *   No deletions.
