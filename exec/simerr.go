package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("ls", "eses")
	cmd2 := exec.Command("grep", "main.go")
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = os.Stdout
	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	cmd2.Wait()
}
