package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "3333"
	connType = "tcp"
)

func main() {
	listen, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listen.Close()

	fmt.Println("Listening on " + connHost + ": " + connPort)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024) // Make a buffer to hold incoming data.

	_, err := conn.Read(buf) // Read the incoming connection into the buffer.
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("Message received.\n"))
	conn.Close()
}
