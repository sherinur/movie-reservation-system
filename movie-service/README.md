# Movie Service

## Overview

Movie Service is a backend system designed to manage cinemas, movies, screenings, and seat bookings. Built with Go and MongoDB, the service adheres to clean architecture principles and provides RESTful endpoints for seamless interaction.

## Prerequisites

- **Go**: Version 1.22.2 or higher
- **MongoDB**: Ensure you have a running MongoDB instance (local or hosted).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/movie-service.git
   cd movie-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables in a `.env` file or export them in your terminal:
   ```env
   PORT=:8080
   DB_URI=mongodb://localhost:27017
   DB_NAME=movieDB
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

3. Verify the service is running:
   ```bash
   curl http://localhost:8080/health
   ```

## Endpoints

### Movies

#### **POST** `/movie/add`
Add one or more movies to the database.
All fields must be filled in.

**Request Body:**
```json
[
  {
    "title": "string",
    "genre": "string",
    "description": "string",
    "posterImage": "string",
    "duration": 120,
    "language": "string",
    "releaseDate": "YYYY-MM-DD",
    "rating": "string",
    "pgrating": "string",
    "production": "string",
    "producer": "string",
    "status": "string"
  }
]
```

**Responses:**
- `200 OK`: Movie(s) added successfully.
- `400 Bad Request`: Invalid input.

#### **GET** `/movie/get`
Retrieve all movies from the database.

**Responses:**
- `200 OK`: List of movies.
- `500 Internal Server Error`: Server error.

#### **PUT** `/movie/update/{id}`
Update a movie by its ID.
All fields must be filled in.

**Request Body:**
```json
{
  "title": "string",
  "genre": "string",
  "description": "string",
  "posterImage": "string",
  "duration": 120,
  "language": "string",
  "releaseDate": "YYYY-MM-DD",
  "rating": "string",
  "pgrating": "string",
  "production": "string",
  "producer": "string",
  "status": "string"
}
```

**Responses:**
- `200 OK`: Movie updated successfully.
- `400 Bad Request`: Invalid input.

#### **DELETE** `/movie/delete/{id}`
Delete a movie by its ID.

**Responses:**
- `204 No Content`: Movie deleted successfully.
- `500 Internal Server Error`: Server error.

### Cinemas

#### **POST** `/cinema/add`
Add one or more cinemas to the database.
All fields must be filled in.

**Request Body:**
```json
[
  {
    "name": "string",
    "address": "string",
    "rating": 4.5,
    "hallList": [
      {
        "number": 1,
        "rowCount": 10,
        "columnCount": 20,
        "seats": [
          { "row": "A", "column": "1", "status": "available" }
        ],
        "screenings": [
          {
            "movieID": "string",
            "startTime": "2024-01-01T10:00:00Z",
            "endTime": "2024-01-01T12:00:00Z",
            "hallNumber": 1,
            "availableSeats": 200
          }
        ]
      }
    ]
  }
]
```

**Responses:**
- `200 OK`: Cinema(s) added successfully.
- `400 Bad Request`: Invalid input.

#### **GET** `/cinema/get`
Retrieve all cinemas from the database.

**Responses:**
- `200 OK`: List of cinemas.
- `500 Internal Server Error`: Server error.

#### **PUT** `/cinema/update/{id}`
Update a cinema by its ID.
All fields must be filled in.

**Request Body:**
```json
{
  "name": "string",
  "address": "string",
  "rating": 4.5,
  "hallList": [
    {
      "number": 1,
      "rowCount": 10,
      "columnCount": 20,
      "seats": [
        { "row": "A", "column": "1", "status": "available" }
      ],
      "screenings": [
        {
          "movieID": "string",
          "startTime": "2024-01-01T10:00:00Z",
          "endTime": "2024-01-01T12:00:00Z",
          "hallNumber": 1,
          "availableSeats": 200
        }
      ]
    }
  ]
}
```

**Responses:**
- `200 OK`: Cinema updated successfully.
- `400 Bad Request`: Invalid input.

#### **DELETE** `/cinema/delete/{id}`
Delete a cinema by its ID.

**Responses:**
- `204 No Content`: Cinema deleted successfully.
- `500 Internal Server Error`: Server error.

