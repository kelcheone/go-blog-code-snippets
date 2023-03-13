package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

// use toml.Tree to unmarshal
func main() {
	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	// get value
	fmt.Println(tree.Get("Database.Server"))
}
