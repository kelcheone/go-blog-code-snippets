package main

import (
    "fmt"
    "time"
)

func main(){
    go func(){
        fmt.Println("Sleeping for 5 seconds...")
        time.Sleep(5 * time.Second)
        fmt.Println("Done sleeping.")
    }()
    fmt.Println("Waiting for goroutine to finish...")
    time.Sleep(10 * time.Second)
    fmt.Println("Done waiting.")
}
