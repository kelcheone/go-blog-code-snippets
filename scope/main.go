package main

import "fmt"

func someFunc(){
    i := 20

    fmt.Printf("We have %d initialized.", i)
}

func wasmain() {
    someFunc()
    fmt.Printf("Trying to print %v .", i)
}
