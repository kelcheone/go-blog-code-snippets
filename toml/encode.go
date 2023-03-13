package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Postgres struct {
		Server   string `toml:"server"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
}

func main() {
	var config Config
	config.Postgres.Server = "localhost"
	config.Postgres.Port = 5432
	config.Postgres.Username = "postgres"
	config.Postgres.Password = "postgres"

	f, err := os.Create("postgre.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(config)
	if err != nil {
		panic(err)
	}
}
