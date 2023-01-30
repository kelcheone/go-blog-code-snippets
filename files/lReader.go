package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  file, err := os.Open("file.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    fmt.Print(line)
  }
}
