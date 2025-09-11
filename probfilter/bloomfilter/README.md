# Bloom Filter

*   **Space and Time Complexity**:
    *   Space: `O(m)` bits.
    *   Time (Add, Contains): `O(k)` operations.
    *   Delete: Not supported.

*   **Use Case**: Checking for existing usernames to reduce database load by quickly identifying unavailable usernames.

*   **Pros**:
    *   Very space-efficient.
    *   Fast, constant-time insertions and lookups.
    *   No false negatives.
*   **Cons**:
    *   False positives are possible.
    *   Cannot delete elements.
    *   Size must be decided in advance.
