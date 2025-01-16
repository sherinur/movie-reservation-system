# MongoDB image
FROM mongo:latest

# Expose MongoDB port
EXPOSE 27017

# Set the working directory
WORKDIR /app

# Copy binary file
COPY user-service /app/user-service

# Start
CMD ["sh", "-c", "mongod --bind_ip_all & ./user-service"]
