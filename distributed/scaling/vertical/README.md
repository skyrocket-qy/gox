# Vertical Scaling

This section explains Vertical Scaling (scaling up), a method of increasing capacity by adding more resources (CPU, RAM, storage) to an existing machine in a distributed system.

## Which service use it?



-   **Traditional Monolithic Applications:** Many legacy applications that were not designed for distributed environments often rely on vertical scaling by running on increasingly powerful single servers.

-   **Single-Instance Relational Databases:** Databases like Oracle, SQL Server, or even a standalone PostgreSQL/MySQL instance are often vertically scaled by upgrading the server's CPU, RAM, and storage to handle more load.

-   **Specialized High-Performance Computing (HPC) Tasks:** Certain computational workloads that are difficult to parallelize across multiple machines might benefit most from running on a single, extremely powerful server with a large amount of memory and many CPU cores.

-   **In-Memory Databases:** Databases designed to operate primarily in RAM (e.g., SAP HANA) often require significant vertical scaling to accommodate large datasets.
