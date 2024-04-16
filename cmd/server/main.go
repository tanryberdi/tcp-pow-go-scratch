// server.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var quotes = []string{
	"Do not take life too seriously. " +
		"You will never get out of it alive",

	"The best way out is always through",

	"Knowledge speaks, but wisdom listens",

	"Always do what you are afraid to do",

	"Life is what happens to us while we are making other plans",
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

	// Handle incoming connections in a separate goroutine
	go func() {
		for {
			conn, _ := ln.Accept()
			go handleConnection(conn)
		}
	}()

	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay interrupt signals to the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine that listens on the signal channel and gracefully shuts down the server
	go func() {
		<-c
		ln.Close()
		os.Exit(0)
	}()

	// Keep the main function alive indefinitely
	select {}
}
