# Reservation Service

## Overview

Reservation Service is a golang service which main functionality is to implement tickets booking reservation for each movie session. 

## Installation
Clone the repository:
   ```bash
   git clone https://github.com/your-username/reservation-service.git
   cd reservation-service
   ```

## Getting Started

1. Run MongoDB:
   Start your local MongoDB instance on specified port (27017).

2. Start the service:
   ```bash
   go run cmd/main.go
   ```

## Endpoints

### **POST** `/booking/add`
Adds a new reservation (on real situation: after arrange payment)

**Request Body:**
```json
{
  "movie_title": "Name of the movie",
  "email": "YourEmail@domain.com",
  "tickets": [
    {
      "seat": "SeatNumber",
      "price": 9999,
      "type": "adult"
    }
  ]
}
``` 


**Responses:**
- `201 Created`: Return the new created reservation.
- `400 Bad Request`: Invalid request body.

### **POST** `/booking/delete/id`
Deletes a reservation by its id

**Responses:**
- `200 OK`: Returns the ID of the newly created user.
- `400 Bad Request`: Invalid request data.

## MongoDB Setup

```json
{
  "_id": "ObjectId",
  "movie_title": "string",
  "email": "string",
  "status": "string",
  "boughttime": "string",
  "tickets": [
    {
      "seat": "string",
      "price": "double",
      "type": "string"
    }
  ],
  "qrcode": "string"
}
```

## Architecture

- **Project Structure:**
  ```
  reservation-service
  │  
  ├── cmd
  │   └── main.go         
  ├── internal
  │   ├── dal              
  │   ├── db               
  │   ├── handler          
  │   ├── models           
  │   └── service          
  ├── go.mod               
  └── README.md          
  ```

## Author
This project is written by Bolat Danial.
