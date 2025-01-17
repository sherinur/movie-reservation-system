# MongoDB image
FROM golang:1.22.2

# Expose MongoDB port
EXPOSE 8080

# Set the working directory
WORKDIR /app

# Copy binary file
COPY user-service /app/user-service
COPY .env /app/.env

# Start
ENTRYPOINT [ "./user-service" ]
