# Bloom Filter

*   **Space and Time Complexity**:
    *   Space: `O(m)` bits.
    *   Time (Add, Contains): `O(k)` operations.
    *   Delete: Not supported.

*   **Use Case**:
    *   **Checking for existing usernames**: To reduce database load by quickly identifying unavailable usernames.
    *   **Database Query Optimization**: In a database system, quickly checking if a record exists in a large table before performing an expensive disk I/O operation.
    *   **Caching Layer**: As a first-level cache check in a backend system to avoid querying slower data stores for non-existent items.
    *   **Spam Filtering**: Identifying known spam emails or messages by checking their hashes against a Bloom filter of known spam signatures.
    *   **Deduplication in Data Pipelines**: Quickly identifying and discarding duplicate records in a data ingestion pipeline to reduce processing load and storage requirements.
    *   **Network Packet Filtering**: In network devices or firewalls, quickly checking if a packet's source/destination IP or port is on a blacklist/whitelist.

*   **Pros**:
    *   Very space-efficient.
    *   Fast, constant-time insertions and lookups.
    *   No false negatives.
*   **Cons**:
    *   False positives are possible.
    *   Cannot delete elements.
    *   Size must be decided in advance.
