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

## Pros & Cons

### Pros

-   **Flexibility:** The ability to tune the system's behavior allows it to be adapted to a wide range of use cases with different requirements.
-   **Dynamic Adaptation:** Some systems can even change their behavior dynamically in response to changing network conditions or application needs.

### Cons

-   **Complexity:** The flexibility of tunable systems comes at the cost of increased complexity. It can be difficult to reason about the system's behavior under all possible conditions.
-   **Risk of Misconfiguration:** If not configured correctly, a tunable system can provide weaker guarantees than intended, leading to data loss or other problems.
