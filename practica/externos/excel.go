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
	excelFile.SetCellValue("Sheet1", "A1", "Hola, soy una celda hecha en Golang")
	isError := excelFile.SaveAs("ExcelGo.xlsx") // Si existe un error lo retorna
	if isError != nil {
		fmt.Println(isError)
	}
}
