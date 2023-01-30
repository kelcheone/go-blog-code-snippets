package main

import "fmt"

func main() {
    x:= make(map[string]int)

    // Adding a key-value pair
    x["foo"] = 1
    x["bar"] = 2

    // Deleting a key-value pair
    delete(x, "foo")

    fmt.Println(x)
}
