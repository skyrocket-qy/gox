# System Modes

## Core

In the context of distributed systems, a **system mode** (also known as an architectural pattern or deployment model) describes the fundamental way in which the components of the system are organized and interact with each other. The choice of a system mode is a critical architectural decision that has a profound impact on the system's properties, such as its scalability, fault tolerance, consistency, and complexity.

Different system modes are suited for different types of applications and use cases. For example, a simple web application might be well-served by a traditional client-server model, while a large-scale, data-intensive application might require a more complex, peer-to-peer or microservices architecture.

### Common System Modes

This section provides an overview of some of the most common system modes found in distributed systems:

-   **Client-Server:** A traditional model where a central server provides services to a number of clients.
-   **Peer-to-Peer (P2P):** A decentralized model where all nodes are equal peers and can act as both clients and servers.
-   **Leader-Follower:** A model used in replicated systems where one node is designated as the leader and is responsible for coordinating the other nodes (the followers).
-   **Multi-Master:** A replication model where multiple nodes can accept writes, and the changes are replicated to all other nodes.
-   **Sharded (Partitioned):** A model where data is partitioned across a number of nodes, allowing the system to be scaled out horizontally.
-   **Shared-Nothing:** An architectural pattern where each node is independent and self-sufficient, and there is no single point of contention across the system.
-   **Microservices:** An architectural style that structures an application as a collection of loosely coupled services.

Understanding the trade-offs between these different system modes is essential for designing and building effective distributed systems.


## Comparison

| Mode | Scalability | Consistency | Availability | Complexity | Use Case |
|---|---|---|---|---|---|
| **[Leader-Follower](./leader-follower)** | High (read), Low (write) | Strong (read from leader) | High | Medium | Databases, replicated systems |
| **[Master-Slave](./master-slave)** | High (read), Low (write) | Strong (read from master) | High | Medium | Databases, replicated systems |
| **[Multi-Master](./multi-master)** | High | Eventual | High | High | Multi-datacenter deployments |
| **[Peer-to-Peer](./peer-to-peer)** | High | Eventual | High | High | File sharing, content delivery |
| **[Sharded/Partitioned](./sharded-partitioned)** | High | Varies | High | High | Large-scale databases |
| **[Shared-Nothing](./shared-nothing)** | High | Varies | High | High | Distributed databases |
| **[Shared-Everything](./shared-everything)** | Low | Strong | Low | Low | Traditional databases |
| **[Quorum-Based](./quorum-based)** | Medium | Strong | High | High | Distributed storage systems |
| **[Strongly Consistent](./strongly-consistent)** | Low | Strong | Low | High | Financial systems, critical data |
| **[Eventually Consistent](./eventually-consistent)** | High | Eventual | High | Medium | Social media, e-commerce |
| **[Event-Driven](./event-driven)** | High | Eventual | High | Medium | Microservices, real-time data processing |
| **[Log-Based](./log-based)** | High | Strong | High | High | Distributed databases, event sourcing |
| **[Federated Multi-Cluster](./federated-multi-cluster)** | High | Varies | High | High | Large-scale, multi-cloud deployments |
| **[CAP Tradeoff (Tunable)](./cap-tradeoff-tunable)** | Varies | Tunable | Tunable | High | Systems requiring flexibility |