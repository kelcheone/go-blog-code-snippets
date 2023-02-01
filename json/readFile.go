package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var student []Student

	err = json.NewDecoder(strings.NewReader(string(data))).Decode(&student)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student)
}
