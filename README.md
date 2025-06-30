# Go Microservices Architecture

This project was built to learn and practice **Golang** microservices architecture, event-driven communication patterns, and containerized deployments using Docker.

## Architecture Overview

A distributed system consisting of 6 microservices communicating via HTTP APIs, message queues, and gRPC. Each service handles a specific domain responsibility following microservices principles.

```
Frontend → Broker Service → Authentication/Logger/Mail Services
                ↓
            RabbitMQ → Listener Service
```

## Services

### Broker Service
**Port:** 8080  
Central API gateway that orchestrates communication between other services. Routes incoming requests to appropriate microservices and handles service-to-service communication via HTTP and RabbitMQ.

### Authentication Service  
**Port:** 8081  
Handles user authentication and authorization. Manages user registration, login, password hashing, and JWT token generation. Uses PostgreSQL for persistent user data storage.

### Logger Service
Centralized logging service that collects and stores application logs from all services. Implements both HTTP API and gRPC server for log ingestion. Uses MongoDB for log storage and aggregation.

### Mail Service
Email service responsible for sending transactional emails (welcome emails, password resets, notifications). Templates are processed server-side with HTML and plain text formats.

### Listener Service
Event consumer that listens to RabbitMQ message queues. Processes asynchronous events from other services and triggers appropriate business logic based on message types.

### Frontend Service
Web interface built with Go templates. Provides user interface for interacting with the microservices ecosystem through API calls to the broker service.

## Technology Stack

### Core Languages & Frameworks
- **Go 1.18** - Primary programming language
- **Chi Router v5** - HTTP routing and middleware
- **Go Templates** - Frontend templating

### Databases & Storage
- **PostgreSQL** - User authentication data
- **MongoDB** - Centralized logging storage
- **Docker Volumes** - Data persistence

### Message Queue & Communication
- **RabbitMQ** - Asynchronous message passing
- **gRPC** - High-performance service communication
- **Protocol Buffers** - Message serialization
- **HTTP/REST** - Service-to-service communication

### Email & Development Tools
- **go-simple-mail** - SMTP email client
- **MailHog** - Email testing in development
- **go-premailer** - HTML email processing

### Database Drivers
- **pgx/v4** - PostgreSQL driver and connection pooling
- **mongo-driver** - MongoDB official driver

### DevOps & Deployment
- **Docker & Docker Compose** - Containerization
- **Multi-stage Dockerfiles** - Optimized container builds
- **Makefile** - Build automation

