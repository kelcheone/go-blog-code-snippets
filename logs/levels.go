package main

import (
    "log"
)

func main() {
    log.Panicln("Panic message")
    log.Fatalln("Fatal message")
    log.Println("Info message")
    log.Println("Warning message")
}
