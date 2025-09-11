# t-digest

### Explanation

A t-digest is a probabilistic data structure for estimating the rank of an element in a sorted sequence of numbers, and for estimating percentiles (e.g., the 99th percentile). It is particularly useful for summarizing the distribution of a large dataset or a stream of data in a space-efficient way.

### Scenario: Monitoring API response times

An e-commerce platform wants to monitor the response times of its API endpoints. It needs to calculate the 95th and 99th percentile response times to ensure that most users are having a good experience. Storing all response times to calculate these percentiles would be very memory-intensive.

Instead, the platform can use a t-digest. It adds each response time to the t-digest, which maintains a compact summary of the distribution. The platform can then query the t-digest to get accurate estimates of the 95th and 99th percentile response times.

### Comparison

*   **Pros**:
    *   Very space-efficient for estimating percentiles.
    *   Can be merged, allowing for distributed calculations.
    *   More accurate at the tails of the distribution (e.g., p99, p99.9) than some other methods.
*   **Cons**:
    *   It is an approximation, not an exact calculation.
    *   The accuracy depends on the size of the digest.

### Mathematical Foundations

A t-digest summarizes a distribution by maintaining a set of "centroids," where each centroid represents a cluster of data points. Each centroid has a `mean` (the average value of the points it represents) and a `count` (the number of points it represents). The key idea is that centroids representing data points in denser regions (e.g., near the median) are smaller (represent fewer points) and more numerous, while centroids in sparser regions (e.g., tails of the distribution) are larger (represent more points) and fewer. This adaptive compression allows t-digests to provide more accurate quantile estimates at the tails of the distribution compared to other methods.

The `compression` parameter controls the number of centroids and thus the accuracy. A higher compression value leads to more centroids and better accuracy.

### Implementation Considerations

*   **Centroid Management**: The core of a t-digest implementation involves efficiently adding new data points, merging existing centroids, and maintaining the desired number of centroids based on the compression parameter.
*   **Merging Strategy**: The merging strategy is crucial for maintaining accuracy, especially at the tails. Centroids are typically merged based on their proximity and the number of data points they represent, ensuring that smaller centroids are preserved in areas of high density.
*   **Sorting**: Centroids need to be sorted by their mean value for efficient quantile estimation and merging.
*   **Quantile Estimation**: Quantiles are estimated by linearly interpolating between the means of adjacent centroids, weighted by their counts.
*   **Mergeability**: A significant advantage of t-digests is their mergeability. Two t-digests can be combined into a single t-digest, making them suitable for distributed and parallel processing of data streams.
*   **No Deletions**: Standard t-digests do not support the deletion of individual data points.

### Performance Analysis

*   **Space Complexity**: `O(compression)` or `O(number_of_centroids)`. The memory usage is proportional to the compression factor, not the number of data points, making it very space-efficient for large datasets.
*   **Time Complexity**:
    *   **Add**: `O(log(number_of_centroids))` on average, as it involves finding the closest centroid and potentially merging.
    *   **Quantile Estimation**: `O(number_of_centroids)` to iterate through centroids, or `O(log(number_of_centroids))` if centroids are kept sorted and a binary search is used.
    *   **Merge**: `O(number_of_centroids)` to merge two t-digests.
*   **Practical Performance**: Efficient for both additions and quantile queries, making it suitable for real-time monitoring and streaming data.

### Trade-offs

*   **Approximation**: The primary trade-off is that t-digest provides an approximation of quantiles, not exact values. The accuracy is tunable by adjusting the `compression` parameter.
*   **Accuracy at Tails**: While generally more accurate at the tails than some other methods (like KLL sketches), its accuracy is still an approximation.
*   **No Deletions**: Standard t-digests do not support the deletion of individual data points.
*   **Complexity**: More complex to implement than simpler probabilistic data structures due to the centroid management and merging logic.
*   **Mergeability**: A significant advantage for distributed systems, allowing for efficient aggregation of quantile information from multiple sources.

## Code Example

A basic Go implementation of the t-digest can be found [here](code/t_digest.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd t-digest/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go t_digest.go
    ```