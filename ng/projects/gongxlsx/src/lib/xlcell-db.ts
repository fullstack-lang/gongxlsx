// insertion point for imports
import { XLRowDB } from './xlrow-db'
import { XLSheetDB } from './xlsheet-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class XLCellDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	X?: number
	Y?: number

	// insertion point for other declarations
	XLRow_CellsDBID?: NullInt64
	XLRow_CellsDBID_Index?: NullInt64 // store the index of the xlcell instance in XLRow.Cells
	XLRow_Cells_reverse?: XLRowDB

	XLSheet_SheetCellsDBID?: NullInt64
	XLSheet_SheetCellsDBID_Index?: NullInt64 // store the index of the xlcell instance in XLSheet.SheetCells
	XLSheet_SheetCells_reverse?: XLSheetDB

}
