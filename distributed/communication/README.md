# Communication

## Core

In distributed systems, **communication** is the mechanism by which processes running on different nodes interact and exchange information. Effective communication is fundamental to the functioning of any distributed system, as it enables coordination, data sharing, and the execution of distributed algorithms.

This section explores various communication patterns, protocols, and models used for inter-process communication. Key challenges in distributed communication include handling network latency, ensuring message delivery, maintaining order, and dealing with failures.

## Comparison

| Pattern | Coupling | Synchronicity | Topology | Use Case |
|---|---|---|---|---|
| **[Client-Server](./client-server)** | Tight | Synchronous | Centralized | Web services, databases |
| **[Message Queue](./message-queue)** | Loose | Asynchronous | Decoupled | Task processing, microservices |
| **[Publish-Subscribe](./pubsub)** | Loose | Asynchronous | Decoupled | Event-driven systems, notifications |
| **[Peer-to-Peer (P2P)](./p2p)** | Loose | Both | Decentralized | File sharing, content delivery |
| **[Actor Model](./actor-model)** | Loose | Asynchronous | Decentralized | Concurrent and parallel systems |

## Which service use it?



-   **Client-Server:** The most common communication pattern, used in virtually all web applications, email systems, file servers, and traditional enterprise software.

-   **Message Queue:** Asynchronous task processing, decoupling microservices, buffering requests, and enabling reliable communication between distributed components.

-   **Publish-Subscribe:** Real-time data feeds, event notifications, chat applications, and IoT data ingestion.

-   **Peer-to-Peer (P2P):** File sharing networks, cryptocurrencies, distributed content delivery, and some online gaming.

-   **Actor Model:** Concurrent and parallel systems, especially those requiring high concurrency and fault tolerance, such as Erlang/OTP applications, Akka-based systems, and some distributed databases.
