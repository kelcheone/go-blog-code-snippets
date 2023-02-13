package main

import (
    "flag"
    "fmt"
)

func main() {
    name := flag.String("name", "Kevin", "The name of the person to greet.")
    flag.Parse()
    fmt.Printf("Hello %s!\n", *name)
}
