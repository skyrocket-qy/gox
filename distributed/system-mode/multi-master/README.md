# Multi-Master Mode

## Core

**Multi-master** (or active-active) replication is a system mode where multiple nodes in a distributed system can accept write operations. This is in contrast to a single-master (or master-slave) architecture, where only one node is designated to handle writes.

In a multi-master setup, each master node can process writes independently and then replicate its changes to the other master nodes. This model is often used to achieve high availability, distribute write traffic, and reduce write latency in geographically distributed systems.

## How It Works

When a write is made to any master node, that node updates its local data store and then propagates the change to all other master nodes in the system. Since writes can occur concurrently on different masters, conflicts can arise. For example, two clients might update the same piece of data on two different masters at the same time.

**Conflict Resolution** is a critical aspect of multi-master systems. Common strategies for handling conflicts include:
-   **Last Write Wins (LWW):** The write with the later timestamp overwrites the earlier one. This is simple to implement but can lead to data loss.
-   **Vector Clocks:** A more sophisticated timestamping mechanism that can detect concurrent updates and flag them for resolution.
-   **Application-Specific Logic:** The application itself is responsible for resolving conflicts, often by merging the conflicting changes or prompting a user for a decision.

Replication between masters can be synchronous or asynchronous. Asynchronous replication is more common, as synchronous replication would introduce high latency for every write.

## Pros & Cons

### Pros

-   **High Availability for Writes:** If one master node fails, other masters can continue to accept write requests, eliminating the single point of failure for writes.
-   **Low Write Latency:** Clients can connect to the nearest master, reducing network latency for write operations.
-   **Write Scalability:** The write load is distributed across multiple nodes, increasing the system's overall write throughput.

### Cons

-   **Conflict Resolution Complexity:** The need to handle and resolve write conflicts adds significant complexity to the system.
-   **Eventual Consistency:** Due to replication lag, data is typically eventually consistent. Different nodes might have different versions of the data for a short period.
-   **Difficult to Reason About:** The possibility of concurrent writes and conflicts can make it harder to reason about the state of the data and the behavior of the system.
