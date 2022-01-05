package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	creacionLibroExcel()
}

func creacionLibroExcel() {
	excelFile := excelize.NewFile()
	excelSheet := excelFile.NewSheet("Main Sheet")
	excelFile.SetCellValue("Main Sheet", "A1", "Hola, soy una celda hecha en Golang")
	excelFile.SetActiveSheet(excelSheet)
	if error := excelFile.SaveAs("ExcelGo.xlsx"); error != nil {
		fmt.Println(error)
	}
}
