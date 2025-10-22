# Causal Consistency

This section describes the causal consistency model, which is a weaker guarantee than strong consistency but stronger than eventual consistency. It ensures that operations that are causally related are seen by all processes in the same order.

## Which service use it?

-   **Collaborative Editing Systems:** In applications like Google Docs, causal consistency ensures that if user A types something, and user B then replies to it, all users will see A's typing before B's reply, even if updates arrive out of order due to network delays.
-   **Distributed Social Networks:** When a user posts a comment on another user's post, causal consistency ensures that all users see the original post before the comment, regardless of which server they are reading from.
-   **Some Distributed Databases:** Certain NoSQL databases or specialized distributed data stores might offer causal consistency as a consistency option, providing stronger guarantees than eventual consistency without the full overhead of strong consistency.
-   **Message Queues with Causal Ordering:** Advanced message queuing systems might implement causal ordering to ensure that messages that are causally related are processed in the correct sequence.
