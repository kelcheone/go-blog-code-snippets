package main

import (
    "log"
    "os"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds)
    logger.Println("Logger created  and writing to os.Stdout")
}
