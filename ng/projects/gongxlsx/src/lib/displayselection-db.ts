// insertion point for imports
import { XLFileDB } from './xlfile-db'
import { XLSheetDB } from './xlsheet-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplaySelectionDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	XLFile?: XLFileDB
	XLFileID: NullInt64 = new NullInt64 // if pointer is null, XLFile.ID = 0

	XLSheet?: XLSheetDB
	XLSheetID: NullInt64 = new NullInt64 // if pointer is null, XLSheet.ID = 0

}
