package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func funcToWithIO() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func funcToWithBufio() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data := make([]byte, 100)
	for {
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func createFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 1000000; i++ {
		file.Write([]byte("Hello World!"))
	}
}
