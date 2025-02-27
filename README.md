# README

## Overview

My self-study project demonstrates a microservices architecture using a **Gateway**, **Consul** for service discovery and load balancing, and a **Micro Service** that interacts with both **Read** and **Write** storages. The two diagrams provided illustrate how these components work together.

## Architecture

### Diagram 1: Gateway, Consul for service discovery and load balancing

![Load Balancing Diagram](https://github.com/bhtoan2204/video-stream-backend/blob/main/assets/Youtube%20Backend%20Architecture-Load%20Balancing.drawio.png?raw=true)

#### Components

1. **Gateway**
   - Acts as a single entry point for external requests.
   - Forwards requests to Consul for service discovery and load balancing.
2. **Consul**
   - Maintains a registry of available service instances.
   - Performs **reverse proxy** functionality and **load balancing**.
   - Routes incoming requests to one of the **Micro-Service** instances.
3. **Microservice Service Instances**
   - Multiple instances run in parallel.
   - Microservices are being health-check to decide should it be distributed
   - Requests are distributed to these instances by Consul, ensuring scalability and high availability.
   - Consul also can be considered as Service Mesh in the future (not sure)

### Diagram 2: User Service with Read/Write Storage and Event Streaming

![CQRS Architecture](https://github.com/bhtoan2204/video-stream-backend/blob/main/assets/Youtube%20Backend%20Architecture-CQRS.drawio.png?raw=true)

#### Components

1. **User Service**

   - **Query Endpoint**: Fetches data from the **Read Storage** (Elasticsearch).
   - **Command Endpoint**: Writes or updates data in the **Write Storage** (MySQL).
   - **Event Consumer**: Listens to events published by Kafka (originating from Debezium).

2. **Read Storage (Elasticsearch)**

   - Optimized for fast queries and indexing.
   - Receives updates when data changes in MySQL (via events).

3. **Write Storage (MySQL)**

   - Primary system of record for data writes.
   - Changes are captured by **Debezium** and published to Kafka.

4. **Debezium**

   - Monitors MySQL transaction logs.
   - Publishes change events to **Kafka** for other services (e.g., User Service) to consume.

5. **Kafka**
   - Serves as the messaging backbone.
   - Decouples services by allowing asynchronous communication.

## Data Flow

1. **Command Flow**

   - A client sends a request to the **Gateway**.
   - The Gateway forwards it to **Consul**, which routes it to a **User Service** instance’s **Command Endpoint**.
   - The User Service writes data to **MySQL**.
   - **Debezium** captures changes from MySQL’s binlog and sends them to **Kafka**.
   - The **Event Consumer** in the User Service (or other interested services) processes these Kafka events.

2. **Query Flow**
   - A client sends a query request to the **Gateway**.
   - Consul routes it to a User Service instance’s **Query Endpoint**.
   - The Query Endpoint fetches data from **Elasticsearch** (the Read Storage).
   - Elasticsearch may be updated in near-real-time via events from Debezium -> Kafka -> User Service -> Elasticsearch.

## Some notice about Clean Architect in my project

- I have only been learning Clean Architecture for two weeks from now, so my implementation may not be perfect yet. However, I plan to fully adopt Clean Architecture principles and add unit tests in the future.

### **Application**

- The **Application layer** defines how the outside world (controllers, UI, or other services) interacts with the domain.
- By isolating commands, queries, and events, i can clearly separate **write** operations (commands) from **read** operations (queries), supporting a **CQRS** approach.

### **Domain**

- The **Domain layer** is the **Heart** of the application: it should remain free of external frameworks or infrastructure concerns.
- This ensures the **business Rules** do not depend on technical details like databases or message brokers.

### **Infrastructure**

- The **Infrastructure layer** holds all the technical details (DB, message brokers, logging).
- By keeping these details separate, i can swap or modify implementations without impacting the domain logic.
