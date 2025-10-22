# Timestamps with Logical Clocks for Conflict Resolution

This section explains how timestamps combined with logical clocks (e.g., Lamport timestamps, Vector Clocks) are used to resolve conflicts in distributed systems by establishing a causal order of events.

## Which service use it?



-   **Distributed Databases:** Many distributed databases use logical clocks (like Lamport timestamps or vector clocks) internally to order events and resolve conflicts, especially in eventually consistent or multi-master replication scenarios.

-   **Event Sourcing Systems:** In event-sourced architectures, logical clocks can be used to ensure the correct ordering of events in the event log, which is crucial for reconstructing the state of an aggregate.

-   **Distributed Transaction Systems:** While complex, some distributed transaction protocols might leverage logical clocks to help in ordering operations and ensuring consistency across multiple participating nodes.

-   **Distributed Caching:** When multiple caches can be updated concurrently, logical clocks can help in determining the most recent version of a cached item.
