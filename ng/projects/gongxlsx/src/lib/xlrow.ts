// generated code - do not edit

import { XLRowDB } from './xlrow-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLCell } from './xlcell'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLRow {

	static GONGSTRUCT_NAME = "XLRow"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RowIndex: number = 0

	// insertion point for pointers and slices of pointers declarations
	Cells: Array<XLCell> = []
}

export function CopyXLRowToXLRowDB(xlrow: XLRow, xlrowDB: XLRowDB) {

	xlrowDB.CreatedAt = xlrow.CreatedAt
	xlrowDB.DeletedAt = xlrow.DeletedAt
	xlrowDB.ID = xlrow.ID
	
	// insertion point for basic fields copy operations
	xlrowDB.Name = xlrow.Name
	xlrowDB.RowIndex = xlrow.RowIndex

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlrowDB.XLRowPointersEncoding.Cells = []
    for (let _xlcell of xlrow.Cells) {
		xlrowDB.XLRowPointersEncoding.Cells.push(_xlcell.ID)
    }
	
}

export function CopyXLRowDBToXLRow(xlrowDB: XLRowDB, xlrow: XLRow, frontRepo: FrontRepo) {

	xlrow.CreatedAt = xlrowDB.CreatedAt
	xlrow.DeletedAt = xlrowDB.DeletedAt
	xlrow.ID = xlrowDB.ID
	
	// insertion point for basic fields copy operations
	xlrow.Name = xlrowDB.Name
	xlrow.RowIndex = xlrowDB.RowIndex

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlrow.Cells = new Array<XLCell>()
	for (let _id of xlrowDB.XLRowPointersEncoding.Cells) {
	  let _xlcell = frontRepo.XLCells.get(_id)
	  if (_xlcell != undefined) {
		xlrow.Cells.push(_xlcell!)
	  }
	}
}
