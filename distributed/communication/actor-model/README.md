# Actor Model

This section describes the Actor Model, a concurrency model for distributed systems where "actors" are the universal primitives of concurrent computation, communicating via asynchronous message passing.

## Which service use it?



-   **Erlang/OTP Applications:** Erlang is a programming language built around the Actor Model, and its OTP (Open Telecom Platform) framework is widely used for building highly concurrent, fault-tolerant distributed systems in telecommunications, messaging, and other industries.

-   **Akka Framework (Scala, Java):** Akka provides an implementation of the Actor Model for the JVM, enabling developers to build scalable and resilient concurrent applications in Scala and Java.

-   **Microsoft Orleans (.NET):** Orleans is a framework that implements the Actor Model for building distributed, high-scale computing applications in .NET.

-   **Some Distributed Databases:** Certain distributed databases or data processing systems might use actor-like principles internally for managing concurrency and communication between components.

-   **High-Concurrency Systems:** Any system requiring a high degree of concurrency, fault tolerance, and distributed processing can benefit from the Actor Model, such as real-time analytics, gaming servers, and IoT platforms.
