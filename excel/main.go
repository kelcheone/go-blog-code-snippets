package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	file := excelize.NewFile()

	headers := []string{"ID", "Name", "Age", "Math", "History", "English", "Science", "Total", "Average"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{1, "John", 30, 80, 90, 70, 60},
		{2, "Alex", 20, 90, 60, 90, 90},
		{3, "Bob", 40, 60, 70, 80, 70},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}
	file.InsertRows("Sheet1", 1, 1)
	// merge cells
	file.MergeCell("Sheet1", "A1", "I1")

	// Enter the title
	file.SetCellValue("Sheet1", "A1", "Students Report")

	if err := file.SaveAs("students.xlsx"); err != nil {
		log.Fatal(err)
	}
}
