package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  data, err := ioutil.ReadFile("file.txt")
  if err != nil {
    fmt.Println("File reading error", err)
    return
  }

  fmt.Println("Contents of file:", string(data))
}
