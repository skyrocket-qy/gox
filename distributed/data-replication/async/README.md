# Asynchronous Replication

This section describes Asynchronous Replication, a data replication mode where the primary node commits a transaction before receiving confirmation that the replica nodes have received or applied the changes, prioritizing performance and availability over immediate consistency.

## Which service use it?



-   **Web Applications:** Many large-scale web applications use asynchronous replication for their databases to handle high write throughput and distribute read loads, accepting eventual consistency.

-   **Caching Layers:** Distributed caches often use asynchronous replication to propagate updates to cached data across multiple nodes, prioritizing low latency for writes.

-   **Analytics Databases and Data Warehouses:** Data is often asynchronously replicated from operational databases to analytical systems, where immediate consistency is less critical than data volume and query performance.

-   **Disaster Recovery (DR) Setups:** Asynchronous replication is commonly used for DR sites, where the primary goal is to have a recent copy of data available in case of a catastrophic failure, even if it means some potential data loss.

-   **Content Delivery Networks (CDNs):** Content updates are asynchronously replicated to edge servers around the globe.
