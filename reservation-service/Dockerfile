FROM golang:1.23

WORKDIR /app

COPY reservation-service/ . 

# RUN go build -o movie-service cmd/main.go

EXPOSE 8090

CMD ["./reservation-service"]