package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd, err := exec.Command("ls", "-la").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(cmd))

}
