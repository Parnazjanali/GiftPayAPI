
# Call Flow Description - Retrieve Transaction History

## Overview
This call flow describes the process of retrieving the last **10 transaction logs** for a user through the API Gateway. 

## Call Flow Steps

1. **User Request**  
   - The user sends a request to the **API Gateway**, including a **JWT token** in the header.  
   - By default, the request retrieves the last **10 transactions**.

2. **Token Verification**  
   - The **API Gateway** forwards the request to the **Profile Manager** to validate the token.  

3. **User ID Retrieval**  
   - If the token is valid, the **Profile Manager** extracts the **User ID** and sends it back to the **API Gateway**.

4. **Request to Wallet Manager**  
   - The **API Gateway** sends a request to the **Wallet Manager**, asking for transactions related to the retrieved **User ID**.

5. **Transaction History Retrieval**  
   - The **Wallet Manager** finds the corresponding **Wallet ID** for the given **User ID**.  
   - It then retrieves the transaction history and returns the data to the **API Gateway**.

6. **Response to User**  
   - The **API Gateway** sends the **transaction logs** back to the user as a response.

## Notes
- The **JWT token** is required for authentication.  
- The default number of transactions returned is **10**, but this can be adjusted using query parameters.  
- If the **token is invalid**, the request is rejected.  
- If no transactions are found, an empty list is returned.
