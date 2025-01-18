# API Documentation

## URL: `/api/v1/login`
### METHOD: **POST**


The `/login` endpoint allows users to authenticate by providing their credentials (usually username and password). Upon successful authentication, a token or session will be returned.

##### Request Body:
The request body must be in `JSON` format and contain the following fields:

| Field       | Type     | Description                                |
|-------------|----------|--------------------------------------------|
| `email`  | `string` | The email of the user.         |
| `password`  | `string` | The password of the user.                  |

**Example Request:**
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

