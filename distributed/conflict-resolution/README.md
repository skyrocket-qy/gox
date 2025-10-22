# Conflict Resolution

## Core

**Conflict resolution** is the process of managing and resolving inconsistencies that arise when multiple nodes in a distributed system attempt to update the same piece of data concurrently. In any system where data is replicated and can be modified in more than one location, conflicts are inevitable.

Without robust conflict resolution mechanisms, a distributed system can suffer from data divergence, where replicas of the same data become inconsistent over time. This can lead to data corruption, incorrect application behavior, and a violation of the system's integrity guarantees. This is a particularly important challenge in eventually consistent systems.

This section addresses various strategies and mechanisms for resolving these conflicts, including:
- **Last-Write-Wins (LWW):** A simple approach where the update with the latest timestamp is chosen as the winner.
- **Vector Clocks:** A more sophisticated mechanism that can detect concurrent updates and leave the resolution to the application.
- **Conflict-free Replicated Data Types (CRDTs):** Data structures that are designed to be concurrently modified without causing conflicts.
- **Application-specific Logic:** In some cases, the application itself is best equipped to resolve conflicts based on business rules.

## Comparison

| Strategy | Complexity | Data Loss Risk | Resolution Logic | Use Case |
|---|---|---|---|---|
| **[Last-Write-Wins (LWW)](./last-write-wins)** | Low | High | Timestamp-based | Simple, non-critical data |
| **[Vector Clocks](./vector-clocks)** | Medium | Low | Causal history | Detecting concurrency, manual resolution |
| **[CRDTs](./crdts)** | High | None | Automatic, deterministic | Collaborative editing, real-time apps |
| **[Timestamps with Logical Clocks](./timestamps-with-logical-clocks)** | Medium | Medium | Causal ordering | Distributed databases, event sourcing |

## Which service use it?



-   **Last-Write-Wins (LWW):** Often used in systems where simplicity is prioritized and occasional data loss due to concurrent updates is acceptable, such as caching systems or some eventually consistent key-value stores.

-   **Vector Clocks:** Employed in distributed databases (e.g., Riak) and collaborative systems to detect concurrent updates and allow for application-level conflict resolution or merging.

-   **CRDTs (Conflict-free Replicated Data Types):** Ideal for real-time collaborative applications (e.g., text editors, whiteboards) where multiple users can concurrently modify data, and conflicts need to be resolved automatically and deterministically.

-   **Timestamps with Logical Clocks:** Used in distributed databases and event sourcing systems to establish a causal order of events and resolve conflicts based on that order, providing stronger consistency guarantees than simple physical timestamps.

-   **Application-specific Logic:** Many complex distributed systems, especially those dealing with business logic, implement custom conflict resolution strategies tailored to their specific domain requirements.
