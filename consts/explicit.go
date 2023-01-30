package main

import "fmt"

func main() {
  const (
    Pi float64 = 3.14
    Language string = "Go"
    Version int = 10e6
  )

    fmt.Printf("%T %T %T\n", Pi, Language, Version)
}
