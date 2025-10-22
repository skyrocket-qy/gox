# Topology

## Core

In the context of a distributed system, **topology** refers to the arrangement of the nodes and the communication links between them. The network topology has a significant impact on the performance, reliability, and scalability of the system.

Different topologies have different characteristics. For example, some topologies are more fault-tolerant than others, while some are more scalable. The choice of which topology to use depends on the specific requirements of the application.

### Common Network Topologies

There are a number of different network topologies that can be used in a distributed system. Some of the most common ones include:

- **Bus:** All nodes are connected to a single, shared communication link. This is a simple topology, but it has a single point of failure.
- **Star:** All nodes are connected to a central hub. This is a more fault-tolerant topology than the bus, but the hub can become a bottleneck.
- **Ring:** All nodes are connected in a closed loop. This is a simple and reliable topology, but it can be difficult to add and remove nodes.
- **Mesh:** All nodes are connected to all other nodes. This is the most fault-tolerant topology, but it is also the most expensive to build and maintain.
- **Tree:** A hybrid topology that combines the bus and star topologies. This is a scalable and flexible topology, but it can be complex to manage.

The choice of which topology to use is a trade-off between a number of factors, including cost, performance, reliability, and scalability.

## Comparison

| Topology | Reliability | Scalability | Cost | Use Case |
|---|---|---|---|---|
| **[Bus](./bus)** | Low | Low | Low | Small networks |
| **[Star](./star)** | Medium | Low | Medium | Local area networks (LANs) |
| **[Ring](./ring)** | High | Low | Medium | Telecom networks |
| **[Mesh](./mesh)** | High | High | High | Wide area networks (WANs) |
| **[Tree](./tree)** | Medium | High | High | Large networks |

## Which service use it?



-   **Bus Topology:** Historically used in early Ethernet networks (e.g., 10Base2, 10Base5 coaxial cables) and still found in some industrial control systems or embedded networks.

-   **Star Topology:** Widely used in modern Local Area Networks (LANs) where all devices connect to a central switch or hub. Most home and office networks are star topologies.

-   **Ring Topology:** Historically used in Token Ring networks and some fiber optic networks. Also found in some Storage Area Networks (SANs) and metropolitan area networks (MANs).

-   **Mesh Topology:** Employed in critical infrastructure like military communications, backbone networks of the internet, and some wireless sensor networks where high redundancy and fault tolerance are paramount.

-   **Tree Topology:** Often used in large corporate networks, combining multiple star networks into a hierarchical structure, allowing for easy expansion and management.
