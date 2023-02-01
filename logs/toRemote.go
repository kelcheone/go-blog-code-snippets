package main

import (
    "log"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatalln("Failed to connect to remote server:", err)
    }
    defer conn.Close()

    logger := log.New(conn, "INFO: ", log.LstdFlags)
    logger.Println("Logger created  and writing to remote server")

}
