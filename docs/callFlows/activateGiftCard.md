### **Call Flow Description:**

1. **User Login (Initial Phase):**
   - The user attempts to log in by sending their **username** and **password** in the request body to the **API Gateway**.
   
2. **API Gateway Validates Credentials:**
   - The **API Gateway** validates the user's credentials.
   - If the credentials are valid, it generates a **JWT token** and returns it to the user.

3. **Request to Link Gift Card:**
   - The user sends a **request** to link a **gift card**, placing the **JWT token** in the header and the **Gift Card ID** in the Path.
   - The **API Gateway** validates the token and extracts the **User ID** from the JWT token.

4. **Retrieving Wallet ID:**
   - The **API Gateway** sends a request to the **Wallet Manager** to retrieve the **Wallet ID** using the extracted **User ID**.

5. **Wallet Manager Returns Wallet ID:**
   - The **Wallet Manager** returns the **Wallet ID** associated with the user.

6. **Linking Gift Card:**
   - The **API Gateway** sends a request to the **Gift Card Manager** to link the provided **Gift Card ID** with the obtained **Wallet ID**.

7. **Gift Card Manager Confirms Linking:**
   - The **Gift Card Manager** links the **gift card** to the wallet and confirms the operation.

8. **Crediting the Wallet:**
   - The **API Gateway** sends a request to the **Wallet Manager** to credit the wallet with **1,000,000 Toman**.

9. **Wallet Manager Completes the Transaction:**
   - The **Wallet Manager** completes the transaction and confirms the credit to the **API Gateway**.

10. **Final Response to User:**
   - The **API Gateway** sends a final response to the user, confirming that the **gift card was activated** and providing the updated **wallet balance**.
