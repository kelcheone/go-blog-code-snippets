package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	reader.Comma = '|'

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, each := range csvData {
		fmt.Println(each)
	}
}
