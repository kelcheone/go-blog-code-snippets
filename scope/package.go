package main

import "fmt"

var (
	i = 100
)

func someFunc2() {
	i += 200
	fmt.Println(i)
}

func main() {
	someFunc2()
	fmt.Println(i)
	otherFunc()
}
