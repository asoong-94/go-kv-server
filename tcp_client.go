package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)


func main() {
    // Connect to the server, creates a network connection
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()

    // Send a message to the server
    fmt.Fprintf(conn, "Hello server!\n")

    // Read the server's response
    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading:", err.Error())
        return
    }
    fmt.Print("Message from server: ", response)
}