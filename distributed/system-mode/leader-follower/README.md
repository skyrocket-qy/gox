# Leader-Follower Mode

## Core

The **leader-follower** (also known as **primary-replica** or **master-slave**) system mode is a common architectural pattern for achieving fault tolerance and high availability in distributed systems. In this model, one node is designated as the **leader** (or **primary**), and all other nodes are designated as **followers** (or **replicas**).

The leader is responsible for handling all write requests. It processes the write, updates its own state, and then replicates the change to all of the followers. The followers are responsible for handling all read requests. This allows the system to be scaled out for read-heavy workloads by adding more followers.

## How It Works

The leader-follower model provides a simple way to achieve fault tolerance. If the leader fails, one of the followers can be promoted to be the new leader. This process is known as **leader election**.

There are a number of different algorithms that can be used for leader election, such as:

-   **Bully algorithm:** In the Bully algorithm, any node can initiate an election. It sends a message to all other nodes with a higher ID. If it doesn't receive a response, it declares itself the leader.
-   **Raft:** Raft is a consensus algorithm that is designed to be easy to understand. It is used in a number of popular systems, such as etcd and Consul.
-   **ZooKeeper:** ZooKeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services. It is often used for leader election in distributed systems.

## Pros & Cons

### Pros

-   **Strong Consistency:** The leader-follower model provides strong consistency for reads from the leader.
-   **High Availability:** The leader-follower model is highly available, as it can tolerate the failure of the leader.
-   **Read Scalability:** The leader-follower model is highly scalable for read-heavy workloads.

### Cons

-   **Write Bottleneck:** The leader is a single point of failure for write requests.
-   **Replication Lag:** There can be a delay between when a write is made to the leader and when it is replicated to the followers. This can be a problem for applications that require real-time data.
-   **Complexity:** The leader-follower model can be complex to implement, as it requires a mechanism for leader election and failover.
