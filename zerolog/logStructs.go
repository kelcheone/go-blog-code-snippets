package main

import (
    "os"

    "github.com/rs/zerolog"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{
        Name: "John Doe",
        Age:  30,
    }

// implemet a zerolog.LogObjectMarshaler interface

    marshledPerson := zerolog.ObjectMarshalerFunc(func(e *zerolog.Event) {
        e.Str("name", person.Name)
        e.Int("age", person.Age)
    })

    logger := zerolog.New(os.Stdout)
    logger.Info().Object("person", marshledPerson).Msg("")
}
