package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  file, err := os.Create("file-n.txt")
  if err != nil {
    fmt.Println("File creation error", err)
    return
  }
  defer file.Close()

  writer := bufio.NewWriter(file)
  writer.Write([]byte("Here via bufio.Writer"))
  writer.Flush()

  fmt.Println("Data successfully written to file")
}
