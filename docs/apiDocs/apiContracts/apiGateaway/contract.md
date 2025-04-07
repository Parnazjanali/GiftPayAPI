
# API Documentation

## Login Route

### Route Information
- **Path:** `/Login`
- **Method:** `POST`

### Headers
| Header         |Description|
|----------------|---------|
| Content-Type   | `application/json`|


### Request Body
```json
{
  "username": "string",
  "password": "string"
}
```

### Responses

#### Success (`200 OK`)
```json
{
  "token": "string",
  "expires_in": 3600
}
```
- **`token`**: JWT authentication token for subsequent requests
- **`expires_in`**: Token validity duration in seconds

### Errors

#### Bad Request (`400`)

```json
{
    "Code": "400",
    "Message": "Username or Password not provided."
}
```

#### Unauthorized (`401`)
```json
{
  
    "Code": "401",
    "Message": "Incorrect username or password."
  
}
```

#### Internal Server Error (`500`)
```json
{
    "Code": "1003",
    "Message": "Internal server error."
  
}
```

## Validate Token Route

### Route Information
- **Endpoint:** `/validate-token`
- **Method:** `POST`

### Headers
| Header         | Type    | Description                          |
|----------------|---------|--------------------------------------|
| Authorization  | string  | Bearer JWT Token                     |
| Content-Type   | string  | `application/json`                   |


### Response

#### Success (`200 OK`)
```json
{
  "user_id": "12345",
  "status": "valid"
}
```
- **`user_id`**: Unique identifier of the authenticated user
- **`status`**: Indicates whether the token is valid

### Errors

#### Unauthorized (`401`)
```json
{
  "error": {
    "code": "401",
    "message": "Invalid or expired token."
  }
}
```

#### Internal Server Error (`500`)
```json
{
  "error": {
    "code": "500",
    "message": "Internal server error."
  }
}
```


