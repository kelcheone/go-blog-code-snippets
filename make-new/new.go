package main

import "fmt"

func main() {
    // create a pointer to an integer
    i := new(int)
    fmt.Println(i)

    // create a pointer to a string
    s := new(string)
    fmt.Println(s)

    // create a pointer to a slice
    sl := new([]int)
    fmt.Println(sl)

    // create a pointer to a map
    m := new(map[int]int)
    fmt.Println(m)

    // create a pointer to a channel
    c := new(chan int)
    fmt.Println(c)
}
