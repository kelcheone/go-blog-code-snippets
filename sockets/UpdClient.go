package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    // Create a UDP address at port 8000
    addr, err := net.ResolveUDPAddr("udp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    // Create a UDP connection
    conn, err := net.DialUDP("udp", nil, addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    buf := make([]byte, 1024)
    for {
        // Write data to the connection
        _, err = conn.Write([]byte("Hello UDP Server"))
        if err != nil {
            log.Fatal(err)
        }
        // Read data from the connection
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(buf[:n]))
        time.Sleep(time.Second)
    }
}
