# **API_Documentation : UserRoutes**

### **1. /login**
#### URL: `/api/v1/login`
### METHOD: `POST`
`POST /login`: This route is used to handle user login.

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

____________________________________________________________________________________________________________________________

### **2. /register**

#### URL: `/api/v1/register`
### METHOD: `POST`
`POST /register`: This route is used to handle user registration.

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


____________________________________________________________________________________________________________________________

### **3. /users**
#### **URL:** `api/v1/users`
### METHOD: `GET`
`GET /users`: This endpoint retrieves a list of all users.

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
curl --location GET '' \
--header 'Content-Type: application/json' \
--data-raw '[
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
]'
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

____________________________________________________________________________________________________________________________
 
### **4. /users/:id**

#### **URL:** `/users/:id`
### METHOD: `GET`

`GET /users/:id`: This endpoint retrieves the details of a specific user by their id. 

The `/users/:id` endpoint allows you to retrieve the details of a specific user by their `id`. It provides the user information such as `id`, `name`, and `email`.

#### **Example Request:**
```bash
curl --location 'https://your-api-url.com/api/v1/users/1' \
--header 'Content-Type: application/json'
```
 **Response format**
| Field     | Type     | Description                                |
|-----------|----------|--------------------------------------------|
| `id`      | `int`    | The unique identifier for the user.        |
| `name`    | `string` | The name of the user.                      |
| `email`   | `string` | The email address of the user.             |

