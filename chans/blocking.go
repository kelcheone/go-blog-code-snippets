package main

import "fmt"

func someFunc(ch chan int) {
    ch <- 200
}

func main() {
    ch := make(chan int)

    go someFunc(ch)

    fmt.Println(<-ch)
}
