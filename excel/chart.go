package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// add a chart to the file
func main() {
	file, err := excelize.OpenFile("students.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	file.InsertCols("Sheet1", "J", 2)
	file.SetCellValue("Sheet1", "J2", "IDK")

	if err := file.AddChart("sheet1", "h5", &excelize.Chart{
		Type: "line",
		Series: []excelize.ChartSeries{
			{
				Name:       "sheet1!$D$2",
				Categories: `sheet1!$A$3:$A$5`,
				Values:     `sheet1!$D$3:$D$5`,
			},
			{
				Name:       "sheet1!$E$2",
				Categories: `sheet1!$A$3:$A$5`,
				Values:     `sheet1!$E$3:$E$5`,
			},
			{
				Name:       "sheet1!$F$2",
				Categories: `sheet1!$A$3:$A$5`,
				Values:     `sheet1!$F$3:$F$5`,
			},
			{
				Name:       "sheet1!$G$2",
				Categories: `sheet1!$A$3:$A$5`,
				Values:     `sheet1!$G$3:$G$5`,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	if err := file.SaveAs("students.xlsx"); err != nil {
		log.Fatal(err)
	}
}
