package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database struct {
		Server   string `toml:"server"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}

func main() {
	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Database.Server)
}
