package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server.
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// Create a new zlib writer.
	w := zlib.NewWriter(connection)

	defer w.Close()

	// Continously write data to the server.
	for {
		// get user input and wait for enter
		fmt.Print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		// write to server
		w.Write([]byte(text))
		w.Flush()
	}

}
