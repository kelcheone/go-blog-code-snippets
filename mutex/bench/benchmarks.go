package main

import (
//  "fmt"
  "sync"
)

func useMutex() {
  var wg sync.WaitGroup
  var mutex sync.Mutex
  var counter int64

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
      mutex.Lock()
      counter++
      mutex.Unlock()
      wg.Done()
    }()
  }
  wg.Wait()

 // fmt.Println(counter)
}

func useChannel() {
 var w sync.WaitGroup
    c := make(chan bool, 1)
    counter := 0
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go func() {
            c <- true
            counter++
            <- c
            w.Done()
        }()
    }
    w.Wait()
   // fmt.Println(counter)
}


func main() {
  useMutex()
  useChannel()
}
