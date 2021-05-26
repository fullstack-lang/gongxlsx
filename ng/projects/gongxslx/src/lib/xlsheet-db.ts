// insertion point for imports
import { XLFileDB } from './xlfile-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class XLSheetDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	MaxRow?: number
	MaxCol?: number

	// insertion point for other declarations
	XLFile_SheetsDBID?: NullInt64
	XLFile_SheetsDBID_Index?: NullInt64 // store the index of the xlsheet instance in XLFile.Sheets
	XLFile_Sheets_reverse?: XLFileDB

}
