package main

import "fmt"

func main() {
  cancel := defer fmt.Println("world")

  fmt.Println("hello")

  cancel()
}
