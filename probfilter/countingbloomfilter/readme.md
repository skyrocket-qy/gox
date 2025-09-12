# Counting Bloom Filter

*   **Space and Time Complexity**:
    *   Space: `O(m * c)` bits, where `m` is number of counters and `c` is bits per counter.
    *   Time (Add, Contains, Remove): `O(k)` operations.

*   **Use Case**:
    *   **Managing a list of malicious URLs**: Where URLs need to be added and removed dynamically.
    *   **Dynamic Caching/Eviction Policies**: In a backend caching system, tracking items in cache and allowing for their removal when evicted.
    *   **Rate Limiting/Throttling**: For backend APIs, tracking the number of requests from specific users or IP addresses within a time window.
    *   **Distributed Blacklisting/Whitelisting**: Maintaining a shared blacklist or whitelist of entities where entries can be added and removed dynamically across nodes.
    *   **Session Management**: In a distributed web application, tracking active user sessions across multiple servers.

*   **Pros**:
    *   Supports deletion of elements.
    *   Relatively space-efficient.
*   **Cons**:
    *   Requires more space than standard Bloom filter.
    *   Counters can overflow.
    *   "Soft" deletions can introduce errors.
