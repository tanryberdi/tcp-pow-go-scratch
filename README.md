# "Word of Wisdom" TCP-server with protection from DDOS based on Proof of Work

## 1. Description
This project is a solution for some interview question on Golang.

## 2. Getting started
### 2.1 Requirements
+ [Go 1.21+](https://go.dev/dl/) installed (to run tests, start server or client without Docker)
+ [OrbStack](https://orbstack.dev/download) (alternative for Docker) installed (to run docker-compose)

### 2.2 Start server and client by docker-compose:
```
make start
```

### 2.3 Start only server:
```
make start-server
```

### 2.4 Start only client:
```
make start-client
```

### 2.5 Check for lint:
```
make lint
```

## 3. Problem description
Design and implement “Word of Wisdom” tcp server.
TCP server should be protected from DDOS attacks with the [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work),
the challenge-response protocol should be used.  
The choice of the PoW algorithm should be explained.  
After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.  
Docker file should be provided both for the server and for the client that solves the PoW challenge.

## 4. Proof of Work
Idea of Proof of Work for DDOS protection is that client, which wants to get some resource from server,
should firstly solve some challenge from server.
This challenge should require more computational work on client side and verification of challenge's solution - much less on the server side.

### 4.1 Selection of an algorithm
To be honest i heard about Proof of Work before, but never touched it. After some investigation i found that
Hashcash is the most popular algorithm for Proof of Work. I implemented Hashcash for TCP-server.
For my solution, server can verify client's solution immediately after receiving it.