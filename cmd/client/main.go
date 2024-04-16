// client.go
package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

func main() {
	// Connect to the server
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	// Read the server's challenge
	challenge, _ := bufio.NewReader(conn).ReadString('\n')
	challenge = strings.TrimSpace(challenge)

	// Find a number that, when hashed with the challenge, results in a hash with 4 leading zeros
	var response string
	for {
		response = fmt.Sprintf("%x", rand.Int63())
		hash := sha256.Sum256([]byte(challenge + response))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), "0000") {
			break
		}
	}

	// Send the response to the server
	conn.Write([]byte(response + "\n"))

	// Delete the challenge after using it
	challenge = ""

	// Read and print the server's quote
	quote, _ := bufio.NewReader(conn).ReadString('\n')
	log.Println(quote)
}
