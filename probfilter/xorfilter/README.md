# XOR Filter

*   **Space and Time Complexity**:
    *   Space: `O(N)` bits, ~1.23 bits per item.
    *   Time (Construction): `O(N)` average.
    *   Time (Contains): `O(1)`.
    *   Add/Delete: Not supported.

*   **Use Case**:
    *   **Serving static assets from a CDN**: Where the set of assets is large but changes infrequently.
    *   **Backend API Whitelisting**: For a backend service that only allows requests from a predefined set of API keys or client IDs, enabling quick validation without database hits.
    *   **Known Malicious IP/Domain Lists**: A backend security service can maintain a static list of known malicious IP addresses or domains for extremely fast lookups to block traffic.
    *   **Routing Table Optimization**: In network devices or backend routing services, for very fast lookup of destination addresses to determine if a route exists, reducing latency.

*   **Pros**:
    *   Less space than Bloom or Cuckoo filters.
    *   Faster lookups than Bloom or Cuckoo filters.
    *   No false positives for items in the set.
*   **Cons**:
    *   Static data structure; requires complete rebuild for modifications.
    *   More complex construction.
