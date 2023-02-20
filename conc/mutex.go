package main

import (
  "fmt"
  "sync"
)

func main() {
  var wg sync.WaitGroup
  var m sync.Mutex
  counter := 0
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      m.Lock()
      counter++
      m.Unlock()
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println(counter)
}
