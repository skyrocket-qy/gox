# Consistency Models

## Core

In a distributed system, a **consistency model** is a contract between the system and the application that specifies the guarantees the system will provide regarding the visibility and ordering of updates to data. When data is replicated across multiple nodes, it's not always possible to ensure that all clients see the same data at the same time. A consistency model defines the rules for how and when the system will propagate updates and make them visible to clients.

The choice of a consistency model has a significant impact on the performance, availability, and complexity of a distributed system. There is often a trade-off between the strength of the consistency guarantee and the performance and availability of the system. Stronger consistency models are easier for developers to reason about, but they often come at the cost of higher latency and lower availability. Weaker consistency models can provide better performance and availability, but they require developers to handle the complexities of eventual consistency.

## Comparison