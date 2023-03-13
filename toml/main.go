package main

import (
    "fmt"
    "log"

    "github.com/BurntSushi/toml"
)

type Server struct {
    Host string `toml:"host"`
    Port int `toml:"port"`
    Running bool `toml:"running"`
}

type Development struct {
    Host string `toml:"host"`
    Port int `toml:"port"`
    Running bool `toml:"running"`
}

type Config struct {
    Server Server `toml:"server"`
    Development Development `toml:"development"`
    Data []string `toml:"data"`
}

func main() {
    var config Config

    _, err := toml.DecodeFile("example.toml", &config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(config)
}
