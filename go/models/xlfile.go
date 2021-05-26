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

		sheet := new(XLSheet).Stage()
		sheet.Name = sh.Name
		sheet.sheet = sh
		sheet.MaxCol = sh.MaxCol
		sheet.MaxRow = sh.MaxRow

		xlfile.Sheets = append(xlfile.Sheets, sheet)
	}
	fmt.Println("----")

	xlfile.NbSheets = len(xlfile.file.Sheets)
}
