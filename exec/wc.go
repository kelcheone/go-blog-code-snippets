package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	wcCmd := exec.Command("wc")

	// Open a pipe to the process's standard input
	stdin, err := wcCmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error opening stdin pipe: %v\n", err)
		return
	}

	// Open a pipe to the process's standard output
	stdout, err := wcCmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error opening stdout pipe: %v\n", err)
		return
	}

	// Start the wc command
	if err := wcCmd.Start(); err != nil {
		fmt.Printf("Error starting wc command: %v\n", err)
		return
	}

	// Write some text to the process's standard input
	text := "Hello, World!\n"
	if _, err := stdin.Write([]byte(text)); err != nil {
		fmt.Printf("Error writing to stdin: %v\n", err)
		return
	}

	// Close the pipe to signal end of input
	if err := stdin.Close(); err != nil {
		fmt.Printf("Error closing stdin: %v\n", err)
		return
	}

	// Read the process's standard output
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from stdout: %v\n", err)
		return
	}

	// Wait for the process to finish
	if err := wcCmd.Wait(); err != nil {
		fmt.Printf("Error waiting for wc command: %v\n", err)
		return
	}
}
