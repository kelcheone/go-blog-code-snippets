package main

import "fmt"

// get user input
func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

// main function
func main() {
	// get the sentence from the user
	fmt.Println("Enter a sentence: ")
	sentence := getUserInput()
	fmt.Println("You entered: ", sentence)
}
