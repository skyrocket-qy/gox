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