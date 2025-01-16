# User Service

## Overview

User Service is a backend service that provides user authentication and management functionality, including login, registration, and profile retrieval. The service is built using Go and MongoDB, adhering to clean architecture principles.

## Prerequisites

- **Go**: Version 1.22.2 or higher
- **MongoDB**: Ensure you have a running MongoDB instance. You can use a local instance or connect to a hosted MongoDB Atlas cluster.
- **Docker** (optional): For containerized deployment

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/user-service.git
   cd user-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables in a `.env` file or export them in your terminal:
   ```env
   PORT=:8080
   DB_URI=mongodb://localhost:27017
   DB_NAME=userdb
   SECRET_KEY=your_secret_key
   ```

## Starting the Service

1. Run MongoDB:
   - If you have Docker installed, you can start MongoDB with:
     ```bash
     docker run -d -p 27017:27017 --name mongo mongo
     ```
   - Alternatively, start your local or hosted MongoDB instance.

2. Start the service:
   ```bash
   go run cmd/main.go
   ```

   The service will start on the port specified in the `PORT` environment variable (default is `:8080`).

3. Check the health endpoint to verify the service is running:
   ```bash
   curl http://localhost:8080/health
   ```

## Endpoints

### **POST** `/login`
Authenticate a user and return a JWT token.

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Responses:**
- `200 OK`: Returns the JWT token.
- `400 Bad Request`: Invalid request body.
- `401 Unauthorized`: Invalid credentials.

### **POST** `/register`
Register a new user.

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Responses:**
- `201 Created`: Returns the ID of the newly created user.
- `400 Bad Request`: Invalid request data.
- `409 Conflict`: User already exists.

### **GET** `/profile`
Retrieve user profiles (JWT required).

**Headers:**
```http
Authorization: Bearer <JWT>
```

**Responses:**
- `200 OK`: Returns a list of users.
- `401 Unauthorized`: Missing or invalid JWT.
- `500 Internal Server Error`: An error occurred on the server.

### **GET** `/health`
Check the health status of the service.

**Responses:**
- `200 OK`: Service is running.

## MongoDB Setup

The service requires a MongoDB database. By default, the database name is `userdb`. You can change this in the `.env` file. The `users` collection is used to store user data.

Example MongoDB user document:
```json
{
  "_id": "ObjectId",
  "username": "string",
  "passwordHash": "string"
}
```

## Development Notes

- **Project Structure:**
  ```
  user-service
  ├── cmd
  │   └── main.go          # Entry point of the application
  ├── internal
  │   ├── dal              # Data access layer
  │   ├── db               # Database connection utilities
  │   ├── handler          # HTTP handlers for the API
  │   ├── models           # Data models
  │   ├── service          # Business logic layer
  │   └── utils            # Utility functions (e.g., JWT, validation)
  ├── pkg
  │   └── logger           # Logging utilities
  ├── go.mod               # Go module definition
  └── README.md            # Project documentation
  ```

- **Error Handling:** All errors are handled gracefully and appropriate HTTP responses are returned.

- **Testing:** Use `go test` to run unit tests.

## Future Improvements

- Implement graceful shutdown.
- Add rate limiting to protect against abuse.
- Enhance logging for better observability.
- Add unit and integration tests for comprehensive coverage.

## Author
This project is written by Nurislam Sheri [(sherinur)](https://github.com/sherinur).
