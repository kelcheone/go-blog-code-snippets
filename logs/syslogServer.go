package main

import (
    "log"
    "log/syslog"
)

// Syslog server in port 514 and make the server listen for incoming connections
func main() {
    server, err := syslog.New(syslog.LOG_INFO, "myapp")
    if err != nil {
        log.Fatalln("Failed to create syslog server:", err)
    }
    defer server.Close()

    server.Info("Logger created  and writing to syslog server")
}
