// generated code - do not edit

import { XLSheetDB } from './xlsheet-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLRow } from './xlrow'
import { XLCell } from './xlcell'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLSheet {

	static GONGSTRUCT_NAME = "XLSheet"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	MaxRow: number = 0
	MaxCol: number = 0
	NbRows: number = 0

	// insertion point for pointers and slices of pointers declarations
	Rows: Array<XLRow> = []
	SheetCells: Array<XLCell> = []
}

export function CopyXLSheetToXLSheetDB(xlsheet: XLSheet, xlsheetDB: XLSheetDB) {

	xlsheetDB.CreatedAt = xlsheet.CreatedAt
	xlsheetDB.DeletedAt = xlsheet.DeletedAt
	xlsheetDB.ID = xlsheet.ID

	// insertion point for basic fields copy operations
	xlsheetDB.Name = xlsheet.Name
	xlsheetDB.MaxRow = xlsheet.MaxRow
	xlsheetDB.MaxCol = xlsheet.MaxCol
	xlsheetDB.NbRows = xlsheet.NbRows

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlsheetDB.XLSheetPointersEncoding.Rows = []
	for (let _xlrow of xlsheet.Rows) {
		xlsheetDB.XLSheetPointersEncoding.Rows.push(_xlrow.ID)
	}

	xlsheetDB.XLSheetPointersEncoding.SheetCells = []
	for (let _xlcell of xlsheet.SheetCells) {
		xlsheetDB.XLSheetPointersEncoding.SheetCells.push(_xlcell.ID)
	}

}

// CopyXLSheetDBToXLSheet update basic, pointers and slice of pointers fields of xlsheet
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlsheetDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLSheetDBToXLSheet(xlsheetDB: XLSheetDB, xlsheet: XLSheet, frontRepo: FrontRepo) {

	xlsheet.CreatedAt = xlsheetDB.CreatedAt
	xlsheet.DeletedAt = xlsheetDB.DeletedAt
	xlsheet.ID = xlsheetDB.ID

	// insertion point for basic fields copy operations
	xlsheet.Name = xlsheetDB.Name
	xlsheet.MaxRow = xlsheetDB.MaxRow
	xlsheet.MaxCol = xlsheetDB.MaxCol
	xlsheet.NbRows = xlsheetDB.NbRows

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlsheet.Rows = new Array<XLRow>()
	for (let _id of xlsheetDB.XLSheetPointersEncoding.Rows) {
		let _xlrow = frontRepo.map_ID_XLRow.get(_id)
		if (_xlrow != undefined) {
			xlsheet.Rows.push(_xlrow!)
		}
	}
	xlsheet.SheetCells = new Array<XLCell>()
	for (let _id of xlsheetDB.XLSheetPointersEncoding.SheetCells) {
		let _xlcell = frontRepo.map_ID_XLCell.get(_id)
		if (_xlcell != undefined) {
			xlsheet.SheetCells.push(_xlcell!)
		}
	}
}
