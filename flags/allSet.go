package main

import (
    "flag"
    "fmt"
    "time"
)

func main() {
    name := flag.String("name", "Kevin", "The name of the user")
    age := flag.Int("age", 26, "The age of the user")
    married := flag.Bool("married", false, "Whether the user is married")
    duration := flag.Duration("duration", 5*time.Second, "The duration of the user")

    flag.Parse()

    fmt.Println("Name:", *name)
    fmt.Println("Age:", *age)
    fmt.Println("Married:", *married)
    fmt.Println("Duration:", *duration)
}
