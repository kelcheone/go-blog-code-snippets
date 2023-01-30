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

    jsonStr := ``
    var student Student

    err := json.Unmarshal([]byte(jsonStr), &student)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(student)
}
