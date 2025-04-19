# 🧩 Microservices Monorepo – GiftCard, Wallet, API Gateway

## Project Overview

This project showcases a microservices-based architecture designed for managing **Gift Cards**, **Wallet balances**, and handling authentication via an **API Gateway**.

Driven by the need to overcome the challenges of scaling and managing large, monolithic applications, this architecture embraces independent development and HTTP communication between services. This approach fosters scalability, modularity, and clear separation of responsibilities, all while maintaining well-defined domain boundaries.

The services are built using **Go (Golang)** with the **Fiber** web framework, and JWT is implemented for authentication. Go's performance and concurrency capabilities, combined with the benefits of a microservices architecture, make this project a robust foundation for building high-performance and resilient systems.

---
## ✨ Why This Project? In a Nutshell

This project was born out of the challenges of scaling and managing large, monolithic applications. The goal is to showcase the power and flexibility of a **microservices architecture** built with **Go (Golang)** to create high-performance and resilient systems.

Why this architecture?

- **Scalability**: Ability to independently scale each service based on demand.
- **Separation of Concerns**: Each service (Gift Cards, Wallets, API Gateway) has a specific responsibility, simplifying maintenance and development.
- **Go's Performance**: Leveraging Goroutines for high concurrency with minimal overhead.
- **Security with JWT**: Secure and stateless authentication between services.
- **Future Extensibility**: Easily add new features without impacting existing components.

In short, this project is an exploration of building a scalable, maintainable, and high-performance system using microservices and Go.

---
## 📁 Project Structure
api contract: [click me](docs/apiDocs/apiContracts/apiGateway/contract.md)

wallet contract : [click me](docs/apiDocs/apiContracts/walletManager/contract.md)

gift card contract : [click me](docs/apiDocs/apiContracts/giftCardManager/contract.md)

Manifest : [clickme](docs/manifest/manifest.md)

Activate Giftcard callFlows : [click me](docs/rawFiles/activateGIftCard.drawio)

Transaction Log callFlows : [click me](docs/rawFiles/transactioLog.drawio)

SwaggerFiles : [click me](docs/swaggerFiles/swagger.json)

Each microservice has its own environment config, Dockerfile, and dedicated routes.

## 🔁 Call Flow & Contracts

Each component exposes RESTful endpoints with clearly defined request/response structures.

- Contracts (JSON schemas) are documented per endpoint.
- All inter-service communication is through the API Gateway.
- JWT Token is required for all protected routes.
- Internal errors follow a consistent error model.

## 🚀 Running the Project

### Requirements

- Go 1.21+
- Docker & Docker Compose

### Run with Docker Compose

```bash
docker-compose up --build
```


### 🙌 Author
Made by Parnaz Janali

