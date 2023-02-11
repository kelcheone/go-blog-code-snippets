package main

import (
	"os"
	"os/exec"
	// "strings"
)

// pipe the output of one command to another
func main() {
	cmd1 := exec.Command("ls", "-l")
	cmd2 := exec.Command("grep", "anotherPiping.go")

	// cmd1Out, _ := cmd1.Output()
	// cmd2.Stdin = strings.NewReader(string(cmd1Out))
	// cmd2Out, _ := cmd2.Output()
	// fmt.Println(string(cmd2Out))

	// pipe the output of cmd1 to cmd2
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = os.Stdout
	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	cmd2.Wait()
}
