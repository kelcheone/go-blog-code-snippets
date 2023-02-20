package main

import (
    "context"
    "fmt"
    "time"
)

func someFunc(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Done")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go someFunc(ctx)
    time.Sleep(5 * time.Second)
    cancel()
    time.Sleep(1 * time.Second)
}
