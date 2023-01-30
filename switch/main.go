package main

import (
	"fmt"
)

func main() {

	slice := []int{1, 2, 3}
	slice = append([]int{0}, slice...)
	fmt.Println(slice)
}

// In
