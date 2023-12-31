package main

import (
    "bufio" // buffered IO
    "fmt"   // formatted IO
    "net"   // networking IO
    // "os"    // operating system IO
)

/*
	This is a simple TCP server that listens on port 8080 and returns
	1. Listen on a TCP Port: Your server should listen on a specific port for incoming TCP connections.
	2. Accept Connections: Once a connection is made, your server should accept it and possibly handle it in a separate goroutine for concurrent processing.
	3. Read Data from the Connection: For each connection, your server will need to read incoming data.
	4. Process Requests: After reading the data, your server should process it according to your application's logic.
	5. Send Responses: Optionally, your server can send responses back to the client.
	6. Handle Errors and Close Connections: Properly handle any errors and ensure connections are closed gracefully.
*/

func handleRequest(conn net.Conn) {
	// close connection on exit

	// create a buffer stream and read from it, only until \n. otherwise will need a multi-line parser
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')


	// handle errors
	if err != nil {
        fmt.Println("Error reading:", err.Error())
        conn.Close()
        return
    }

	// Process the received data
	fmt.Printf("Received: %s", string(buffer))

	// Send a response back to the client
	response := []byte("Message received.\n")
	conn.Write(response)

	// Close the connection
	conn.Close()
}

func listen() {
	// Listen for TCP connections on localhost 8080
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// By using defer listener.Close(), you're ensuring that the listener is closed
	// and its resources are freed up when the enclosing function (main in this case) returns.
	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:8080...")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}

		// Handle the connection in a new goroutine, without 'go' it will be run synchronously
		// Starting a new go routine allows the server to continue accepting new connections concurrently.
		go handleRequest(conn)
	}
}

func main() {
	listen()
}