package main

import (
  "log"
  "os/exec"
)


func main() {

  cmd := exec.Command("firefox")
  if err := cmd.Start();err != nil{
     log.Fatal(err)
  }

 if err := cmd.Wait();err != nil{
   log.Fatal(err)
  }
}
