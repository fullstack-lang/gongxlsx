package models

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

type XLFile struct {
	Name string

	file *xlsx.File

	NbSheets int
	Sheets   []*XLSheet
}

func (xlfile *XLFile) Open(path string) {

	xlfile.Name = path
	var err error
	xlfile.file, err = xlsx.OpenFile(path)

	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range xlfile.file.Sheets {
		fmt.Println(i, sh.Name)

		xlsheet := new(XLSheet).Stage()
		xlsheet.Name = sh.Name
		xlsheet.sheet = sh
		xlsheet.MaxCol = sh.MaxCol
		xlsheet.MaxRow = sh.MaxRow

		xlfile.Sheets = append(xlfile.Sheets, xlsheet)

		emptyRow := false
		rowIndex := 0
		for !emptyRow {

			cell, err := sh.Cell(rowIndex, 0)
			rowIndex = rowIndex + 1
			if err != nil {
				continue
			}
			if cell.Value == "" {
				emptyRow = true
				continue
			}
			xlsheet.NbRows = rowIndex
		}
		fmt.Println("Sheet ", xlsheet.Name, "Nb Rows", xlsheet.NbRows)
	}
	fmt.Println("----")

	xlfile.NbSheets = len(xlfile.file.Sheets)
}

var nbRows int

func rowVisitor(r *xlsx.Row) error {
	nbRows = nbRows + 1
	return nil
}
