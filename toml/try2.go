package main

import (
	"fmt"

	"io"
	"os"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	Database struct {
		Password string
		Port     int
		Server   string
		Username string
	}
}

func main() {
	// type Postgres struct {
	// 	User     string
	// 	Password string
	// }
	// type Config struct {
	// 	Postgres Postgres
	// }

	// doc := []byte(`
	// [postgres]
	// user = "pelletier"
	// password = "mypassword"`)

	// config := Config{}
	// toml.Unmarshal(doc, &config)
	// fmt.Println("user=", config.Postgres.User)
	file, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config
	// into bytes
	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// into struct
	err = toml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Database.Server)

}
