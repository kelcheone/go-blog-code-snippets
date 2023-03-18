package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	r, err := zlib.NewReader(connection)
	// print the buffer
	fmt.Println(r)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, r)
	fmt.Println()
}
