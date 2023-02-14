package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
)

func main() {
    // listen on all interfaces
    ln, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Listening on port 8000")
    // accept connection on port
    conn, err := ln.Accept()
    if err != nil {
        log.Fatal(err)
    }
    // run loop forever (or until ctrl-c)
    for {
        // will listen for message to process ending in newline (\n)
        message, err :=  bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        // output message received
        fmt.Print("Message Received:", string(message))
        // sample process for string received
        newmessage := strings.ToUpper(message)
        // send new string back to client
        conn.Write([]byte(newmessage + "\n"))
    }
}
