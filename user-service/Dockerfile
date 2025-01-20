FROM golang:1.23

WORKDIR /app

COPY user-service/ . 

# RUN go build -o user-service cmd/main.go

EXPOSE 8080

CMD ["./user-service"]