
package main

import (
  "fmt"
  "os"
)

func main() {
  s, err := os.Stat("file.txt")
  fmt.Println(s)
  if os.IsNotExist(err) {
    fmt.Println("File does not exist")
  } else {
    fmt.Println("File exists")
  }
}
