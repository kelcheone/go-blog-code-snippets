package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Student struct{
        Name string
        Age  int
    }

    jsonData := []byte(`{"Name":"John","Age":21}`)

    var student Student

    err := json.Unmarshal(jsonData, &student)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(student)
}
