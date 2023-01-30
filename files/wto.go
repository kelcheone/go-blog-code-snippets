package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  data := []byte("Hello World!")
  err := ioutil.WriteFile("file-n.txt", data, 0644)
  if err != nil {
    fmt.Println("File writing error", err)
    return
  }
  fmt.Println("Data successfully written to file")
}
