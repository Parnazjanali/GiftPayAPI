
# Wallet Manager API Routes


## 1. Retrieve Wallet ID

### Endpoint
- **Method:** `GET`
- **Route:** `/wallet/:userId`
- **Description:**
    - Retrieves the wallet ID associated with a user based on the provided user ID.
    - This route is called by the **API Gateway** to fetch the corresponding wallet ID for subsequent operations.

### Request Example
```json
header
{
    "Authorization": "Bearer <JWT_TOKEN>"
}

body
{
    "userId": "username"
}
```

### Response

#### Success (`200 OK`)
```json
{
    "status": "success",
    "userId": "username",
    "walletId": "67890",
    "giftCardId": "abc123",
}
```

#### Unauthorized (`401 Unauthorized`)
```json
{
    "Code": "Unauthorized access",
    "message": "Invalid or missing token"
}
```

#### Not Found (`404 Not Found`)
```json
{
    "Code": "404",
    "message": "No wallet associated with the provided user ID"
}
```

---

## 2. Set Balance to Wallet

### Endpoint
- **Method:** `POST`
- **Route:** `/wallet/{walletId}/set-balance`
- **Description:**
    - Sets a specified balance to the user's wallet.
    - This route is called after the **Gift Card Manager** confirms the association of the gift card to the wallet, and the user requests to set a specific balance in their wallet.

### Request Example
```json
{
    "walletId": "67890",
    "balance": 100000
}
```

### Response Example

#### Success (`200 OK`)
```json
{
    "status": "success",
    "message": "Wallet balance updated successfully",
    "newBalance": 100000
}
```

#### Unauthorized (`401 Unauthorized`)
```json
{
    "Code": 400,
    "message": "walletId is required in the URL path"
}
```

#### Not Found (`404 Not Found`)
```json
{
    "Code": 404,
    "message": "No wallet associated with the provided wallet ID"
}
```

---

## Notes:
- All requests to this API must include a valid JWT token for authentication.
- If the token is missing or invalid, a `401 Unauthorized` error will be returned.
- If the wallet ID does not match any existing wallet, a `404 Not Found` error will be returned.

---



# Wallet Manager API - Transaction History Route

## 1. Retrieve Wallet Transactions

### Endpoint
- **Method:** `GET`
- **Route:** `/wallet/{walletId}/transactions`
- **Description:**
    - Retrieves the **last 10 transactions** related to a specific wallet by default.
    - If a different number of transactions is needed, a query parameter (`limit`) can be provided.
    - The **API Gateway** sends this request after validating the user's token through the **Profile Manager**.

---

### Request Example
**URL:**
```
GET /wallet/67890/transactions?limit=10
```

**Headers:**
```json
{
    "Authorization": "Bearer jwt_token_here"
}
```

---

### Response Example

#### Success (`200 OK`)
```json
{
    "status": "success",
    "walletId": "67890",
    "transactions": [
        {
            "transactionId": "tx001",
            "amount": 50000,
            "type": "credit",
            "timestamp": "2025-03-09T12:30:00Z"
        },
        {
            "transactionId": "tx002",
            "amount": 20000,
            "type": "debit",
            "timestamp": "2025-03-08T14:15:00Z"
        }
    ]
}
```

#### Unauthorized (`401 Unauthorized`)
```json
{
    "error": "Unauthorized access",
    "statusCode": 401,
    "message": "Invalid or missing token"
}
```

#### Not Found (`404 Not Found`)
```json
{
    "error": "Wallet not found",
    "statusCode": 404,
    "message": "No wallet found for the provided user ID"
}
```

---

## Notes
- The request **requires a valid JWT token** in the Authorization header.
- The default number of transactions returned is **10**, but this can be adjusted using the `limit` query parameter.
- If the wallet is **not found**, a `404 Not Found` error is returned.
