package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Student struct {
        Name string `json:"name"`
        Age  int    `json:"-"`
    }

    student := Student{
        Name: "John",
        Age:  21,
    }

    json, err := json.Marshal(student)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(json))
}
