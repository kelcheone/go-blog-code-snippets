package main

import (
    "encoding/json"
    "fmt"
    "strings"
)


func main() {
    type Student struct{
        Name string
        Age int
    }

    jsonStr := `{"Name":"John","Age":21}`

    decoder := json.NewDecoder(strings.NewReader(jsonStr))

    var student Student

    err := decoder.Decode(&student)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(student)
}
