# JWT Auth Go Example

A simple, production-ready JWT (JSON Web Token) authentication system in Go. This project demonstrates how to secure RESTful APIs using JWTs, with endpoints for login, logout, protected resources, and token refresh. It is ideal for learning or as a starting point for your own Go authentication service.



## Features
- 🔒 User login with JWT token generation
- 🛡️ Protected route (`/home`) accessible only with a valid JWT
- ♻️ Token refresh endpoint
- 🚪 Logout functionality (token invalidation)
- 📦 Modular handler structure for easy extension

## Endpoints
| Method | Endpoint   | Description                |
|--------|------------|----------------------------|
| POST   | /login     | Authenticate and get token |
| POST   | /logout    | Invalidate token/logout    |
| GET    | /home      | Protected resource         |
| POST   | /refresh   | Refresh JWT token          |

### Example Requests
#### Login
```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"youruser","password":"yourpass"}' http://localhost:8080/login
```
**Response:**
```json
{
  "token": "<JWT_TOKEN>",
  "refresh_token": "<REFRESH_TOKEN>"
}
```

#### Access Protected Route
```bash
curl -H "Authorization: Bearer <JWT_TOKEN>" http://localhost:8080/home
```

#### Refresh Token
```bash
curl -X POST -H "Authorization: Bearer <REFRESH_TOKEN>" http://localhost:8080/refresh
```

---

## Project Structure
```
main.go           # Entry point, sets up HTTP routes
handlers/main.go  # Handler functions for each endpoint
```

## Getting Started

### Prerequisites
- Go 1.18 or newer

### Configuration
- Set your JWT secret and other configs in `handlers/main.go` (see comments in code).

### Installation & Run
1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd jwt-auth-go-main
   ```
2. Download dependencies:
   ```bash
   go mod tidy
   ```
3. Set up your environment variables in a `.env` file (see `.env.example`).
  
   ```bash
   export $(grep -v '^#' .env | xargs)
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
5. The server will start on `http://localhost:8080`

---

## Usage
- Use [Postman](https://www.postman.com/) or `curl` to interact with the endpoints.
- Always include the JWT as a Bearer token in the `Authorization` header for protected endpoints.

---

