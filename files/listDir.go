package main

import (
  "fmt"
  "io/ioutil"
)

func main() {
  files, err := ioutil.ReadDir("../../Go")
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, file := range files {
    fmt.Println(file.Name())
  }
}
