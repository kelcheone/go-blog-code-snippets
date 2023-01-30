package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
	}
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
	}
	var p2 Person
	err = json.Unmarshal(decoded, &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

	b := true

	s := strconv.FormatBool(b)

	fmt.Println(s)      // true
	fmt.Printf("%T", s) // string
}
