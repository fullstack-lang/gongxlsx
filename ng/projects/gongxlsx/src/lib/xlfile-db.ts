// insertion point for imports
import { XLSheetDB } from './xlsheet-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class XLFileDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	NbSheets?: number

	// insertion point for other declarations
	Sheets?: Array<XLSheetDB>
}
