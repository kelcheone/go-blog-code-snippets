package main

import "fmt"

type Language string

const (
  Go Language = "Go"
  Python Language = "Python"
  JavaScript Language = "JavaScript"
)

func main() {
  fmt.Println(Go, Python, JavaScript)
}
