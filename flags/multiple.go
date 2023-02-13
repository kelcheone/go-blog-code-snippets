package main

import (
    "flag"
    "fmt"
    "strings"
)

type names []string

func (n *names) String() string {
    return fmt.Sprintf("%s", *n)
}

func (n *names) Set(value string) error {
    if len(value) < 3 {
        return fmt.Errorf("name is too short")
    }

    parts := strings.Split(value, " ")
    if len(parts) != 2 {
        return fmt.Errorf("name must be in the format 'first last'")
    }

    *n = append(*n, value)

    return nil
}

func main() {
    var n names
    flag.Var(&n, "name", "The name of the user")

    flag.Parse()

    fmt.Println("Names:", n)
}
