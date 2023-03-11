package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Language struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Compiler struct {
	Name     string   `yaml:"name"`
	Language Language `yaml:"language"`
}

type Runtime struct {
	Name     string   `yaml:"name"`
	Language Language `yaml:"language"`
}

type Config struct {
	Language Language `yaml:"language"`
	Compiler Compiler `yaml:"compiler"`
	Runtime  Runtime  `yaml:"runtime"`
}

func main() {
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	useNode()
}

func useNode() {
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var config map[string]interface{}
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v ", config["language"])
}

func useYamlNode() {
	var config map[string]yaml.Node

	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v ", config["language"])
}
