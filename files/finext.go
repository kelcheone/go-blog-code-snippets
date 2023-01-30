package main

import (
  "fmt"
  "io/ioutil"
  "path"
)

func main() {
  files, err := ioutil.ReadDir("../files")
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, file := range files {
    if path.Ext(file.Name()) == ".txt" {
      fmt.Println(file.Name())
    }
  }
}
