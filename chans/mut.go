package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    ch := make(chan bool, 1)
    var counter int

    for i := 0; i < 1000; i++{
        wg.Add(1)
        go func() {
            ch <- true
            counter++
            <-ch
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
