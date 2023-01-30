package main

import (
  "fmt"
  "path"
)

func main() {
  fileExtension := path.Ext("file.txt")
  fmt.Println(fileExtension)
}
