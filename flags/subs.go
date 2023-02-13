package main

import (
    "flag"
    "fmt"
)

func main() {
    name := flag.String("name", "Kevin", "The name of the user")
    age := flag.Int("age", 26, "The age of the user")

    flag.Parse()

    fmt.Println("Name:", *name)
    fmt.Println("Age:", *age)

    subcommand := flag.NewFlagSet("subcommand", flag.ExitOnError)
    subcommandName := subcommand.String("name", "Kevin", "The name of the user")
    subcommandAge := subcommand.Int("age", 26, "The age of the user")

    subcommand.Parse(flag.Args()[1:])
    fmt.Println("Subcommand Name:", *subcommandName)
    fmt.Println("Subcommand Age:", *subcommandAge)

}
