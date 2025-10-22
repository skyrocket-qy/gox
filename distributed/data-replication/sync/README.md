# Synchronous Replication

This section describes Synchronous Replication, a data replication mode where the primary node waits for confirmation from replica nodes that they have received and committed the transaction before committing itself, prioritizing strong consistency.

## Which service use it?



-   **Financial Transaction Systems:** Banking, trading, and payment processing systems where data integrity and zero data loss are absolutely critical. Transactions must be committed on multiple nodes before being acknowledged to the client.

-   **Critical Enterprise Databases:** Databases storing highly sensitive or business-critical information where any data loss or inconsistency could have severe consequences.

-   **Distributed Consensus Systems (e.g., Paxos, Raft):** While not strictly a replication strategy in the same sense as database replication, consensus algorithms inherently use synchronous communication to ensure all participants agree on a state before proceeding.

-   **High-Availability Clusters:** In active-passive or active-active setups for critical applications, synchronous replication ensures that the standby node always has an up-to-date copy of the data.
