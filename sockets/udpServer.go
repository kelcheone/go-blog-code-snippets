package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    // Create a UDP address at port 8000
    addr, err := net.ResolveUDPAddr("udp", ":8000")
    if err != nil {
        log.Fatal(err)
    }
    // Create a UDP listener
    ln, err := net.ListenUDP("udp", addr)
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close()
    fmt.Println("Listening on port 8000")
    buf := make([]byte, 1024)
    for {
        // Read data from the connection
        n, addr, err := ln.ReadFromUDP(buf)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Received %s from %s \n", string(buf[:n]), addr)
        // Write data to the connection
        _, err = ln.WriteToUDP([]byte(time.Now().String()), addr)
        if err != nil {
            log.Fatal(err)
        }
    }
}
