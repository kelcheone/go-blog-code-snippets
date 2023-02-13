package main

import (
    "flag"
    "fmt"
    "strings"
)

type name struct {
    firstName string
    lastName  string
}

func (n *name) String() string {
    return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

func (n *name) Set(value string) error {
    if len(value) < 3 {
        return fmt.Errorf("name is too short")
    }

    parts := strings.Split(value, " ")
    if len(parts) != 2 {
        return fmt.Errorf("name must be in the format 'first last'")
    }
    n.firstName = parts[0]
    n.lastName = parts[len(parts)-1]

    return nil
}

func main() {
    var n name
    flag.Var(&n, "name", "The name of the user")

    flag.Parse()

    fmt.Println("Name:", n)
}
