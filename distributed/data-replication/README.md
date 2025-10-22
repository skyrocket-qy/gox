# Data Replication

## Core

**Data replication** is the process of storing copies of the same data on multiple nodes in a distributed system. This is a fundamental technique for building fault-tolerant, available, and high-performance systems.

## Comparison

| Strategy | Consistency | Performance | Complexity | Use Case |
|---|---|---|---|---|
| **[Synchronous](./sync)** | Strong | Low | Low | Financial systems, critical data |
| **[Asynchronous](./async)** | Eventual | High | Low | Caching, non-critical data |
| **[Semi-Synchronous](./semi-sync)** | Stronger than eventual | Medium | Medium | E-commerce, online gaming |
| **[Multi-Leader](./multi-leader)** | Eventual | High | High | Multi-datacenter deployments |
| **[Multi-Source](./multi-source)** | Eventual | High | High | Data aggregation, complex data flows |
| **[CRDT](./crdt)** | Eventual | High | High | Collaborative applications |
