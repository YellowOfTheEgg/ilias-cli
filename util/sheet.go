package util

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"ilias-cli/ilias_api"
)

func CreateCorrectionSheet(name string, corrections []ilias_api.Correction) *excelize.File {
	file := excelize.NewFile()
	file.SetSheetName("Sheet1", name)
	file.SetCellValue(name, "A1", "Benutzer")
	file.SetCellValue(name, "B1", "Punktzahl")

	for i, correction := range corrections {
		file.SetCellValue(name, fmt.Sprintf("A%d", i + 2), correction.Student)
		file.SetCellValue(name, fmt.Sprintf("B%d", i + 2), correction.Points)
	}

	return file
}
