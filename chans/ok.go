package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 200
        close(ch)
    }()

    for {
        v, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(v)
    }
}
