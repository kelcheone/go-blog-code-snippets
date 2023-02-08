package main

import (
    "fmt"
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
            // Output: some error
        }
    }()

    panic("some error")
}
