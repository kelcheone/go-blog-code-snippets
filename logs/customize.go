package main

import (
    "log"
    "os"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.SetPrefix("WARNING: ")
    log.SetOutput(os.Stdout)
    log.Println("Logger created and writing to os.Stdout")
}
