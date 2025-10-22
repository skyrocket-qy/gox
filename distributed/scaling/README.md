# Scaling

## Core

**Scaling** is the process of designing and building a system in a way that it can handle a growing amount of work. In the context of a distributed system, this typically means being able to handle more users, more data, or more transactions. A scalable system is one that can maintain or even improve its performance and cost-effectiveness as the load on the system increases.

There are two main ways to scale a system:

1.  **Vertical Scaling (Scaling Up):** This involves adding more resources (e.g., CPU, memory, disk) to a single node. This is often the simplest way to scale, but it has its limits. There is a physical limit to how much you can scale up a single machine, and it can become very expensive.
2.  **Horizontal Scaling (Scaling Out):** This involves adding more nodes to the system. This is the most common way to scale a distributed system, as it can provide virtually unlimited scalability. However, it also introduces a number of challenges, such as how to distribute the load across the nodes and how to maintain consistency of the data.

### Scaling Techniques

There are a number of different techniques that can be used to scale a distributed system. Some of the most common ones include:

- **Load Balancing:** Distributing incoming requests across a cluster of servers. This can be done at the network level (e.g., with a hardware load balancer) or at the application level (e.g., with a software load balancer).
- **Sharding (Partitioning):** Splitting a large database into a number of smaller, more manageable pieces called shards. Each shard is stored on a separate node, which allows the database to be scaled out horizontally.
- **Caching:** Storing frequently accessed data in a fast, in-memory cache. This can significantly reduce the latency of read operations and reduce the load on the backend database.
- **Asynchronous Processing:** Using a message queue to decouple the processing of a request from the response. This can improve the responsiveness of the system and allow it to handle a higher volume of requests.

The choice of which scaling techniques to use depends on the specific requirements of the application. For example, a system that is read-heavy might benefit from caching, while a system that is write-heavy might benefit from sharding.

## Comparison

| Scaling Method | Approach | Cost | Scalability Limit | Use Case |
|---|---|---|---|---|
| **[Vertical Scaling](./vertical)** | Add resources to a single node | High | Limited by hardware | Monolithic applications, databases |
| **[Horizontal Scaling](./horizontal)** | Add more nodes to the system | Low | High | Microservices, distributed databases |

## Which service use it?

-   **Vertical Scaling:** Traditional monolithic applications, single-instance relational databases (e.g., a powerful SQL Server instance), and specialized high-performance computing tasks that benefit from a single, very powerful machine.
-   **Horizontal Scaling:** Most modern web applications, microservices architectures, distributed databases (NoSQL and sharded SQL), big data processing frameworks (e.g., Apache Spark, Hadoop), and cloud-native applications.
