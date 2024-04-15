# Dockerfile for client
FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o client client.go
CMD ["./client"]