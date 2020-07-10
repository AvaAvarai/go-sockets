package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {

	// Connect to the server.
	conn, _ := net.Dial(connType, connHost+":"+connPort)

	//
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read in input from stdin.
		text, _ := reader.ReadString('\n')

		// Send to socket connection.
		conn.Write([]byte(text))

		// Listen for a reply.
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// Print the server's reply.
		fmt.Print("Message from server: " + message)
	}
}
