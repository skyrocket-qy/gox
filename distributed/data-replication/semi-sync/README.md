# Semi-Synchronous Replication

This section explains Semi-Synchronous Replication, a data replication mode that offers a balance between synchronous and asynchronous replication, providing stronger consistency guarantees than async without the full latency impact of sync.

## Which service use it?



-   **Relational Databases (e.g., MySQL, PostgreSQL):** These databases often offer semi-synchronous replication as a built-in feature to provide a middle ground between the performance of asynchronous and the data safety of synchronous replication.

-   **E-commerce Platforms:** For critical operations like order placement or inventory updates, semi-synchronous replication can ensure that data is safely replicated to at least one other node before the transaction is committed, reducing the risk of data loss compared to pure asynchronous replication.

-   **Online Gaming:** In some online games, certain critical game state updates might use semi-synchronous replication to ensure a higher degree of consistency among players without incurring the full latency of synchronous replication.

-   **Content Management Systems (CMS):** For content updates that are important but not absolutely real-time critical, semi-synchronous replication can provide better data durability than asynchronous replication.
