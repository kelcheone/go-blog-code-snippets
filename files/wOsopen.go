package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.OpenFile("file-n.txt", os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("File opening error", err)
    return
  }
  defer file.Close()

  file.Write([]byte("Here via os.OpenFile"))

  fmt.Println("Data successfully written to file")
}
