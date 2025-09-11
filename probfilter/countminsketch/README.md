# Count-Min Sketch

*   **Space and Time Complexity**:
    *   Space: `O(d * w)` counters.
    *   Time (Add, Estimate): `O(d)` operations.

*   **Use Case**: Identifying trending topics or hashtags in real-time on a social media platform by estimating item frequencies.

*   **Pros**:
    *   Very space-efficient for estimating frequencies.
    *   Fast, constant-time updates and queries.
    *   Can be used for finding "heavy hitters".
*   **Cons**:
    *   Always overestimates frequency, never underestimates.
    *   Cannot delete items (in standard version).
    *   Does not store items themselves.
