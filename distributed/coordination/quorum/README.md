# Quorum

This section explains the concept of Quorum in distributed systems, particularly its role in ensuring consistency and availability for coordination.

## Which service use it?



-   **Distributed Storage Systems (e.g., Apache Cassandra, Amazon DynamoDB):** These systems allow users to configure read and write quorums to achieve desired consistency levels. For example, a write quorum of `W` means `W` replicas must acknowledge a write before it's considered successful.

-   **Consensus Algorithms (e.g., Paxos, Raft):** These algorithms fundamentally rely on a majority (quorum) of nodes to agree on a decision (e.g., leader election, log entry commitment) to ensure consistency and fault tolerance.

-   **Distributed Coordination Services (e.g., Apache ZooKeeper, etcd):** These services use quorum-based protocols to maintain a consistent and highly available shared state, which is crucial for distributed locks, leader election, and configuration management.

-   **Distributed File Systems:** Some distributed file systems use quorum-like mechanisms to ensure data integrity and availability, especially when dealing with metadata or critical control plane operations.
