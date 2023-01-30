package main

import "fmt"

func main() {
    source := []int{1, 2, 3, 4, 5}
    destination := append([]int{}, source...)

    fmt.Println(destination)
}
