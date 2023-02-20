package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func publisher(out chan<- int) {
    for {
        time.Sleep(time.Second)
        out <- rand.Intn(100)
    }
}


func subscriber(in <-chan int, id int) {
    for {
        fmt.Printf("Subscriber %d received: %d\n", id, <-in)
    }
}

func pubSub() {
    out := make(chan int)
    for i := 1; i <= 3; i++ {
        go subscriber(out, i)
    }
    publisher(out)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        pubSub()
        wg.Done()
    }()
    wg.Wait()
}
