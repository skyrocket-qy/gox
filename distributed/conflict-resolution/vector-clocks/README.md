# Vector Clocks for Conflict Resolution

This section explains how Vector Clocks are used in conflict resolution to detect concurrent updates and establish a partial ordering of events across distributed systems.

## Which service use it?



-   **Distributed Databases (e.g., Riak):** Riak uses vector clocks to detect and manage concurrent updates, allowing applications to resolve conflicts when they read data.

-   **Distributed File Systems:** Some distributed file systems use vector clocks to track the causal history of file versions, helping to resolve conflicts when files are modified concurrently.

-   **Collaborative Editing Systems:** In systems where multiple users can edit a document simultaneously, vector clocks can help in merging changes by understanding the causal relationships between different edits.

-   **Version Control Systems:** While not always explicitly called vector clocks, the underlying principles of tracking divergent histories and merging them are similar to how vector clocks help manage concurrent updates.
