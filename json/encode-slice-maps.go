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

    fmt.Println("Marshaling a slice of structs")

    students := []Student{
        {
            Name: "John",
            Age:  21,
        },
        {
            Name: "Jane",
            Age:  22,
        },
    }

    jsonSlice, err := json.Marshal(students)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(jsonSlice))

    fmt.Println("Marshaling a map of structs")

    studentsMap := map[string]Student{
        "John": {
            Name: "John",
            Age:  21,
        },
        "Jane": {
            Name: "Jane",
            Age:  22,
        },
    }

    jsonMap, err := json.Marshal(studentsMap)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(jsonMap))
}
