# Strong Consistency

This section describes the strong consistency model, where all clients are guaranteed to see the same, most up-to-date version of the data at all times.

## Which service use it?

-   **Traditional Relational Databases (e.g., PostgreSQL, MySQL with ACID transactions):** These systems typically provide strong consistency within a single instance or through synchronous replication.
-   **Financial Systems:** Banking applications, stock trading platforms, and payment gateways require strong consistency to ensure accurate transaction processing and prevent data anomalies.
-   **Distributed Databases with Consensus (e.g., Google Spanner, CockroachDB, etcd, ZooKeeper):** These systems use algorithms like Paxos or Raft to ensure that all replicas agree on the order of operations, providing strong consistency guarantees across distributed nodes.
-   **Critical Business Applications:** Any application where data integrity and immediate visibility of changes are paramount, such as inventory management or reservation systems.
