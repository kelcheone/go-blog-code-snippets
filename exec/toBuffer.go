package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
)


func main() {

  cmd := exec.Command("ls", "-la")
  var stdout, stderr bytes.Buffer
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  if err := cmd.Run(); err != nil {
    log.Fatal(err)
  }
  fmt.Println(stdout.String())
  fmt.Println(stderr.String())
}
