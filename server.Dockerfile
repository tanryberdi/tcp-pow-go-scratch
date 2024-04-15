# Dockerfile for server
FROM golang:latest
#WORKDIR /app
#COPY . .
#RUN go build -o server server.go
#CMD ["./server"]

WORKDIR /build

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server