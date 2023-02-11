package main

import (
   "log"
   "os"
  "os/exec"
)

func main() {

cmd1 := exec.Command("ls", "-la")
cmd2 := exec.Command("grep", "main.go")
cmd2.Stdin, _ = cmd1.StdoutPipe()
cmd2.Stdout = os.Stdout
cmd2.Stderr = os.Stderr
if err := cmd1.Start(); err != nil {
  log.Fatal(err)
}
if err := cmd2.Run(); err != nil {
  log.Fatal(err)
}
}
