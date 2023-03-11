package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
)

type Blog struct {
    Contributors []string `yaml:"contributors"`
    Title        string   `yaml:"title"`
    Date         string   `yaml:"date"`
    Tags         []string `yaml:"tags"`
    Draft        bool     `yaml:"draft"`
}

func main() {
    var blog Blog

    yamlFile, err := os.ReadFile("blog.yaml")
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(yamlFile, &blog)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Contributors: %v \n", blog.Contributors)
    fmt.Printf("Title: %v \n", blog.Title)
    fmt.Printf("Date: %v \n", blog.Date)
    fmt.Printf("Tags: %v \n", blog.Tags)
    fmt.Printf("Draft: %v \n", blog.Draft)
}
