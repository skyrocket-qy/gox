# CAP Tradeoff (Tunable) System Mode

## Core

In distributed systems, the **CAP theorem** states that it is impossible for a distributed data store to simultaneously provide more than two out of the following three guarantees:

-   **Consistency:** Every read receives the most recent write or an error.
-   **Availability:** Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
-   **Partition Tolerance:** The system continues to operate despite an arbitrary number of messages being dropped (or delayed) by the network between nodes.

A **CAP tradeoff-tunable** system mode is one that allows system designers to make deliberate, often dynamic, tradeoffs between these three properties. Rather than being strictly CP (consistent and partition-tolerant) or AP (available and partition-tolerant), these systems can be configured to favor one property over the other, depending on the application's requirements.

## How It Works

Tunable systems often expose configuration parameters that allow administrators to control the system's behavior in the face of network partitions. For example, a system might allow you to specify the number of replicas that must acknowledge a write before it is considered successful.

-   **Favoring Consistency:** To favor consistency, a system might require a strict quorum of nodes to be available for reads and writes. If a partition occurs and a quorum cannot be reached, the system will become unavailable rather than risk returning stale data.
-   **Favoring Availability:** To favor availability, a system might allow reads and writes to be served by any available node, even if it is not in the majority partition. This can lead to conflicts, which must be resolved later, but it ensures that the system remains responsive.

## Which service use it?

-   **Apache Cassandra:** Cassandra allows developers to choose consistency levels (e.g., ONE, QUORUM, ALL) on a per-operation basis, effectively tuning the CAP tradeoff. A higher consistency level prioritizes consistency, while a lower one prioritizes availability and performance.
-   **Riak:** Riak is another NoSQL database that provides tunable consistency through its N, R, and W values (number of replicas, read quorum, write quorum), allowing users to decide how many nodes must agree for an operation to succeed.
-   **Amazon DynamoDB:** DynamoDB offers both eventually consistent and strongly consistent read options, allowing applications to choose the appropriate consistency model based on their requirements for a given read operation.
