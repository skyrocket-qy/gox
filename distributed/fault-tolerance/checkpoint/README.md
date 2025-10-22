# Checkpointing

This section explains Checkpointing as a fault-tolerance technique, involving periodically saving the state of a distributed computation to stable storage to enable rollback recovery.

## Which service use it?



-   **Long-Running Batch Processing Jobs (e.g., Apache Spark, Apache Flink):** These frameworks periodically save the state of computations, allowing them to resume from the last successful checkpoint if a worker node fails, rather than restarting the entire job.

-   **Scientific Simulations:** Complex and time-consuming scientific simulations often use checkpointing to save their progress, enabling recovery from system crashes or allowing the simulation to be paused and resumed later.

-   **Stateful Stream Processing Applications:** Stream processing engines (e.g., Apache Flink, Apache Kafka Streams) use checkpointing to persist the state of their operators, ensuring fault tolerance and exactly-once processing guarantees.

-   **Distributed Databases:** Some distributed databases use checkpointing internally to periodically flush in-memory changes to disk, reducing recovery time after a crash.

-   **Virtual Machine (VM) Snapshots:** While not exactly the same, the concept of taking a snapshot of a running VM's state is analogous to checkpointing, allowing the VM to be restored to a previous point in time.
