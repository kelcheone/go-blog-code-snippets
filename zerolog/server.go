package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }

    defer ln.Close()

    for {
        con, err := ln.Accept()
        if err != nil {
            panic(err)
        }

        go handleConnection(con)
    }
}

func handleConnection(con net.Conn) {
    defer con.Close()

    scanner := bufio.NewScanner(con)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
