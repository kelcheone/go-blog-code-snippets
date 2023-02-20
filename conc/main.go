package main

import (
  "fmt"
  "time"
)

func main() {
  go func() {
    fmt.Println("Hello from goroutine")
  }()
  time.Sleep(1 * time.Second)
  fmt.Println("Hello from main")
}
