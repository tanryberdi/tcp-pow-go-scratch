# Dockerfile for client
FROM golang:latest
#WORKDIR /app
#COPY . .
#RUN go build -o client client.go
#CMD ["./client"]

WORKDIR /build

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/client