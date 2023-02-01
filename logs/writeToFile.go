package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file:", err)
    }
    defer file.Close()

    logger := log.New(file, "INFO: ", log.LstdFlags)
    logger.Println("Logger created  and writing to log.txt")
}
