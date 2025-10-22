# Event-Driven System Mode

## Core

An **event-driven** system mode is an architectural pattern in which the flow of the system is determined by events. Events are messages that are sent by one component to signal that something has happened, such as a user action, a sensor reading, or a change in the state of the system.

In an event-driven system, components are loosely coupled and communicate with each other by producing and consuming events. This is in contrast to a traditional request-response model, where components are tightly coupled and communicate with each other by making direct calls.

## How It Works

Event-driven systems typically consist of three main components:

-   **Event Producers:** Components that generate and send events.
-   **Event Consumers:** Components that receive and process events.
-   **Event Bus (or Message Broker):** A central component that receives events from producers and delivers them to consumers.

When an event producer sends an event, it is published to the event bus. The event bus then delivers the event to all of the consumers that have subscribed to that type of event. This allows for a great deal of flexibility, as producers and consumers do not need to have any knowledge of each other.

## Pros & Cons

### Pros

-   **Loose Coupling:** Components in an event-driven system are loosely coupled, which makes it easier to develop, test, and maintain the system.
-   **Scalability:** Event-driven systems are highly scalable, as new consumers can be added to the system without affecting the existing components.
-   **Resilience:** If one component in an event-driven system fails, the rest of the system can continue to operate.

### Cons

-   **Complexity:** Event-driven systems can be more complex than traditional request-response systems, as it can be difficult to reason about the flow of events through the system.
-   **Debugging:** Debugging event-driven systems can be challenging, as it can be difficult to trace the path of an event through the system.
-   **Eventual Consistency:** Event-driven systems are often eventually consistent, which means that there can be a delay between when an event is produced and when it is consumed. This can be a problem for applications that require strong consistency.

## Which service use it?

-   **Microservices Architectures:** Event-driven communication is a cornerstone of many microservices deployments, allowing services to communicate asynchronously and react to changes in other services.
-   **Real-time Data Processing (e.g., IoT platforms, financial trading systems):** Systems that need to process and react to a continuous stream of data often use event-driven patterns.
-   **User Activity Tracking and Analytics:** Websites and applications can generate events for user actions (clicks, views, purchases) that are then consumed by analytics systems.
-   **Serverless Computing (e.g., AWS Lambda, Google Cloud Functions):** Serverless functions are often triggered by events from various sources like database changes, file uploads, or message queues.
