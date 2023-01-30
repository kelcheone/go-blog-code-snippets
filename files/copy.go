package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  source, err := os.Open("file.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  destination, err := os.Create("newfile.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  nBytes, err := io.Copy(destination, source)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(nBytes, "bytes copied")
}
