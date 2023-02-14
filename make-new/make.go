package main

import "fmt"

func main() {
    // make a slice of 5 integers
    slice := make([]int, 5)
    fmt.Println(slice)

    // make a map of 5 integers
    m := make(map[int]int, 5)
    fmt.Println(m)

    // make a channel of 5 integers
    c := make(chan int, 5)

    for i := 0; i < 5; i++ {
        c <- i
    }

    close(c)

    for v := range c {
        fmt.Println(v)
    }

}
