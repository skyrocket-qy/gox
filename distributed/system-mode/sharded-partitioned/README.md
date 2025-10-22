# Sharded/Partitioned System Mode

## Core

**Sharding** (also known as **partitioning**) is a database architecture pattern used to horizontally scale a data store. In a sharded system, a large dataset is broken down into smaller, more manageable pieces called **shards** or **partitions**. Each shard is stored on a separate server or node.

This distribution of data allows the system to handle a larger volume of data and a higher load of read/write operations than would be possible with a single server. Sharding is a key technique for achieving horizontal scalability in distributed databases and other data-intensive applications.

## How It Works

A sharded system requires a **partitioning key** (or **shard key**), which is a specific field or column in the data that is used to determine which shard a particular piece of data belongs to. The system uses a **partitioning function** to map each key to a specific shard.

Common sharding strategies include:
-   **Range-Based Sharding:** Data is partitioned based on a range of values in the shard key. For example, a database of users might be sharded by the first letter of their last name (A-F on shard 1, G-M on shard 2, etc.). This approach is simple but can lead to uneven data distribution (hot spots).
-   **Hash-Based Sharding:** Data is partitioned based on the hash of the shard key. This typically results in a more even distribution of data across shards but makes range queries (e.g., find all users with last names starting with 'S') inefficient, as they would require querying all shards.
-   **Directory-Based Sharding:** A lookup table (or directory) is maintained that maps each shard key to the shard it resides on. This offers the most flexibility but introduces the overhead of maintaining and querying the lookup service.

A sharded architecture also requires a routing or query coordination layer. When a client sends a request, this layer determines which shard(s) hold the relevant data and routes the request accordingly. For queries that span multiple shards, the coordinator must gather results from each relevant shard and combine them before returning the final result to the client.

## Pros & Cons

### Pros

-   **Horizontal Scalability:** The storage capacity and throughput of the system can be increased by adding more shards (nodes).
-   **Improved Performance:** By distributing the data and workload, query latency can be reduced, and the system can handle higher concurrency.
-   **Increased Availability:** If one shard becomes unavailable, it only affects the data on that shard. The rest of the system can remain operational.

### Cons

-   **Increased Complexity:** Sharding adds significant complexity to the system. Developers need to manage data distribution, routing, and cross-shard queries.
-   **Rebalancing:** As data grows, shards may need to be rebalanced (split or moved) to maintain even distribution. This can be a complex and resource-intensive operation.
-   **Loss of ACID Guarantees:** Transactions that span multiple shards are difficult to coordinate and may not provide the same ACID guarantees as single-shard transactions.
-   **Hot Spots:** Poor choice of a shard key can lead to uneven data distribution, with some shards becoming performance bottlenecks ("hot spots").

## Which service use it?

-   **Large-Scale Relational Databases (e.g., sharded MySQL, PostgreSQL):** Many high-traffic web applications shard their relational databases to handle massive amounts of data and user requests.
-   **NoSQL Databases (e.g., MongoDB, Apache Cassandra, Apache HBase):** These databases are inherently designed for horizontal scaling through sharding/partitioning, distributing data across clusters.
-   **Distributed Search Engines (e.g., Elasticsearch, Apache Solr):** These systems shard their indexes across multiple nodes to handle large volumes of data and provide fast search capabilities.
-   **Data Warehouses and Analytics Platforms:** Large data processing systems often partition data to enable parallel processing and efficient querying of petabytes of information.
