package main

import "fmt"


func main() {

    var slice []int
    ch := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)
    }()

    for v := range ch {
        slice = append(slice, v)
    }

    ch2 := make(chan int)

    go func() {
        for v := range ch {
            ch2 <- v * 2
        }
        close(ch2)
    }()

    for v := range ch2 {
        slice = append(slice, v)
    }

    fmt.Println(slice)
}
