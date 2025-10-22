# Peer-to-Peer (P2P) Mode

## Core

A **Peer-to-Peer (P2P)** system mode is a decentralized architectural model where all participating nodes, called peers, have equivalent capabilities and responsibilities. Unlike the client-server model, there is no central server. Peers communicate directly with each other to share resources, workloads, and data.

Each peer can act as both a client (requesting services) and a server (providing services). This creates a resilient and scalable network where the failure of a single peer does not bring down the entire system.

## How It Works

In a P2P network, peers need a way to discover each other. This can be done in several ways:
-   **Centralized Discovery:** A central server (a tracker or bootstrap server) maintains a list of active peers. New peers connect to this server to get a list of other peers they can connect to. While discovery is centralized, the actual data exchange is P2P.
-   **Decentralized Discovery:** Peers discover each other using protocols like Distributed Hash Tables (DHTs). In a DHT, each peer is responsible for storing a small portion of the routing information, creating a decentralized index for the entire network.

Once connected, peers can exchange information directly. For example, in a file-sharing application, a peer can download parts of a file from multiple other peers simultaneously, increasing download speed.

P2P networks can be:
-   **Unstructured:** Peers connect to each other in an ad-hoc manner. Finding specific data can be inefficient as requests may need to be flooded across the network.
-   **Structured:** Peers are organized in a specific topology (e.g., a ring or a tree), and data is placed at specific locations, making searches very efficient. DHTs are an example of structured P2P networks.

## Pros & Cons

### Pros

-   **Scalability:** The total capacity of the system increases as more peers join the network, since each peer adds new resources.
-   **Resilience and Fault Tolerance:** There is no single point of failure. The system can continue to operate even if some peers go offline.
-   **Cost-Effectiveness:** P2P networks do not require powerful and expensive central servers, reducing infrastructure costs.
-   **Censorship Resistance:** The decentralized nature makes it difficult for any central authority to shut down or censor the system.

### Cons

-   **Discovery and Coordination:** Finding peers and coordinating their actions can be complex without a central server.
-   **Data Availability:** Data is only available if the peers storing it are online. If all peers with a piece of data go offline, that data becomes inaccessible.
-   **Security:** P2P networks can be vulnerable to security threats such as malicious peers, data poisoning, and denial-of-service attacks.
-   **Inconsistent Performance:** The performance of the network can be unpredictable and depends on the number of peers online and their network conditions.

## Which service use it?

-   **File Sharing Networks (e.g., BitTorrent):** Users directly share files with each other without a central server mediating the transfers.
-   **Cryptocurrencies (e.g., Bitcoin, Ethereum):** The blockchain is maintained and validated by a decentralized network of peer nodes, where each node holds a copy of the ledger and participates in transaction verification.
-   **Distributed Content Delivery Networks (CDNs):** Some CDNs use P2P principles to distribute content more efficiently, allowing users to download content from nearby peers.
-   **Some Online Gaming Platforms:** Certain games use P2P connections for direct communication between players, especially for real-time multiplayer experiences.
