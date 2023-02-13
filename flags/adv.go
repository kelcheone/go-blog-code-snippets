package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    flag.CommandLine.SetOutput(os.Stdout)
    flag.CommandLine.SetUsage(func() {
        fmt.Println("Usage: go run main.go [flags]")
        flag.PrintDefaults()
    })

    var iter int
    flag.IntVar(&iter, "iter", 5, "The number of iterations")

    flag.Parse()

    for i := 0; i < iter; i++ {
        fmt.Println("Just chillin")
    }
}
