package main

import (
    "fmt"
//    "time"
)

func main() {
    ch := make(chan int, 1)

    select {
    case v := <-ch:
        fmt.Println(v)
    default:
        fmt.Println("No value ready")
    }

    ch <- 200

    select {
    case v := <-ch:
        fmt.Println(v)
    default:
        fmt.Println("No value ready")
    }
}
