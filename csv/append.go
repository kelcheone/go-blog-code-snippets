package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// append data to csv file
func main() {
	file, err := os.OpenFile("data2.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"Felix", "23", "Berlin", "Researcher", "PWC", "555-230-2139"},
		{"Linda", "28", "Chicago", "Data Engineere", "KPMG", "331-359-8311"},
	}

	for _, value := range data {
		writer.Write(value)
	}

	fmt.Println("Data appended to csv file successfully.")
}
