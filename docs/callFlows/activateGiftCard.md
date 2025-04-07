
# Call Flow Description - Gift Card Redemption

## Overview
This call flow describes the process of **redeeming a gift card** and crediting its balance to the user's wallet through the API Gateway.

## Call Flow Steps

### **1. User Authentication**
- The **user** provides their **username and password** to the **API Gateway**.
- The **API Gateway** forwards the credentials to the **Profile Manager** for verification.

### **2. Token Generation**
- If the credentials are valid, the **Profile Manager** generates a **JWT token** and sends it back to the **API Gateway**.
- The **API Gateway** then returns the **JWT token** to the user.

### **3. Gift Card Redemption Request**
- The **user** sends a request to the **API Gateway**, including:  
  - The **JWT token** in the request header  
  - The **Gift Card ID** in the request body  

### **4. Token Validation**
- The **API Gateway** forwards the request to the **Profile Manager** for token validation.
- The **Profile Manager** verifies the token and returns the **User ID**.

### **5. Retrieve Wallet ID**
- The **API Gateway** requests the **Wallet Manager** to retrieve the **Wallet ID** associated with the **User ID**.
- The **Wallet Manager** returns the corresponding **Wallet ID**.

### **6. Link Gift Card to Wallet**
- The **API Gateway** sends a request to the **Gift Card Manager** to **link the Gift Card ID** to the retrieved **Wallet ID**.

### **7. Confirm Gift Card and Credit Amount**
- The **Gift Card Manager** verifies the **gift card** and determines the **amount to be credited**.
- The **API Gateway** instructs the **Wallet Manager** to credit the specified amount (e.g., **100,000 Toman**) to the **wallet**.

### **8. Wallet Balance Update**
- The **Wallet Manager** confirms that the wallet has been **successfully credited**.
- The **API Gateway** sends a **response** to the user, confirming the **new wallet balance**.

## Notes
- The **JWT token** is required for authentication.  
- If the **token is invalid**, the request is rejected.  
- If the **gift card is already used**, the request is denied.  
- If the **wallet update fails**, an error message is returned.  
