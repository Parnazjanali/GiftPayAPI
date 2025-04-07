# ⚽ Project Manifest: Wallet and Discount Code System (Optimized) 

##  1. Project Overview

This project aims to design and implement a **scalable and reliable** system for managing digital wallets  and discount/gift recharge codes , specifically addressing the scenario of the **Iran vs. Brazil match** .

##  2. Key Objectives

* Build a **high-performance** system ⚡ to handle a large volume of concurrent requests .
* Implement a **fair, race-condition-free** mechanism ⚖️ to allocate gift credits to the first **1,000 users** .
* Provide **real-time reporting** on users who successfully redeemed the gift recharge.
* Enable users to check their **wallet balance** via API.
* Design a **decoupled architecture** with two independent services (Wallet and Discount) communicating through APIs.

## ❓ 3. Problem Statement

During the **Iran vs. Brazil final match** ️, a gift recharge code will be announced . The first **1,000 users** to enter this code will receive **1,000,000 Tomans** in wallet credit. The main challenge is managing the high volume of requests and ensuring fair credit allocation.

##  4. Proposed Solution

The system architecture consists of two independent services:

* **Wallet Service:** Manages user balances  and transactions , provides an API to check balances, and receives credit application requests from the Discount Service.
* **Discount Service:** Manages discount and gift recharge codes ️, tracks the first **1,000 users** to redeem the credit, provides real-time reporting, and requests credit applications from the Wallet Service.

##  5. Request Processing Flow

1.  The user submits a gift recharge code and phone number  via API.
2.  The Discount Service checks the **1,000-user limit** .
3.  If valid, the Discount Service requests the Wallet Service to apply the credit ✅.
4.  The Wallet Service updates the user's balance and records the transaction .
5.  Requests beyond the first **1,000 users** are rejected ❌.

## ✅ 6. Assumptions

* The phone number is used as the unique user identifier 🆔.
* The system must be **highly scalable and reliable** ️.
* Wallet transactions must be **atomic** ⚛️.
* Services are independent and communicate via APIs .
* Redis is used for concurrency management and the **1,000-user limit** .

## ⚙️ 7. Technical Stack

* **Programming Language:** Golang 
* **API Framework:** Fiber 
* **Database:** SQLite ️ (for initial development, adaptable to Redis)
* **Concurrency and Caching:** Redis ⚡ (using sorted lists for tracking the first **1,000 users** and counters for limit management)
* **Logging:** Zap Logger 
* **API Documentation:** Swagger 
* **Containerization:** Docker 
* **Session Management:** JWT  (user information stored in tokens, no Refresh Tokens in the first phase)

##  8. Security

* Input validation to prevent injection attacks .
* HTTPS for secure API communication .
* Restricted access to Redis and the database ️.
* Rate Limiting implementation to prevent DDOS attacks .

##  9. Data Management and Archiving

* User data is archived every **6 months** ️.
* Archived data is retained for **2 years** ⏳.
* Archived data will be stored as JSON or CSV files  in a separate storage.

## ⚠️ 10. Error Handling

* Detailed logging for debugging .
* Implementation of data recovery mechanisms in case of errors .
* Clear error messages for users .

##  11. Testing and Deployment

* Unit and integration testing for system reliability ✅.
* Docker for easy and fast deployment .
* CI/CD for automated testing and deployment .

##  12. Out of Scope

* No complex authentication system required ❌.
* No graphical user interface (GUI) needed ️.
* Profile manager will not be implemented .
* Refresh Tokens will not be implemented in phase 1 .