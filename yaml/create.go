package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "os"
    "io"
)

type Person struct {
    Name   string   `yaml:"name"`
    Age    int      `yaml:"age"`
    Hobbies []string `yaml:"hobbies"`
}

func main() {
    Person := Person{
        Name:   "Kevin Kelche",
        Age:    21,
        Hobbies: []string{"Programming", "Reading", "Writing"},
    }

    yamlFile, err := yaml.Marshal(&Person)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(yamlFile))

    // Write to file
    f, err := os.Create("person.yaml")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = io.WriteString(f, string(yamlFile))
    if err != nil {
        panic(err)
    }
}
