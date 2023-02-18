package main

import (
  "fmt"
  "sync"
)

func main() {
    var w sync.WaitGroup
    c := make(chan bool, 10)
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
    fmt.Println(counter)
}
