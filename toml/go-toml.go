package main

import (
    "github.com/pelletier/go-toml"
    "io/ioutil"
)

type Config struct {
    Database struct {
        Server   string
        Port     int
        Username string
        Password string
    }
}

func main() {
    // Create a new Config struct with some data
    config := Config{
        Database: struct {
            Server   string
            Port     int
            Username string
            Password string
        }{
            Server:   "localhost",
            Port:     3306,
            Username: "user123",
            Password: "secret",
        },
    }

    // Encode the Config struct to TOML format
    data, err := toml.Marshal(config)
    if err != nil {
        panic(err)
    }

    // Write the encoded data to a file
    err = ioutil.WriteFile("config.toml", data, 0644)
    if err != nil {
        panic(err)
    }
}
