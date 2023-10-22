// insertion point for imports
import { XLCellDB } from './xlcell-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLRowDB {

	static GONGSTRUCT_NAME = "XLRow"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	RowIndex: number = 0

	// insertion point for pointers and slices of pointers declarations
	Cells: Array<XLCellDB> = []

	XLRowPointersEncoding: XLRowPointersEncoding = new XLRowPointersEncoding
}

export class XLRowPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Cells: number[] = []
}
