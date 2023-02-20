package main

import (
   "fmt"
   "sync"
)


func main(){
   var wg sync.WaitGroup
  ch := make(chan int, 20)

  wg.Add(1)
  go func(){
    for v:= range ch{
      fmt.Println(v)
    }
    wg.Done()
  }()
  wg.Wait()
}
