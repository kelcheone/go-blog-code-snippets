package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Student struct {
        Name string
        Age  int
    }

    student := Student{
        Name: "John",
        Age:  21,
    }

    jsonData, err := json.MarshalIndent(student, "", "    ")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(jsonData))
}
