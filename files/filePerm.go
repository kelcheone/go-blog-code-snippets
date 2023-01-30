package main

import (
  "fmt"
  "os"
)

func main() {
  fileInfo, err := os.Stat("file.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

 fileMode := fileInfo.Mode()
  filePerm := fileMode.Perm()
  fmt.Println(filePerm)
}
