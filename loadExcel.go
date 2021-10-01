package main

// small function to load an Excel-file, extract US publication numbers
// and return a slice of US publication numbers

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func LoadExcel(filename string) []string {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// get name of first worksheet
	table := f.GetSheetName(0)

	//read all cols
	cols, err := f.GetCols(table)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(cols[0])

	var publnos []string
	for _, cell := range cols[0] {
		if cell[:2] == "US" {
			publnos = append(publnos, cell)
		}

	}

	return publnos
}
