package main

import (
    "fmt"
    "os"
)

func main() {
    os.Setenv("SomeVar", "SomeValue")

    someVar := os.Getenv("SomeVar")
    fmt.Println(someVar)
}
