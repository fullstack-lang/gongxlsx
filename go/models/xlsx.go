package models

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

type Xslx struct {
	Name string

	file xlsx.File

	NbSheets int
}

func (xslx *Xslx) Open(path string) {

	file, err := xlsx.OpenFile(path)

	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range file.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")
}
