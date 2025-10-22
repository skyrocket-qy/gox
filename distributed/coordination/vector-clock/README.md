# Vector Clocks

This section explains Vector Clocks as a mechanism for tracking causality and achieving coordination in distributed systems.

## Which service use it?



-   **Distributed Databases (e.g., Riak):** Riak uses vector clocks to detect and manage concurrent updates, allowing applications to resolve conflicts when they read data.

-   **Collaborative Editing Systems:** In applications where multiple users can edit a document simultaneously, vector clocks can help in merging changes by understanding the causal relationships between different edits.

-   **Distributed File Systems:** Some distributed file systems use vector clocks to track the causal history of file versions, helping to resolve conflicts when files are modified concurrently.

-   **Version Control Systems:** While not always explicitly called vector clocks, the underlying principles of tracking divergent histories and merging them are similar to how vector clocks help manage concurrent updates.

-   **Eventual Consistency Systems:** Vector clocks are a fundamental tool in many eventually consistent systems to ensure that causally related events are processed in the correct order, even if physical timestamps are unreliable.
