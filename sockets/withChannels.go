package main

import (
    "fmt"
    "log"
    "net"
    "strings"
)

func handleConnection(conn net.Conn, c chan string) {
    // Close the connection when the function returns
    defer conn.Close()
    // Create a buffer to store the received data
    buf := make([]byte, 1024)
    for {
        // Read the incoming connection into the buffer.
        n, err := conn.Read(buf)
        if err != nil {
            log.Println(err)
            return
        }
        // Print the received data
        fmt.Println(string(buf[:n]))
        // Convert the data to uppercase
        data := strings.ToUpper(string(buf[:n]))
        // Write the response back to the connection
        _, err = conn.Write([]byte(data))
        if err != nil {
            log.Println(err)
            return
        }
        // Send a message to the channel
        c <- "Message sent"
    }
}


func main() {
    // Create a TCP address at port 8000
    addr, err := net.ResolveTCPAddr("tcp", ":8000")
    if err != nil {
        log.Fatal(err)
    }
    // Create a TCP listener
    ln, err := net.ListenTCP("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()
    fmt.Println("Listening on port 8000")
    // Create a channel
    c := make(chan string)
    for {
        // Accept a connection
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal(err)
        }
        // Handle the connection in a new goroutine
        go handleConnection(conn, c)
        // Print the message received from the channel
        fmt.Println(<-c)
    }
}
