# GoNexus

### Scalable, Resilient, and Distributed Microservices in Go

GoNexus is a distributed system built using **Golang**, designed to demonstrate the power of **microservices, messaging, authentication, logging, and email processing**. The system follows best practices in microservices architecture, with services communicating via **REST API, RPC, gRPC, and AMQP (RabbitMQ)**.

---

## 🚀 Features

- **Microservices Architecture** – Loosely coupled, independent services.
- **Multiple Communication Methods** – REST API, RPC, gRPC, and AMQP (RabbitMQ).
- **Scalability & Resilience** – Easily deployable on **Docker Swarm** and **Kubernetes**.
- **Database Flexibility** – Uses **PostgreSQL**, **MongoDB**, and **RabbitMQ** for different services.
- **Secure Authentication** – Manages user authentication with hashed credentials.
- **Centralized Logging** – Logs system events to **MongoDB**.
- **Email Notification Service** – Sends formatted emails via an independent microservice.

---

## 🏗️ Microservices Overview

| Service Name               | Description                                                  |
| -------------------------- | ------------------------------------------------------------ |
| **Frontend Service**       | Serves web pages; interacts with other services.             |
| **Broker Service**         | A centralized gateway that routes requests to microservices. |
| **Authentication Service** | Handles user authentication using PostgreSQL.                |
| **Logging Service**        | Logs system events to MongoDB.                               |
| **Mail Service**           | Sends emails with preformatted templates.                    |
| **Listener Service**       | Listens for RabbitMQ messages and processes them.            |

---

## ⚙️ Technologies Used

- **Go (Golang)** – Core language for all microservices.
- **gRPC & REST API** – Microservices communication.
- **RabbitMQ (AMQP)** – Message-based event processing.
- **Docker & Kubernetes** – Containerization & orchestration.
- **PostgreSQL & MongoDB** – Relational & NoSQL database support.
- **Nginx** – Reverse proxy setup.

---

## 📦 Installation & Setup

### Prerequisites

Ensure you have the following installed:

- [Golang](https://golang.org/doc/install)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [Kubernetes (kubectl, minikube, or k3s)](https://kubernetes.io/docs/tasks/tools/)

### Clone the Repository

```sh
git clone https://github.com/Tanmoy095/GoNexus.git
cd gonexus

Start the Services with Docker Compose
docker-compose up --build

Deploy to Kubernetes
kubectl apply -f k8s/

🛠️ API Endpoints

1️⃣ Authentication Service (auth-service)
Login: POST /auth/login
Register: POST /auth/register
Validate Token: GET /auth/validate
2️⃣ Mail Service (mail-service)
Send Email: POST /mail/send
3️⃣ Logging Service (logging-service)
Log Event: POST /log
4️⃣ Broker Service (broker-service)
Route Requests: POST /broker

🏗️ Architecture Diagram

+------------+        +------------+
| Frontend   | -----> |  Broker    |
+------------+        +------------+
        |                 |
        v                 v
+----------------+    +------------------+
| Auth Service   |    |  Logging Service |
| (PostgreSQL)   |    |  (MongoDB)       |
+----------------+    +------------------+
        |                 |
        v                 v
+----------------+    +------------------+
| Mail Service   |    |  Listener Service|
| (SMTP, Templates) |  | (RabbitMQ)      |
+----------------+    +------------------+

🚀 Future Enhancements

✅ Implement JWT-based authentication for better security.
✅ Add Prometheus & Grafana for monitoring.
✅ Support multi-region deployment with Kubernetes.
✅ Introduce rate limiting & circuit breakers for resilience.
```
