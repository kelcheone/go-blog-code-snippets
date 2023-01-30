
package main

import (
  "fmt"
  "os"
)

func main() {
  fileInfo, err := os.Stat("empty.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  if fileInfo.Size() == 0 {
    fmt.Println("File is empty")
  } else {
    fmt.Println("File is not empty")
  }
}
