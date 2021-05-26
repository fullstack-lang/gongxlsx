package models

import (
	"github.com/tealeg/xlsx/v3"
)

type XLSheet struct {
	Name   string
	MaxRow int
	MaxCol int

	sheet  *xlsx.Sheet
	NbRows int

	Rows       []*XLRow
	SheetCells []*XLCell
}
