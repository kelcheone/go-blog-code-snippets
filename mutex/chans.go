package main

import (
  "fmt"
  "sync"
)

func main() {
    var w sync.WaitGroup
    ch := make(chan bool, 1)
    x := 0
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go func() {
            ch <- true
            x = x + 1
            <- ch
            w.Done()
        }()
    }
    w.Wait()
    fmt.Println("final value of x", x)
}
