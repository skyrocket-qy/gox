# Horizontal Scaling

This section explains Horizontal Scaling (scaling out), a method of increasing capacity by adding more machines to a distributed system.

## Which service use it?



-   **Most Modern Web Applications:** Large-scale websites and online services (e.g., social media, e-commerce) distribute user traffic across many web servers and application servers.

-   **Microservices Architectures:** Microservices are inherently designed for horizontal scaling, allowing individual services to be scaled independently based on demand.

-   **Distributed Databases (NoSQL and Sharded SQL):** Databases like Apache Cassandra, MongoDB, Elasticsearch, and sharded relational databases achieve massive scale by distributing data and query processing across clusters of commodity servers.

-   **Big Data Processing Frameworks (e.g., Apache Hadoop, Apache Spark):** These frameworks process and analyze vast datasets by distributing computation across hundreds or thousands of nodes in a cluster.

-   **Cloud-Native Applications:** Applications built for cloud environments are typically designed to be horizontally scalable, leveraging auto-scaling groups and container orchestration platforms (e.g., Kubernetes) to dynamically adjust resources.
