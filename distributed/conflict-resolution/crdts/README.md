# CRDTs (Conflict-free Replicated Data Types) for Conflict Resolution

This section details how CRDTs are designed to resolve conflicts automatically and deterministically, making them suitable for collaborative and eventually consistent distributed systems.

## Which service use it?



-   **Collaborative Text Editors:** Applications like Google Docs or Atom Teletype use CRDTs to allow multiple users to edit the same document concurrently without explicit locking, and changes are merged automatically.

-   **Shared Whiteboards and Drawing Applications:** CRDTs can manage concurrent drawing operations, ensuring that all users see a consistent final state.

-   **Distributed Counters and Sets:** In scenarios where multiple nodes need to update a shared counter or a set of items, CRDTs provide a way to do this without conflicts, ensuring eventual consistency.

-   **Multiplayer Games:** CRDTs can be used to synchronize game state across multiple players, especially for elements where concurrent updates are common and need to be resolved smoothly.

-   **Offline-First Applications:** Mobile or web applications that need to function offline and synchronize changes when connectivity is restored can leverage CRDTs for seamless merging of user-generated data.
