package main

import (
    "log"
    "os"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", 0)
    logger.Println("Logger created  and writing to os.Stdout")
}

