// main_test.go
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"

	"tcp-pow-go-scratch/config"
)

var quotes2 = []string{""}

func TestHandleConnection(t *testing.T) {
	// Create a pair of connected network connections
	serverConn, clientConn := net.Pipe()

	// Create a configuration object
	conf := config.Config{
		HashcashDuration:      5,
		HashcashZerosCount:    4,
		HashcashMaxIterations: 1000,
	}

	// Run the handleConnection function in a separate goroutine
	go handleConnection(serverConn, conf)

	// Read the challenge from the server
	reader := bufio.NewReader(clientConn)
	challenge, _ := reader.ReadString('\n')
	challenge = strings.TrimSpace(challenge)

	// Generate a valid proof of work for the challenge
	var proofOfWork string
	start := time.Now()
	for i := 0; i < conf.HashcashMaxIterations; i++ {
		proofOfWork = fmt.Sprintf("%x", rand.Int63())
		hash := sha256.Sum256([]byte(challenge + proofOfWork))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), strings.Repeat("0", conf.HashcashZerosCount)) {
			break
		}
		if time.Since(start) > time.Duration(conf.HashcashDuration)*time.Second {
			t.Fatal("Failed to find a valid proof of work within the specified time limit")
		}
	}

	// Write the proof of work to the server
	//nolint:errcheck
	clientConn.Write([]byte(proofOfWork + "\n"))

	// Read the quote from the server
	quote, _ := reader.ReadString('\n')

	// Check if the quote is one of the expected quotes
	found := false
	for _, expectedQuote := range quotes2 {
		if strings.TrimSpace(quote) == expectedQuote {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Received unexpected quote: %s", quote)
	}
}
