# Strongly Consistent System Mode

## Core

**Strong consistency** is a consistency model that guarantees that once a write operation completes, any subsequent read operation, regardless of which node it is directed to, will see the value of that write. It is the most intuitive and strictest consistency model, ensuring that the system behaves as if it were a single, non-distributed system.

This model is often required in systems where data accuracy is critical, such as financial systems, banking applications, and inventory management systems.

## How It Works

Achieving strong consistency in a distributed system is challenging because it requires all replicas of a piece of data to be updated simultaneously. Since network communication is not instantaneous, this typically involves a coordination protocol to ensure that all nodes agree on the state of the data.

Common techniques for achieving strong consistency include:
-   **Single-Master Replication:** All write operations are directed to a single leader node. The leader applies the change and then replicates it to follower nodes. Reads can be served from the leader to guarantee they see the latest write, or from followers if some replication lag is acceptable (which would not be strictly strong consistency).
-   **Quorum-Based Protocols:** As seen in the quorum-based system mode, by configuring the read (R) and write (W) quorums such that **R + W > N** (where N is the total number of replicas), the system can guarantee that any read quorum will overlap with the most recent write quorum, ensuring the latest value is read.
-   **Consensus Algorithms:** Protocols like Paxos and Raft are designed to have a set of nodes agree on a value (or a sequence of values in a log). These algorithms ensure that once a value is committed by the group, it is durable and all nodes will eventually agree on it, providing a foundation for strongly consistent state machine replication.
-   **Two-Phase Commit (2PC):** A protocol used for distributed transactions to ensure that all participating nodes either commit or abort the transaction together, maintaining data integrity across nodes.

## Pros & Cons

### Pros

-   **Data Integrity and Accuracy:** Guarantees that all clients see the most up-to-date data, simplifying application logic.
-   **Predictability:** The behavior of the system is easy to reason about, as it mimics a single, atomic system.
-   **Correctness:** Essential for applications where stale data can lead to incorrect behavior or financial loss.

### Cons

-   **Higher Latency:** Write operations must often be confirmed by multiple nodes before they are acknowledged, which increases latency. Read operations may also be slower if they need to be coordinated.
-   **Lower Availability:** In the event of a network partition, a strongly consistent system may have to become unavailable to write (or even read) operations in order to avoid the risk of serving inconsistent data. This is the tradeoff described in the CAP theorem (choosing Consistency over Availability).
-   **Reduced Scalability:** The coordination overhead required for strong consistency can limit the system's horizontal scalability.

## Which service use it?

-   **Traditional Relational Databases (e.g., PostgreSQL, MySQL with ACID transactions):** These databases typically provide strong consistency within a single instance or through synchronous replication setups, ensuring that all transactions are atomic, consistent, isolated, and durable.
-   **Google Spanner:** A globally distributed, strongly consistent database that uses atomic clocks and GPS to achieve global external consistency.
-   **CockroachDB:** A distributed SQL database that provides strong consistency and transactional guarantees.
-   **etcd and Apache ZooKeeper:** These are distributed key-value stores primarily used for coordinating distributed systems, managing configuration, and implementing leader election, all of which require strong consistency.
-   **Financial Transaction Systems:** Banking and trading systems often require strong consistency to ensure that all transactions are processed correctly and that account balances are always accurate.
