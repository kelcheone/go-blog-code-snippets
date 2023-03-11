package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var config map[string]interface{}

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Language: %v \n", config["language"])
}
