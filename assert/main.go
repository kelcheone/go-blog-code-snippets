package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	a := []byte("hello world")

	fmt.Println(cases.Title(language.English).Bytes(a))         // [72 101 108 108 111 32 87 111 114 108 100]
	fmt.Println(string(cases.Title(language.English).Bytes(a))) // Hello World
}
