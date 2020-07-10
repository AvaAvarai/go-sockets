// Package main is the entry-point for the go-sockets project.
// go-sockets available under the GNU GENERAL PUBLIC LICENSE.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	// Start the server and listen for incoming connections.
	fmt.Println("Starting", connType, "server on port", connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		// Handle connections concurrently in a new goroutine.
		go handleConnection(c)
	}
}

// handleConnection handles the logic handling for a single connection request.
func handleConnection(conn net.Conn) {
	// Read in until a new-line character.
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close down left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	// Concatenate the response message.
	response := string(bufferBytes) + " from " + conn.RemoteAddr().String() + "\n"

	// Print the response message.
	log.Println(response)

	// Send the response message to the client.
	conn.Write([]byte(response))

	// Restart the process if the client stays connected.
	handleConnection(conn)
}
