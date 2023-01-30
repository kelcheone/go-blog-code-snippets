package main

import "fmt"

func funcToBenchmark() {
	var a int
	for i := 0; i < 1000000; i++ {
		a += i
	}
}

func main() {
	fmt.Println("Hello, World!")
}
