package main

import (
    "log"
    "log/syslog"
)

func main() {
    logger, err := syslog.Dial("tcp", "localhost:514", syslog.LOG_INFO, "myapp")
    if err != nil {
        log.Fatalln("Failed to connect to syslog server:", err)
    }
    defer logger.Close()

    logger.Info("Logger created  and writing to syslog server")
}
