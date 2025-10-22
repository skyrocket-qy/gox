# Message Queue Communication

This section describes Message Queue communication, an asynchronous communication pattern where messages are stored in a queue until they are processed by a receiver, enabling decoupling and reliable message delivery.

## Which service use it?



-   **Asynchronous Task Processing:** Sending emails, generating reports, processing images, or performing other time-consuming tasks in the background without blocking the user interface.

-   **Decoupling Microservices:** Allowing different microservices to communicate without direct dependencies, improving resilience and scalability.

-   **Buffering and Load Leveling:** Queues can absorb bursts of requests, protecting backend services from being overwhelmed during peak traffic.

-   **Long-Running Batch Jobs:** Distributing large batch processing tasks across multiple workers, where each worker picks up a message (a unit of work) from the queue.

-   **Order Processing Systems:** Ensuring that orders are processed in a reliable and sequential manner, even if some components are temporarily unavailable.
