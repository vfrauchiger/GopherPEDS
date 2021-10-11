package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func toExcel(table [][]string, filename string) string {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return "Problem at opening file!"
	}
	col := []string{"A", "B", "C", "D", "E", "F"}

	f.NewSheet("TermExtList")

	for num0, slice := range table {
		fmt.Println(len(slice))
		fmt.Printf("num0: %d \n", num0)
		for num1, el := range slice {
			fmt.Println("Wert: " + el)
			fmt.Printf("num1: %d \n", num1)
			cell := col[num1] + strconv.Itoa(num0+1)
			fmt.Println(cell)
			f.SetCellValue("TermExtList", cell, el)

		}
	}
	style, err := f.NewStyle(`{"font":{"bold": true}}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetColWidth("TermExtList", "A", "D", 30.0)
	f.SetCellStyle("TermExtList", "A1", "D1", style)
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return "Problem at saving!"
	}

	return "Success"

}
