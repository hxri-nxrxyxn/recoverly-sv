# **API Documentation**

### **1. /login**
### URL: `/api/v1/login`
#### METHOD: `POST`
The `/login` endpoint allows users to authenticate by providing their credentials (usually username and password). Upon successful authentication, a token or session will be returned.

##### Request Body:
The request body must be in `JSON` format and contain the following fields:

| Field       | Type     | Description                                |
|-------------|----------|--------------------------------------------|
| `email`  | `string` | The email of the user.         |
| `password`  | `string` | The password of the user.                  |

#### **Example Request:**

```json
{
  "email": "john_doe@example.com",
  "password": "pass123"
}
```
```bash
curl --location 'https://hardly-genuine-yeti.ngrok-free.app/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
   "email":"john_doe@email.com",
   "password":"pass123"
}'
```

### **Success Response**
#### Status Code: `200 OK`
Response Body:

``` json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "1",
    "name": "John Doe",
    "email": "user@example.com"
  },
  "expires_in": 3600
}
```

``` bash
curl --location 'https://hardly-genuine-yeti.ngrok-free.app/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "1",
    "name": "John Doe",
    "email": "user@example.com"
  },
  "expires_in": 3600 } '
  ```
 ### **Error Responses**
 #### Status Code: `400 Bad Request`
 Malformed or missing fields.
 
Example Response:
```json
{
  "error": "Email or password is missing."
}
```
#### Status Code: `401 Unauthorized`

Invalid email or password.
Example Response:
```json
{
  "error": "email or password is incorrect."
}
```
#### **The endpoint is defined in Go as follows:**

```go 
Copy
api.Post("/login", controller.Login(db))
```
##### **Handler: `controller.Login(db)`**
This handler manages the login process by checking the provided credentials against the database and returning a token if valid.


### **2. /register**

### URL: `/api/v1/register`
#### METHOD: `POST`
The `/register` endpoint allows users to create a new account by providing their email, password, and other details. Upon successful registration, the user is returned with a confirmation message and their user data.

##### Request Body:
The request body must be in `JSON` format and contain the following fields:

| Field       | Type     | Description                                    |
|-------------|----------|------------------------------------------------|
| `email`     | `string` | The email of the user.                         |
| `password`  | `string` | The password of the user.                      |
| `name`      | `string` | The full name of the user.                     |

#### **Example Request:**

```json
{
  "email": "john_doe@example.com",
  "password": "pass123",
  "name": "John Doe"
}
```
``` bash
curl --location '/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "john_doe@example.com",
  "password": "pass123",
  "name": "John Doe"
}'
```

### **Success Response**
#### Status Code: `201 Created`
##### Returns a success message and the newly created user.
Response Body:
```json
{
  "message": "User registered successfully.",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john_doe@example.com"
  }
}
```
```bash
curl --location '/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
  "message": "User registered successfully.",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john_doe@example.com"
  }
}'
```
### **Error Responses**
#### Status Code: `400 Bad Request`
##### Missing or malformed fields in the request. 
Example Response:
```json
{
  "error": "Email, password, or name is missing."
}
```
#### Status Code: `409 Conflict`
##### The email provided is already in use. 
Example Response:
```json
{
  "error": "Email already exists."
}
```
##### The endpoint is defined in Go as follows:
``` go
api.Post("/register", controller.CreateUser(db))
```
#### Handler: `controller.CreateUser(db)`
##### This handler manages the registration process by validating the provided data, checking if the email is already in use, and then saving the new user to the database.




### **3. /users**
### **URL:** `api/v1/users`
#### METHOD: `GET`

The `/users` endpoint allows you to retrieve a list of all registered users. It provides the user information such as `id`, `name`, and `email`.

#### **Example Request:**

```bash
curl --location '/api/v1/users' \
--header 'Content-Type: application/json'
```
**Response format :**
| Field     | Type     | Description                                |
|-----------|----------|--------------------------------------------|
| `id`      | `int`    | The unique identifier for the user.        |
| `name`    | `string` | The name of the user.                      |
| `email`   | `string` | The email address of the user.             |

### **Success Response**
#### Status Code: `200 OK`
Response Body:
```json
[
  {
    "id": 1,
    "name": "Hari Doe",
    "email": "hari_doe@example.com"
  },
  {
    "id": 2,
    "name": "Johhny Sebastian",
    "email": "janeSeban@email.com"
  }
]
```
```bash
curl --location 'https://your-api-url.com/api/v1/users' \
--header 'Content-Type: application/json'
```
### **Error Responses**
#### Status Code: `500 Internal Server Error`
An error occurred while fetching the users from the database.

Example Response:
```json
{
  "error": "Unable to fetch users."
}
```
#### The endpoint is defined in Go as follows:
```go
api.Get("/users", controller.GetUsers(db))
```
#### **Handler: controller.GetUsers(db)**
 The handler controller.GetUsers(db) is responsible for querying the database and returning the results.


 
 





