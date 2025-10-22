# Shared-Nothing Architecture

## Core

**Shared-nothing** is a distributed computing architecture in which each node is a self-contained, independent system with its own private memory, CPU, and disk storage. The nodes do not share any resources and communicate with each other by passing messages over a network.

This architecture is the foundation for most modern, horizontally scalable distributed systems, including many NoSQL databases, data warehousing solutions, and large-scale web applications. Scaling is achieved by adding more independent nodes to the system, a process known as **horizontal scaling** (or scaling out).

## How It Works

In a shared-nothing system, data is typically partitioned (or sharded) across the nodes. Each node is responsible for storing and processing a subset of the total data. When a request comes in, a routing layer determines which node holds the relevant data and directs the request to that node.

If an operation requires data from multiple nodes, the nodes must communicate over the network to coordinate the operation. This message-passing introduces network latency, which is a key performance consideration in shared-nothing systems.

For example, a distributed database might partition its data by a user ID. All data for a specific user would be stored on a single node. A query for that user's data would be sent directly to the responsible node. A more complex query, like calculating the average of a value across all users, would require all nodes to calculate their local average and then send that result to a coordinating node to compute the final, global average.

## Pros & Cons

### Pros

-   **High Scalability:** The system can be scaled out almost linearly by adding more nodes. There is no central bottleneck to limit the system's growth.
-   **Fault Tolerance and High Availability:** Since nodes are independent, the failure of one node does not bring down the entire system. Other nodes can continue to operate, and the data on the failed node can be replicated for high availability.
-   **Parallel Processing:** The architecture is well-suited for parallel processing, as different nodes can work on different parts of a task simultaneously.

### Cons

-   **Increased Complexity:** Shared-nothing systems are more complex to design and manage than shared-everything systems. Developers must deal with issues like data partitioning, replication, distributed transactions, and network latency.
-   **Network Overhead:** Communication between nodes over the network is significantly slower than communication through shared memory. Performance can be limited by network bandwidth and latency.
-   **Consistency Challenges:** Maintaining data consistency across multiple, independent nodes is a major challenge. Many shared-nothing systems opt for eventual consistency over strong consistency to improve availability and performance.
-   **Rebalancing Costs:** Adding new nodes or handling node failures often requires rebalancing the data across the cluster, which can be a complex and I/O-intensive process.

## Which service use it?

-   **Most NoSQL Databases (e.g., Apache Cassandra, MongoDB, Apache HBase):** These databases are designed from the ground up to be horizontally scalable using a shared-nothing architecture, distributing data and processing across many independent nodes.
-   **Massively Parallel Processing (MPP) Data Warehouses (e.g., Teradata, Greenplum, Amazon Redshift):** These systems distribute large datasets and query processing across many nodes, each with its own CPU, memory, and storage.
-   **Distributed File Systems (e.g., Apache HDFS):** HDFS stores large files across multiple machines, with each DataNode managing its own local storage.
-   **Web Search Engines (e.g., Google Search):** The indexing and serving infrastructure of large search engines are built on shared-nothing principles to handle the immense scale of the web.
