package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := excelize.OpenFile("students.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range rows {
		// skip the title, and header
		if i == 0 || i == 1 {
			continue
		}

		file.SetCellFormula("Sheet1", fmt.Sprintf("H%d", i+1), fmt.Sprintf("SUM(B%d:G%d)", i+1, i+1))
		file.SetCellFormula("Sheet1", fmt.Sprintf("I%d", i+1), fmt.Sprintf("AVERAGE(B%d:G%d)", i+1, i+1))
	}

	if err := file.SaveAs("students.xlsx"); err != nil {
		log.Fatal(err)
	}
}
