package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sleep", "5")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
