package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "hello world")
	}()
	cmd.Stdout = os.Stdout
	cmd.Run()

}
