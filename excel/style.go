package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// Style the file
func main() {

	file, err := excelize.OpenFile("students.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	style, err := file.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{Horizontal: "center"},
			Font:      &excelize.Font{Bold: true, Color: "FF0000"},
			Border: []excelize.Border{
				{Type: "left", Color: "00FF0000", Style: 1},
				{Type: "right", Color: "00FF0000", Style: 1},
				{Type: "top", Color: "00FF0000", Style: 1},
				{Type: "bottom", Color: "00FF0000", Style: 1},
			},
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{"#DDEBF7"},
				Pattern: 1,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	styleTitle, err := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type: "pattern",
			// navy blue
			Color:   []string{"#000080"},
			Pattern: 1,
		},
		Font:      &excelize.Font{Bold: true, Color: "FFFFFF"},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	file.SetCellStyle("Sheet1", "A1", "I5", style)
	file.SetCellStyle("Sheet1", "A1", "I1", styleTitle)

	if err := file.Save(); err != nil {
		log.Fatal(err)
	}
}
