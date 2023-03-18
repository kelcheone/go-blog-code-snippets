package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a title: ")
	title, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", title)
}
