// server.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

var quotes = []string{
	"Quote 1",
	"Quote 2",
	"Quote 3",
	// Add more quotes as needed
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Generate a random challenge
	challenge := fmt.Sprintf("%x", rand.Int63())

	// Send the challenge to the client
	conn.Write([]byte(challenge + "\n"))

	// Wait for the client's response
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	// Check if the client's response is a valid proof of work
	response := strings.TrimSpace(string(buf[:n]))
	hash := sha256.Sum256([]byte(challenge + response))
	if strings.HasPrefix(hex.EncodeToString(hash[:]), "0000") {
		// If the proof of work is valid, send a random quote
		conn.Write([]byte(quotes[rand.Intn(len(quotes))]))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a TCP server
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()

	// Handle incoming connections
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
