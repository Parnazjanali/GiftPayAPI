# Project Manifest: Wallet and Discount Code System

## 1. Project Overview

This project aims to design and implement a **scalable and reliable** system for managing digital wallets and discount/gift recharge codes.

## 2. Key Objectives

* Build a **high-performance** system ⚡ to handle a large volume of concurrent requests.
* Enable users to check their **wallet balance** via API.
* Design a **decoupled architecture** with two independent services (Wallet and Discount) communicating through APIs.

## ❓ 3. Problem Statement

During high-traffic promotional events, a limited number of gift recharge codes are made available to users. The first 1,000 users who submit a valid code will receive a wallet credit of 1,000,000 Tomans.

The challenge lies in handling a large volume of concurrent requests while ensuring that the credit allocation remains fair, atomic, and free of race conditions. The system must be able to handle this traffic spike without overloading or creating inconsistencies in user balances.

## 4. Proposed Solution

The system architecture consists of two independent services:

* **Wallet Service:** Manages user balances and transactions, provides an API to check balances, and receives credit application requests from the Discount Service.
* **Discount Service:** Manages discount and gift recharge codes, tracks the first **1,000 users** to redeem the credit, provides real-time reporting, and requests credit applications from the Wallet Service.

## 5. Request Processing Flow

1. The user submits a gift recharge code and phone number via API.
2. The Discount Service checks the **1,000-user limit**.
3. If valid, the Discount Service requests the Wallet Service to apply the credit ✅.
4. The Wallet Service updates the user's balance and records the transaction.
5. Requests beyond the first **1,000 users** are rejected ❌.

## ✅ 6. Assumptions

* **Username and Password** are used as the unique user identifier 🆔.
* The system must be **highly scalable and reliable**.
* Services are independent and communicate via **APIs**.

## ⚙️ 7. Technical Stack

* **Programming Language:** Golang
* **API Framework:** Fiber
* **Database:** SQLite
* **API Documentation:** Swagger
* **Containerization:** Docker
* **Session Management:** JWT

## 8. Security

* JWT token verification (including expiration check) on each request to ensure user identity.
* API-level responsibility for access control and protection of sensitive routes.

## 🚀 9. Testing and Deployment

* Dockerized architecture for easy local development and containerized deployment.
* Project hosted on Git for version control and collaboration.
* No automated CI/CD pipeline is implemented at this stage — deployments are managed manually.

## 10. Out of Scope

* No complex authentication system required ❌.
* No graphical user interface (GUI) needed.
* Profile manager will not be implemented.