
# Gift Card Manager API Routes

## 1. Link Gift Card to Wallet

### Endpoint
- **Method:** `POST`
- **Route:** `/api/gift-card/link`
- **Description:**
    - This route is used to link a **gift card** to a **wallet**.
    - The request is sent by the **API Gateway** after validating the user's token and retrieving the **walletId** from **Wallet Manager**.
    - If the **gift card is valid and not already linked**, its amount will be confirmed and credited to the wallet.

---

### Request Example
```json
{
    "giftCardId": "abc123"
}
```

---

### Response Example

#### Success (`200 OK`)
```json
{
    "status": "success",
    "message": "Gift card linked successfully",
    "walletId": "67890",
    "giftCardId": "abc123",
    "amount": 100000
}
```

#### Unauthorized (`401 Unauthorized`)
```json
{
    "Code": "401",
    "Message": "Invalid or missing token"
}
```

#### Not Found (`404 Not Found`)
```json
{
    "Code": "404",
    "Message": "No gift card found with the provided ID"
}
```

#### Conflict (`409 Conflict`)
```json
{
    "Code": "409",
    "Message": "This gift card has already been linked to a wallet"
}
```

---

## Notes
- This request requires a **valid JWT token**; otherwise, a **401 Unauthorized** error will be returned.
- If the **gift card is invalid** or not found in the system, a **404 Not Found** error will be returned.
- If the **gift card has already been used**, the request will be rejected with a **409 Conflict** error.
