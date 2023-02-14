package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    // connect to this socket
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    for {
        // read in input from stdin
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        // send to socket
        fmt.Fprintf(conn, text + "\n")
        // listen for reply
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Message from server: "+message)
    }
}
