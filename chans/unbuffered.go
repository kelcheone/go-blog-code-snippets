package main

import "fmt"

func main() {
    ch := make(chan int, 1)

    ch <- 200
    fmt.Println(<-ch)
}
