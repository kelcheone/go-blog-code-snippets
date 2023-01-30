package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Student struct {
        Name int
        Age  int
    }

    // with the Age field of type int, this will throw an error
    jsonStr := `{"Name":"John","Age":"21"}`
    var student Student

    err := json.Unmarshal([]byte(jsonStr), &student)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(student)
}

