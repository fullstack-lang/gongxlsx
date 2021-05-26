// insertion point for imports
import { XLSheetDB } from './xlsheet-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class XLRowDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	RowIndex?: number
	NbCols?: number

	// insertion point for other declarations
	XLSheet_RowsDBID?: NullInt64
	XLSheet_RowsDBID_Index?: NullInt64 // store the index of the xlrow instance in XLSheet.Rows
	XLSheet_Rows_reverse?: XLSheetDB

}
