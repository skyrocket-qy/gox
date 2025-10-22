# Fault Tolerance

## Core

**Fault tolerance** is a critical property of distributed systems that enables them to continue operating correctly even in the presence of failures (or "faults"). In a distributed system, where many independent components are communicating over an unreliable network, failures are not the exception—they are the norm. Therefore, designing systems that can withstand and recover from these failures is essential for building reliable and highly available services.

### Types of Faults

Faults in a distributed system can be broadly categorized into three types:

1.  **Node Failures:** A process or server crashes. This can be a "fail-stop" failure, where the node simply stops working, or a "Byzantine" failure, where the node continues to operate but in an incorrect or malicious way.
2.  **Network Failures:** Messages between nodes can be lost, delayed, or delivered out of order. The network can also become partitioned, where two or more groups of nodes are unable to communicate with each other.
3.  **Storage Failures:** Data stored on disk can be corrupted or lost.

### Techniques for Fault Tolerance

There are a number of techniques that can be used to build fault-tolerant distributed systems. Some of the most common ones include:

-   **Redundancy:** Having multiple components that can perform the same function. This can be achieved through:
    -   **Replication:** Maintaining multiple copies of data on different nodes. If one node fails, the data is still available on other nodes.
    -   **Redundant Components:** Having spare servers or network links that can take over if a primary component fails.
-   **Failure Detection:** Mechanisms to detect when a component has failed. This is often done using heartbeating, where nodes periodically send "I'm alive" messages to each other.
-   **Checkpointing and Recovery:** Periodically saving the state of a process to stable storage. If the process fails, it can be restarted from the last checkpoint.
-   **Graceful Degradation:** Designing the system to continue operating at a reduced level of functionality when some components have failed, rather than failing completely.

The goal of fault tolerance is to mask failures from the end-user, providing the illusion of a continuously available and correct system. The specific techniques used will depend on the system's requirements for availability, consistency, and performance.


## Comparison

| Technique | Recovery Time | Resource Overhead | Complexity | Use Case |
|---|---|---|---|---|
| **[Active-Passive Failover](./active-passive-failover)** | Fast | Low | Low | Stateful applications |
| **[Active-Active Cluster](./active-active-cluster)** | Instant | High | High | Stateless applications, load balancing |
| **[Checkpointing](./checkpoint)** | Medium | Medium | Medium | Long-running computations |
| **[Snapshot](./snapshot)** | Fast | High | Medium | Stateful applications, databases |
| **[Consensus-Based Recovery](./consensus-recovery)** | Slow | High | High | Distributed databases, state machine replication |
| **[Gossip-Based Recovery](./gossip-recovery)** | Slow | Low | Low | Peer-to-peer networks, cluster membership |