# Consensus-Based Recovery

This section discusses Consensus-Based Recovery, where distributed consensus algorithms are leveraged to ensure consistent state recovery after failures in a distributed system.

## Which service use it?



-   **Distributed Databases (e.g., Google Spanner, CockroachDB):** These databases use consensus algorithms (like Paxos or Raft) to ensure that all replicas agree on the state of the data, enabling consistent recovery after node failures.

-   **State Machine Replication:** Systems that implement state machine replication (where all nodes execute the same sequence of operations) rely on consensus for fault tolerance and consistent recovery.

-   **Distributed Coordination Services (e.g., etcd, Apache ZooKeeper):** These services use consensus to maintain a consistent, highly available store for configuration data, service discovery, and leader election, allowing them to recover their state even after multiple node failures.

-   **Blockchain Technologies:** The underlying consensus mechanisms (e.g., Proof of Work, Proof of Stake) in blockchains are a form of consensus-based recovery, ensuring the integrity and consistency of the distributed ledger even in the presence of malicious actors or failures.
