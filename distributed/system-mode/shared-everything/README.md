# Shared-Everything Architecture

## Core

**Shared-everything** (also known as **shared-memory**) is a traditional computing architecture where a single, centralized system contains all the necessary components, including CPUs, memory, and storage, which are all shared and accessible by all parts of the system. This is the architecture found in most single-server systems, such as traditional mainframes and single-instance databases.

In this model, scaling is achieved by adding more resources to the single server, a process known as **vertical scaling** (or scaling up). This could involve adding more powerful CPUs, increasing the amount of RAM, or expanding the storage capacity.

## How It Works

In a shared-everything architecture, all processors (CPUs) have direct and uniform access to a single main memory and a shared pool of disk storage. The operating system manages access to these shared resources, ensuring that different processes and threads can coordinate and share data efficiently.

Data consistency is straightforward because all operations go through the same memory and storage controllers. The system can provide strong ACID (Atomicity, Consistency, Isolation, Durability) guarantees relatively easily.

This architecture is in contrast to distributed architectures like shared-nothing, where each node has its own private memory and storage.

## Pros & Cons

### Pros

-   **Simplicity:** The architecture is simpler to design, manage, and program for, as there is no need to deal with the complexities of distributed systems like network partitions, replication, or distributed consensus.
-   **Strong Consistency:** Strong data consistency is a natural feature of this model, making it well-suited for transactional applications (e.g., OLTP systems).
-   **High Performance for Local Operations:** Inter-process communication is very fast since it happens through shared memory, avoiding network latency.

### Cons

-   **Limited Scalability:** The system can only be scaled vertically. There is a physical limit to the amount of resources that can be added to a single machine, and the cost of high-end hardware increases exponentially.
-   **Single Point of Failure:** Since all components are in a single system, the failure of any critical component (like the motherboard or storage system) can bring the entire system down.
-   **Resource Contention:** As the number of processors and applications increases, there can be significant contention for shared resources like memory bandwidth and I/O channels, which can become a bottleneck.

## Which service use it?

-   **Traditional Single-Instance Relational Databases (e.g., Oracle, SQL Server, PostgreSQL, MySQL):** Before distributed versions became prevalent, these databases ran on a single powerful server, sharing its CPU, memory, and storage.
-   **Operating Systems on Multi-Core Machines:** Modern operating systems manage shared resources (CPU cores, RAM) for multiple processes and threads running on a single physical machine.
-   **High-Performance Computing (HPC) Systems with Shared Memory:** Many supercomputers and specialized HPC systems use shared-memory architectures for extremely fast inter-process communication within a single node.
-   **In-Memory Data Grids (within a single node):** Some in-memory data grids or caches might operate in a shared-everything manner on a single large server to maximize performance by avoiding network overhead.
