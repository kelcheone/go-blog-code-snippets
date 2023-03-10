package main
import "fmt"

func main() {
    fmt.Println("Start")
    callback(func() {
        fmt.Println("Callback")
    })
    fmt.Println("End")
}

func callback(f func()) {
    f()
}
