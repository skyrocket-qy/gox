# Active-Active Cluster

This section describes the Active-Active Cluster configuration for fault tolerance, where multiple nodes are simultaneously active and capable of handling requests, providing high availability and improved performance.

## Which service use it?



-   **Load-Balanced Web Servers:** Multiple web servers running concurrently behind a load balancer, all actively serving user requests. If one server fails, the load balancer redirects traffic to the remaining active servers.

-   **Stateless Microservices:** Microservices designed to be stateless can easily be deployed in an active-active configuration, as any instance can handle any request without needing to maintain session-specific data locally.

-   **Distributed Caching Systems:** Caching solutions like Redis Cluster or Memcached can operate in an active-active manner, distributing cached data across multiple nodes, all of which are available for reads and writes.

-   **Multi-Master Databases:** Databases configured for multi-master replication (e.g., some NoSQL databases, PostgreSQL BDR) allow multiple nodes to accept writes simultaneously, forming an active-active setup for data modification.

-   **Global Traffic Management (GTM) for Geo-Redundancy:** Systems that direct user traffic to the closest or healthiest active data center, where each data center is an active cluster.
