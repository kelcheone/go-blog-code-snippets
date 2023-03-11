package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
)

func main() {
    var blog map[string]interface{}

    yamlFile, err := os.ReadFile("blog.yaml")
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(yamlFile, &blog)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Contributors: %v \n %v \n", blog["contributors"].([]interface{})[0], blog["contributors"].([]interface{})[1])
    fmt.Printf("Title: %v \n", blog["title"])
    fmt.Printf("Date: %v \n", blog["date"])
    fmt.Printf("Tags: %v \n %v \n", blog["tags"].([]interface{})[0], blog["tags"].([]interface{})[1])
    fmt.Printf("Draft: %v \n", blog["draft"])
}
