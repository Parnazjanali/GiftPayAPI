# High-Level Design (HLD) - Wallet and Discount Code System

## 1. Project Goal and Scope

This project aims to design and implement a scalable and reliable system for managing digital wallets and discount/gift recharge codes. The core objective is to handle a high volume of concurrent requests, particularly during promotional events, while ensuring fair and atomic credit allocation. The system comprises two independent services: a Wallet Service and a Discount Service, communicating via APIs.

## 2. Architectural Overview

[High-Level Architecture Diagram](docs/rawFiles/ActivateGiftCard)

**Key Components:**

* **Wallet Service:** Manages user balances and transactions. Provides an API to check balances and processes credit application requests from the Discount Service.
* **Discount Service:** Manages discount/gift recharge codes and tracks the first 1,000 successful redemptions. Provides real-time reporting and initiates credit application requests to the Wallet Service.

**Communication:** The Wallet Service and Discount Service will communicate via RESTful APIs over HTTP/HTTPS.

## 3. Key System Components

### 3.1. Wallet Service

* **Function:** Manages user wallet balances and transaction history.
* **Technology:** Golang, Fiber, SQLite, JWT (for session management)
* **Responsibilities:**
    * Storing and managing user wallet balances.
    * Processing credit and debit transactions.
    * Providing an API endpoint to check wallet balance (authenticated via JWT).
    * Receiving and processing credit application requests from the Discount Service.
    * Recording transaction history.

### 3.2. Discount Service

* **Function:** Manages discount/gift recharge codes and tracks redemption limits.
* **Technology:** Golang, Fiber, SQLite, JWT (for session management)
* **Responsibilities:**
    * Storing and managing discount/gift recharge codes.
    * Tracking the number of successful redemptions (limited to the first 1,000).
    * Providing an API endpoint to submit recharge codes (authenticated via JWT).
    * Communicating with the Wallet Service via API to request credit applications upon successful code redemption (within the limit).
    * Providing real-time reporting on the number of successful redemptions.

## 4. Request Processing Flow (Gift Code Redemption)

1.  A user submits a gift recharge code and their username (used as unique identifier) via an API endpoint exposed by the Discount Service.
2.  The Discount Service checks if the 1,000-user redemption limit has been reached.
3.  If the limit has not been reached and the code is valid, the Discount Service sends a request to the Wallet Service via its API to apply a credit of 1,000,000 Tomans to the user's wallet.
4.  The Wallet Service receives the credit application request, updates the user's balance, and records the transaction.
5.  The Discount Service records the successful redemption and increments the redemption count.
6.  If the redemption limit has been reached, subsequent redemption requests will be rejected by the Discount Service.

## 5. Non-Functional Requirements

* **High Performance:** The system must be able to handle a large volume of concurrent requests, especially during peak promotional periods.
* **Scalability:** The decoupled architecture allows for independent scaling of the Wallet and Discount Services based on their respective loads.
* **Reliability:** The system should be designed to minimize downtime and ensure data consistency, especially during concurrent credit applications.
* **Security:** JWT will be used for authenticating user requests to both services. API-level access control will protect sensitive endpoints.

## 6. Key Architectural Decisions

* **Decoupled Microservices:** The system is designed with two independent services to allow for better scalability, maintainability, and fault isolation.
* **RESTful APIs:** Communication between the Wallet and Discount Services will be through standard RESTful APIs over HTTP/HTTPS, ensuring interoperability.
* **Golang and Fiber:** Chosen for their high performance and efficiency in handling concurrent requests, aligning with the project's performance objectives.
* **SQLite:** Selected as the initial database for simplicity, acknowledging that a more robust and scalable database might be required for production environments.
* **JWT for Session Management:** Used for authenticating user requests, ensuring secure access to the APIs.

## 7. Potential Challenges and Risks

* **Concurrency Control:** Ensuring atomicity and preventing race conditions when applying credits during high-traffic events will be a critical challenge.
* **API Communication Reliability:** Implementing mechanisms to handle potential failures in communication between the Wallet and Discount Services.
* **Scalability of SQLite:** While suitable for development, the scalability limitations of SQLite might necessitate a migration to a more robust database in the future.
* **Manual Deployment:** The lack of an automated CI/CD pipeline introduces potential risks in terms of deployment consistency and speed.

---

**Note:** This HLD provides a high-level overview of the system. A visual architecture diagram would further enhance understanding. Low-Level Design (LLD) documents will detail the specific implementation choices, database schemas, API contracts, and concurrency control mechanisms.