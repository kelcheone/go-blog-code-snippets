package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 200
    }()
    fmt.Println(<-ch)
}
