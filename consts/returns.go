package main

import "fmt"

func returnFunc() int {
  defer fmt.Println("world")
  return 1
}

func main() {
  fmt.Println(returnFunc())
}
