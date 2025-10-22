# Consistency Models

## Core

In a distributed system, a **consistency model** is a contract between the system and the application that specifies the guarantees the system will provide regarding the visibility and ordering of updates to data. When data is replicated across multiple nodes, it's not always possible to ensure that all clients see the same data at the same time. A consistency model defines the rules for how and when the system will propagate updates and make them visible to clients.

The choice of a consistency model has a significant impact on the performance, availability, and complexity of a distributed system. There is often a trade-off between the strength of the consistency guarantee and the performance and availability of the system. Stronger consistency models are easier for developers to reason about, but they often come at the cost of higher latency and lower availability. Weaker consistency models can provide better performance and availability, but they require developers to handle the complexities of eventual consistency.

## Comparison

| Model | Guarantee | Performance | Use Case |
|---|---|---|---|
| **[Strong Consistency](./strong-consistency)** | All replicas are always up-to-date | High latency | Financial systems, critical data |
| **[Eventual Consistency](./eventual-consistency)** | Replicas will eventually be consistent | Low latency | Social media, e-commerce |
| **[Causal Consistency](./causal-consistency)** | Causal order of operations is preserved | Medium latency | Collaborative editing, chat |

## Which service use it?



-   **Strong Consistency:** Financial transaction systems, banking applications, and critical data management systems where immediate data accuracy is paramount.

-   **Eventual Consistency:** Social media feeds, e-commerce product catalogs, DNS, and large-scale web services where high availability and performance are prioritized over immediate consistency.

-   **Causal Consistency:** Collaborative editing applications, distributed social networks, and systems that need to preserve the causal order of events without requiring global strong consistency.
