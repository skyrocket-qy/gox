# Eventually Consistent System Mode

## Core

**Eventual consistency** is a consistency model used in distributed computing that informally guarantees that, if no new updates are made to a given data item, all accesses to that item will eventually return the last updated value.

In an eventually consistent system, there is a period of time during which different nodes in the system may have different copies of the data. This is in contrast to a strongly consistent system, where all nodes have the same copy of the data at all times.

## How It Works

Eventual consistency is typically achieved through replication. When a write is made to one node, that node updates its local copy of the data and then asynchronously sends the update to the other nodes in the system.

There are a number of different ways to propagate updates to other nodes, including:

-   **Anti-entropy:** A process in which nodes periodically compare their data with other nodes and resolve any differences.
-   **Gossip:** A process in which nodes randomly exchange data with other nodes.
-   **Read repair:** A process in which inconsistencies are detected and resolved when a read is performed.

## Pros & Cons

### Pros

-   **High Availability:** Eventually consistent systems are highly available, as they can continue to operate even if some nodes are offline.
-   **Low Latency:** Eventually consistent systems have low latency, as writes can be acknowledged as soon as they are made to a single node.
-   **Scalability:** Eventually consistent systems are highly scalable, as new nodes can be added to the system without affecting the existing nodes.

### Cons

-   **Stale Data:** The main drawback of eventual consistency is that it is possible to read stale data. This can be a problem for applications that require strong consistency, such as financial systems.
-   **Complexity:** Eventual consistency can add complexity to the system, as developers need to be aware of the possibility of reading stale data and design their applications accordingly.

## Which service use it?

Several popular services and systems leverage eventual consistency to achieve high availability and scalability:

-   **Domain Name System (DNS):** DNS is a classic example of an eventually consistent system. Changes to DNS records propagate across the internet over time, meaning that different DNS servers may have slightly different information at any given moment.
-   **Amazon S3 (Simple Storage Service):** Amazon S3 provides eventual consistency for most operations. When you write a new object to S3 or overwrite an existing one, it may take some time for the change to be fully propagated across all storage nodes.
-   **Apache Cassandra:** Cassandra is a NoSQL database designed for high availability and scalability, and it uses eventual consistency as its primary consistency model.
-   **DynamoDB:** Amazon's DynamoDB also offers eventual consistency as one of its consistency options, allowing for faster write operations at the cost of potentially reading stale data for a short period.
-   **Many social media platforms:** Platforms like Facebook and Twitter often use eventual consistency for things like follower counts, likes, and news feeds. This allows them to handle massive amounts of data and traffic without sacrificing availability.