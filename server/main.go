// Package main is the entry-point for the go-sockets server sub-project.
// The go-sockets project is available under the GPL-3.0 License in LICENSE.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Application constants, defining host, port, and protocol.
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	// Start the server and listen for incoming connections.
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
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

		// Print client connection address.
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go handleConnection(c)
	}
}

// handleConnection handles logic for a single connection request.
func handleConnection(conn net.Conn) {
	// Buffer client input until a newline.
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	// Print response message, stripping newline character.
	log.Println("Client message:", string(buffer[:len(buffer)-1]))

	// Send response message to the client.
	conn.Write(buffer)

	// Restart the process.
	handleConnection(conn)
}
