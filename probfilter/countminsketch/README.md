# Count-Min Sketch

### Explanation

A Count-Min Sketch is a probabilistic data structure that serves as a frequency table of events in a stream of data. It can be used to estimate the frequency of an item. It is similar to a counting Bloom filter, but it's designed to provide frequency estimates rather than just membership. A key characteristic that it may overestimate frequencies, but it never underestimates them.

### Scenario: Identifying trending topics

A social media platform wants to identify trending topics or hashtags in real-time. Tracking the exact count for every hashtag would require a massive amount of memory. Instead, the platform can use a Count-Min Sketch.

1.  As new posts with hashtags are created, the hashtags are added to the sketch, which increments their estimated frequency.
2.  To find trending topics, the platform can query the sketch for the estimated frequency of various hashtags.
3.  This allows the platform to identify which hashtags are being used most frequently and are likely trending, without storing exact counts for every single hashtag.

### Comparison

*   **Pros**:
    *   Very space-efficient for estimating frequencies in a large stream of data.
    *   Fast, constant-time updates and queries.
    *   Can be used to solve a range of frequency-related problems, such as finding the most frequent items ("heavy hitters").
*   **Cons**:
    *   It always provides an overestimation of the true frequency, never an underestimation. The amount of error can be reduced by increasing the size of the sketch.
    *   It cannot delete items (though variations that support this exist).
    *   It does not store the items themselves, only their estimated frequencies.

### Mathematical Foundations

A Count-Min Sketch uses `d` (depth) hash functions and `w` (width) counters for each hash function, forming a `d x w` matrix. When an item arrives, its count is incremented in `d` locations, one for each hash function. To estimate the frequency of an item, the minimum value across its `d` corresponding counters is taken. This minimum value is guaranteed to be an overestimation (never an underestimation) due to collisions.

The parameters `d` and `w` are chosen based on the desired error bounds:
*   `w = ceil(e / epsilon)`: `w` determines the width, which controls the absolute error `epsilon`.
*   `d = ceil(ln(1 / delta))`: `d` determines the depth, which controls the probability `delta` that the error exceeds `epsilon`.

### Implementation Considerations

*   **Hash Functions**: `d` independent hash functions are required. These should be chosen carefully to ensure uniform distribution and minimize collisions. Using a family of universal hash functions is common.
*   **Counter Array**: The `d x w` matrix of counters needs to be efficiently managed. Each counter typically stores an integer representing the frequency.
*   **Error Bounds**: Understanding and setting `epsilon` (absolute error) and `delta` (probability of error exceeding epsilon) is crucial for practical applications. A smaller `epsilon` requires a larger `w`, and a smaller `delta` requires a larger `d`.
*   **Memory Usage**: The memory usage is `O(d * w)`, which is very space-efficient for large data streams compared to storing exact counts.
*   **No Deletions**: Standard Count-Min Sketches do not support deletions. Variations like the "Count-Min Sketch with Conservative Updates" or "Count-Min Sketch with Deletions" exist but are more complex.
*   **Heavy Hitters**: Count-Min Sketches are often used as a component in algorithms for finding "heavy hitters" (most frequent items) in a stream.

### Performance Analysis

*   **Space Complexity**: `O(d * w)` counters, where `d` is the depth and `w` is the width. This is very space-efficient, as it's independent of the number of unique items in the stream.
*   **Time Complexity**:
    *   **Add**: `O(d)` operations, involving `d` hash computations and `d` counter increments.
    *   **Estimate**: `O(d)` operations, involving `d` hash computations and `d` counter lookups.
*   **Practical Performance**: Both additions and estimations are extremely fast, making it suitable for high-throughput data streams.

### Trade-offs

*   **Overestimation**: The primary trade-off is that the Count-Min Sketch always overestimates the true frequency of an item. It never underestimates. The amount of overestimation is bounded by `epsilon`.
*   **Approximation**: It provides an approximate count, not an exact one. The accuracy is tunable by adjusting `epsilon` and `delta`.
*   **No Deletions**: Standard Count-Min Sketches do not support deletions.
*   **Parameter Tuning**: Choosing appropriate `epsilon` and `delta` values is crucial and depends on the application's requirements for accuracy and confidence. A smaller error requires more memory.
*   **Heavy Hitters**: While it can estimate frequencies, finding the exact "heavy hitters" (most frequent items) often requires additional data structures or algorithms built on top of the Count-Min Sketch.

## Code Example

A basic Go implementation of the Count-Min Sketch can be found [here](code/count_min_sketch.go).

### How to Run the Example

1.  Navigate to the `code` directory:
    ```bash
    cd count-min-sketch/code
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go count_min_sketch.go
    ```