package main

import "fmt"

var x int

func idk(x *int) {
	*x = 5
  	fmt.Println(x)
}

func main() {
	x = 90
//	fmt.Println(&x)
	idk(&x)
	fmt.Println(x)
}
