package main

import (
  "crypto/rand"
  "fmt"
  "math/big"
)

func main() {
  n, err := rand.Int(rand.Reader, big.NewInt(100))
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(n)
}
