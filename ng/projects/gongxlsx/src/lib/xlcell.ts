// generated code - do not edit

import { XLCellDB } from './xlcell-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLCell {

	static GONGSTRUCT_NAME = "XLCell"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X: number = 0
	Y: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyXLCellToXLCellDB(xlcell: XLCell, xlcellDB: XLCellDB) {

	xlcellDB.CreatedAt = xlcell.CreatedAt
	xlcellDB.DeletedAt = xlcell.DeletedAt
	xlcellDB.ID = xlcell.ID
	
	// insertion point for basic fields copy operations
	xlcellDB.Name = xlcell.Name
	xlcellDB.X = xlcell.X
	xlcellDB.Y = xlcell.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyXLCellDBToXLCell(xlcellDB: XLCellDB, xlcell: XLCell, frontRepo: FrontRepo) {

	xlcell.CreatedAt = xlcellDB.CreatedAt
	xlcell.DeletedAt = xlcellDB.DeletedAt
	xlcell.ID = xlcellDB.ID
	
	// insertion point for basic fields copy operations
	xlcell.Name = xlcellDB.Name
	xlcell.X = xlcellDB.X
	xlcell.Y = xlcellDB.Y

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
