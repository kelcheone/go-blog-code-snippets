package main

import (
  "fmt"
  "os"
)

func main(){
  file, err := os.OpenFile("file-n.txt", os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("File opening error", err)
    return
  }
  defer file.Close()

  file.Write([]byte(" The earth can be flat but the sky is always blue."))
  fmt.Println("Data successfully appended to file")
}
