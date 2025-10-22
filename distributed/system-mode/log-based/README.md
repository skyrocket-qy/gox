# Log-Based System Mode

## Core

A **log-based** system mode is an architectural pattern in which the state of a system is represented as an immutable, append-only log of events. The log is the single source of truth, and the current state of the system is derived by replaying the events in the log.

This is in contrast to a traditional approach where the current state of the system is stored in a mutable database. In a log-based system, the database is considered to be a cache of the current state, and it can be rebuilt at any time by replaying the log.

## How It Works

Log-based systems are often used in conjunction with a consensus algorithm, such as Raft or Paxos, to ensure that all nodes in the system agree on the order of the events in the log.

When a write is made to the system, it is appended to the log. The log is then replicated to all of the nodes in the system. Once a majority of the nodes have acknowledged that they have received the write, it is considered to be committed.

The state of the system can be derived at any point in time by replaying the log up to that point. This makes it possible to do a number of interesting things, such as:

-   **Time travel:** It is possible to view the state of the system at any point in the past by replaying the log up to that point.
-   **Auditing:** The log provides a complete audit trail of all of the changes that have been made to the system.
-   **Debugging:** The log can be used to debug problems in the system by replaying the events that led up to the problem.

## Pros & Cons

### Pros

-   **Strong Consistency:** Log-based systems provide strong consistency, as all nodes in the system agree on the order of the events in the log.
-   **Durability:** The log is immutable, which means that it is very durable.
-   **Auditing:** The log provides a complete audit trail of all of the changes that have been made to the system.

### Cons

-   **Storage Overhead:** The log can grow to be very large, which can lead to storage overhead.
-   **Replay Time:** Replaying the log can be time-consuming, especially if the log is very large.
-   **Complexity:** Log-based systems can be complex to implement.

## Which service use it?

-   **Apache Kafka:** Kafka is a distributed streaming platform that fundamentally operates as a distributed commit log. It stores streams of records in a fault-tolerant, durable, and highly available manner.
-   **Event Sourcing Architectures:** In event-sourced systems, all changes to application state are stored as a sequence of immutable events in an event log. The current state is derived by replaying these events.
-   **Distributed Databases (e.g., PostgreSQL WAL, MySQL Binlog):** Many traditional databases use a Write-Ahead Log (WAL) or binary log to ensure durability and enable replication. These logs are essentially append-only sequences of changes.
-   **Change Data Capture (CDC) Systems:** CDC tools often rely on reading database transaction logs to capture and propagate changes to other systems in real-time.
