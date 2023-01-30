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

    fmt.Println("Unmarshaling a slice of structs")

    jsonDataSlice := []byte(`[{"Name":"John","Age":21},{"Name":"Jane","Age":22}]`)
    var students []Student

    err := json.Unmarshal(jsonDataSlice, &students)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(students)

    fmt.Println("Unmarshaling a map of structs")

    jsonDataMap := []byte(`{"John":{"Name":"John","Age":21},"Jane":{"Name":"Jane","Age":22}}`)

    var studentsMap map[string]Student

    err = json.Unmarshal(jsonDataMap, &studentsMap)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(studentsMap)
}
