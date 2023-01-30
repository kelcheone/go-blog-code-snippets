package main

import "fmt"
type Mystruct struct {
    Name string
    Age  int
}

func main()  {
    myStruct := Mystruct{
        Name: "John",
        Age:  20,
    }

    fmt.Printf("%v", myStruct)
    fmt.Printf("%#v", myStruct)
}
