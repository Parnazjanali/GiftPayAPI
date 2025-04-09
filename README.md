# 🧩 Microservices Monorepo – GiftCard, Wallet, API Gateway


##  Project Overview

This project is a microservices-based architecture designed for managing **Gift Cards**, **Wallet balances**, and handling authentication via an **API Gateway**.

Each service is developed independently and communicates with others via HTTP. The architecture is scalable, modular, and ideal for separating responsibilities while maintaining clean domain boundaries.

The services are written in **Go (Golang)** using the **Fiber** web framework, and JWT is used for authentication.

---

## 💡 Why This Project?

The motivation behind building this system:

- Practice and learn advanced backend concepts in Go.
- Implement domain-driven design with clear separation of concerns.
- Create a real-world scalable structure that can be used for productized services.
- Leverage API Gateway for centralizing access, security, and routing.

---

## 📁 Project Structure

```bash
.
├── api-gateway/
│   ├── main.go
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── service/
│   ├── utils/
│   └── Dockerfile
├── giftcard-manager/
│   ├── main.go
│   ├── config/
│   ├── handler/
│   ├── service/
│   └── Dockerfile
├── wallet-manager/
│   ├── main.go
│   ├── config/
│   ├── handler/
│   ├── service/
│   └── Dockerfile
├── docker-compose.yml
└── README.md
‍‍‍
```

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


## 📌 TODOs

- [ ] Add unit testing
- [ ] Add rate limiting middleware
- [ ] Migrate to PostgreSQL in production
- [ ] Add OpenAPI docs



## 🙌 Author

Made by Parnaz Janali 