### **Success Response**
#### Status Code: `200 OK`
Response Body:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john_doe@example.com"
}
```
```bash
curl --location --request GET '' \
--header 'Content-Type: application/json' \
--data-raw '{
  "id": 1,
  "name": "John Doe",
  "email": "john_doe@example.com"
}'
```
### **Error Responses**
#### Status Code: `404 Not Found`
The user with the given id was not found. 

Example Response:
```json
{
  "error": "User not found."
}
```
#### The endpoint is defined in Go as follows:
```go
api.Get("/users/:id", controller.GetUser(db))
```
#### Handler: `controller.GetUser(db)`
This handler fetches the user details from the database based on the provided id and returns them as a JSON response. If the user with the given id does not exist, an appropriate error message will be returned.

____________________________________________________________________________________________________________________________


### **5. /users/:id**
#### **URL:** `/users/:id`
### METHOD: `PATCH`
`PATCH /users/:id`: This endpoint is used to update a specific user's details.


The `/users/:id` endpoint allows you to update the details of a specific user by their `id`. It supports updating fields such as `name` and `email`.


#### **Example Request:**

```bash
curl --location --request PATCH '/api/v1/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "John Doe Updated",
  "email": "updated_email@example.com"
}'
```
#### **Request Body**

| Field     | Type     | Description                                         |
|-----------|----------|-----------------------------------------------------|
| `name`    | `string` | The updated name of the user (optional).            |
| `email`   | `string` | The updated email address of the user (optional).   |

#### **Response Format**

| Field     | Type     | Description                                |
|-----------|----------|--------------------------------------------|
| `id`      | `int`    | The unique identifier for the user.        |
| `name`    | `string` | The updated name of the user.              |
| `email`   | `string` | The updated email address of the user.     |

### **Success Response**
#### Status Code: `200 OK`
Response Body:
```json
{
  "id": 1,
  "name": "John Doe Updated",
  "email": "updated_email@example.com"
}
```
```bash
curl --location --request PATCH 'https://your-api-url.com/api/v1/users/1' \
--header 'Content-Type: application/json' \
--data-raw '{
  "id": 1,
  "name": "John Doe Updated",
  "email": "updated_email@example.com"
}'
```
### **Error Responses**
#### Status Code: `400 Bad Request`
Invalid data was provided for the update. 

Example Response:
```json
{
  "error": "Invalid data provided."
}
```
#### Status Code: `404 Not Found`
The user with the given id was not found. 

Example Response:
```json
{
  "error": "User not found."
}
```
#### The endpoint is defined in Go as follows:
```go
api.Patch("/users/:id", controller.UpdateUser(db))
```
#### Handler: `controller.UpdateUser(db)`
This handler processes the incoming request to update user details in the database. It validates the request body, updates the user, and returns the updated user information or an appropriate error message.

______________________________________________________________________________

# **API Documentation - ChatRoutes**

### **1. /chat**
#### **URL:** `/chat`
#### METHOD: `POST`
The `/chat` endpoint allows users to create a new chat message.


#### **Request Body**

| Field       | Type     | Description                                   |
|-------------|----------|-----------------------------------------------|
| `message`   | `string` | The content of the chat message.              |
| `sender`    | `string` | The identifier of the user sending the message. |
| `group`     | `string` | The group identifier for the chat (optional). |

#### **Response Format**

| Field       | Type     | Description                                   |
|-------------|----------|-----------------------------------------------|
| `id`        | `int`    | The unique identifier for the chat message.   |
| `message`   | `string` | The content of the created chat message.      |
| `sender`    | `string` | The identifier of the sender.                 |
| `group`     | `string` | The group identifier for the chat message.    |
| `createdAt` | `string` | The timestamp when the message was created.   |

#### **Example Request**

```bash
curl --location --request POST 'https://your-api-url.com/api/v1/chat' \
--header 'Content-Type: application/json' \
--data-raw '{
  "message": "Hello, world!",
  "sender": "User1",
  "group": "group1"
}'
```
### Success Response
#### Status Code: `201 Created`
Response Body:
```json
{
  "id": 1,
  "message": "Hello, world!",
  "sender": "User1",
  "group": "group1",
  "createdAt": "2025-01-18T10:00:00Z"
}
```
```bash
curl --location --request '' \
--header 'Content-Type: application/json' \
--data '{
  "id": 1,
  "message": "Hello, world!",
  "sender": "User1",
  "group": "group1",
  "createdAt": "2025-01-18T10:00:00Z"
}'
```
### Error Responses
#### Status Code: `400 Bad Request`
The request body is invalid or missing required fields. 
Example Response:
```json
{
  "error": "Invalid request body."
}
```
#### status Code: `500 Internal Server Error`
An error occurred while creating the chat message.
Example Response:
```json
{
  "error": "Unable to create chat message."
}
```
#### The endpoint is defined in Go as follows:
```go
api.Post("/chat", controller.CreateChat(db))
```
#### Handler: `controller.CreateChat(db)`
This handler processes the incoming request to create a new chat message in the database. It validates the request body, stores the message, and returns the created message details or an appropriate error message.
______________________________________________________________________________

### **2. /chats/:group**
#### **URL:** `/chats/:group`
### METHOD: `GET`
The `/chats/:group` endpoint retrieves all chat messages for a specific group.


#### **Response Format**

| Field       | Type     | Description                                   |
|-------------|----------|-----------------------------------------------|
| `id`        | `int`    | The unique identifier for the chat message.   |
| `message`   | `string` | The content of the chat message.              |
| `sender`    | `string` | The identifier of the user who sent the message. |
| `group`     | `string` | The group identifier for the chat message.    |
| `createdAt` | `string` | The timestamp when the message was created.   |

#### **Example Request**

```bash
curl --location '/api/v1/chats/group1' \
--header 'Content-Type: application/json'
```
### **Success Response**
#### Status Code: `200 OK`
Response Body:
```json
[
  {
    "id": 1,
    "message": "Hello, team!",
    "sender": "User1",
    "group": "group1",
    "createdAt": "2025-01-18T10:00:00Z"
  },
  {
    "id": 2,
    "message": "Good morning!",
    "sender": "User2",
    "group": "group1",
    "createdAt": "2025-01-18T10:05:00Z"
  }
]
```
```bash
curl --location GET '' \
--header 'Content-Type: application/json' \
--data '[
  {
    "id": 1,
    "message": "Hello, team!",
    "sender": "User1",
    "group": "group1",
    "createdAt": "2025-01-18 T10:00:00Z"
  },
  {
    "id": 2,
    "message": "Good morning!",
    "sender": "User2",
    "group": "group1",
    "createdAt": "2025-01-18 T10:05:00Z"
  }
]
'
```
### **Error Responses**
#### Status Code: `404 Not Found`
If the group is not found or no chats exist for the specified group. 

Example Response:
```json
{
  "error": "No chats found for the specified group."
}
```
#### Status Code: `500 Internal Server Error`
If an internal error occurs while fetching the chats.

Example Response:
```json
{
  "error": "Unable to fetch chats."
}
```
#### The endpoint is defined in Go as follows:
```go
api.Get("/chats/:group", controller.GetChats(db))
```
#### Handler: `controller.GetChats(db)`
This handler fetches all chat messages for the specified group from the database and returns them as a JSON response. If no messages are found, an appropriate error message is returned.


________________________________________________________________________________

