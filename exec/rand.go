package main

import (
	"fmt"
	"io"
	"log"

	//  "os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-la")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "Some old Falcon")
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
