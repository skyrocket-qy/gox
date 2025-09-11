# Top-K

### Explanation

Top-K is a probabilistic data structure that allows you to find the most frequent items in a data stream. It is a generalization of the "heavy hitters" problem. It keeps a small, fixed-size list of the most frequent items seen so far.

### Scenario: Real-time "Top 10" list

A video streaming service wants to show a real-time list of the "Top 10 most-watched videos right now". Tracking the exact watch count for millions of videos in real-time is challenging.

Instead, the service can use a Top-K algorithm. As users watch videos, the video IDs are fed into the Top-K structure. At any time, the service can query the structure to get a list of the top 10 most frequent video IDs, which represents the most-watched videos.

### Comparison

*   **Pros**:
    *   Very space-efficient, as it only stores a small number of items (K).
    *   Can provide real-time estimates of the most frequent items.
*   **Cons**:
    *   It is an approximation. It may not always find the true top K items, especially if the frequencies are close.
    *   The accuracy depends on the size of the structure and the algorithm used (e.g., "heavy hitters" with Count-Min Sketch).

### Mathematical Foundations

Top-K algorithms are designed to find the `k` most frequent items in a data stream, often with limited memory. Many Top-K algorithms are based on probabilistic data structures like the Count-Min Sketch. The Count-Min Sketch provides an estimated frequency for each item. A common approach is to maintain a small data structure (e.g., a min-heap or a hash map) that stores the `k` items with the highest estimated frequencies seen so far.

The accuracy of the Top-K list depends on the accuracy of the underlying frequency estimation mechanism (e.g., the `epsilon` and `delta` parameters of the Count-Min Sketch). While the algorithm aims to find the true top K items, it's an approximation, and there's a probability of missing some true top items or including some false positives, especially when frequencies are very close.

### Implementation Considerations

*   **Underlying Frequency Estimator**: A robust frequency estimation data structure (like Count-Min Sketch) is crucial for the accuracy of the Top-K list.
*   **Top-K Data Structure**: A min-heap is commonly used to efficiently maintain the `k` largest elements. When a new item's estimated frequency is higher than the smallest element in the heap, the smallest element is removed, and the new item is inserted.
*   **Handling Collisions/Approximations**: Since the frequency estimates are approximate, the Top-K list might not always be perfectly accurate. Strategies might be needed to periodically re-evaluate the top items or to handle cases where items have very similar frequencies.
*   **Memory Management**: The memory usage is primarily determined by the size of the underlying frequency estimator and the `k` items stored in the heap.
*   **Dynamic Updates**: Top-K algorithms are well-suited for streaming data, as they can be updated incrementally as new items arrive.

### Performance Analysis

*   **Space Complexity**: `O(d * w + k)`, where `d * w` is the size of the underlying frequency estimator (e.g., Count-Min Sketch) and `k` is the number of top items to track. This is very space-efficient for large streams.
*   **Time Complexity**:
    *   **Add**: `O(d + log k)` on average, where `O(d)` is for updating the frequency estimator and `O(log k)` is for heap operations.
    *   **GetTopK**: `O(k log k)` for sorting the heap elements, or `O(k)` if just retrieving without sorting.
*   **Practical Performance**: Efficient for real-time tracking of top items in high-throughput data streams.

### Trade-offs

*   **Approximation**: The primary trade-off is that the Top-K list is an approximation. It might not always identify the true top K items, especially if their frequencies are very close or if the underlying frequency estimator has significant error.
*   **Accuracy vs. Resources**: The accuracy depends on the parameters of the underlying frequency estimator (e.g., `epsilon` and `delta` for Count-Min Sketch) and the value of `k`. Higher accuracy or larger `k` generally requires more memory and computation.
*   **Dynamic Updates**: Supports dynamic updates, making it suitable for streaming data.
*   **No Exact Counts**: While it tracks estimated frequencies, it doesn't store exact counts for all items in the stream, only for the top K (or those in the heap).

## Code Example

A basic Go implementation of the Top-K algorithm can be found [here](code/top_k.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd top-k/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go top_k.go
    ```