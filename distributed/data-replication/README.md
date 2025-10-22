# Data Replication

## Core

**Data replication** is the process of storing copies of the same data on multiple nodes in a distributed system. This is a fundamental technique for building fault-tolerant, available, and high-performance systems.

## Comparison

| Strategy | Consistency | Performance | Complexity | Use Case |
|---|---|---|---|---|
| **[Synchronous](./sync)** | Strong | Low | Low | Financial systems, critical data |
| **[Asynchronous](./async)** | Eventual | High | Low | Caching, non-critical data |
| **[Semi-Synchronous](./semi-sync)** | Stronger than eventual | Medium | Medium | E-commerce, online gaming |
| **[Multi-Leader](./multi-leader)** | Eventual | High | High | Multi-datacenter deployments |
| **[Multi-Source](./multi-source)** | Eventual | High | High | Data aggregation, complex data flows |
| **[CRDT](./crdt)** | Eventual | High | High | Collaborative applications |

## Which service use it?

-   **Synchronous Replication:** Used in systems requiring strong consistency and zero data loss, such as financial transaction systems, critical enterprise databases, and distributed consensus systems.
-   **Asynchronous Replication:** Common in web applications, caching layers, and analytics databases where high write throughput and low latency are prioritized over immediate consistency. Also used for disaster recovery where some data loss is acceptable.
-   **Semi-Synchronous Replication:** Often found in relational databases (e.g., MySQL, PostgreSQL) to provide a balance between consistency and performance, ensuring at least one replica has received the data before committing.
-   **Multi-Leader Replication:** Employed in geographically distributed systems or multi-datacenter deployments where local writes need to be fast and available, and conflicts are resolved later (e.g., some NoSQL databases, distributed file systems).
-   **Multi-Source Replication:** Used in data warehousing, data integration scenarios, or complex ETL (Extract, Transform, Load) pipelines where data from various sources needs to be consolidated into a single target.
-   **CRDT (Conflict-free Replicated Data Types):** Ideal for collaborative applications (e.g., real-time text editors, shared whiteboards), distributed counters, and other scenarios where concurrent updates need to be merged automatically without manual conflict resolution.
