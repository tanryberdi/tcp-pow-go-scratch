// server.go
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"tcp-pow-go-scratch/config"
)

var quotes = []string{
	"Do not take life too seriously. " +
		"You will never get out of it alive",

	"The best way out is always through",

	"Knowledge speaks, but wisdom listens",

	"Always do what you are afraid to do",

	"Life is what happens to us while we are making other plans",
}

//nolint:errcheck
func handleConnection(conn net.Conn, conf config.Config) {
	defer conn.Close()

	// Generate a random challenge
	challenge := fmt.Sprintf("%x", rand.Int63())

	// Send the challenge to the client
	conn.Write([]byte(challenge + "\n"))

	// Wait for the client's response
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Duration(conf.HashcashDuration) * time.Second))

	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Failed to receive a valid proof of work within the specified time limit")
		return
	}

	// Check if the client's response is a valid proof of work
	response := strings.TrimSpace(string(buf[:n]))
	hash := sha256.Sum256([]byte(challenge + response))
	if strings.HasPrefix(hex.EncodeToString(hash[:]), strings.Repeat("0", conf.HashcashZerosCount)) {
		// If the proof of work is valid, send a random quote
		conn.Write([]byte(quotes[rand.Intn(len(quotes))]))
	}
}

func main() {
	// Load the configuration
	conf, err := config.LoadConfig("config/config.json")
	if err != nil {
		fmt.Println("error loading config:", err)
		return
	}

	//nolint:staticcheck
	rand.Seed(time.Now().UnixNano())

	// Create a TCP server
	ln, _ := net.Listen("tcp", ":8080")

	// Handle incoming connections in a separate goroutine
	go func() {
		for {
			conn, _ := ln.Accept()
			go handleConnection(conn, conf)
		}
	}()

	// Create a channel to receive OS signals
	c := make(chan os.Signal, 1)

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
