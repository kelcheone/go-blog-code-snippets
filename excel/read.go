package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// read students.xlsx

func main() {
	file, err := excelize.OpenFile("students.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}
