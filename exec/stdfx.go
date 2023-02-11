package main

import (
  "log"
  "fmt"
  "os/exec"
)

func main() {
   out, err := exec.Command("firefox").Output()
  if err != nil{
    log.Fatal(err)
  }

 fmt.Println(string(out))
}

