// generated code - do not edit

import { DisplaySelectionDB } from './displayselection-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLFile } from './xlfile'
import { XLSheet } from './xlsheet'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DisplaySelection {

	static GONGSTRUCT_NAME = "DisplaySelection"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	XLFile?: XLFile

	XLSheet?: XLSheet

}

export function CopyDisplaySelectionToDisplaySelectionDB(displayselection: DisplaySelection, displayselectionDB: DisplaySelectionDB) {

	displayselectionDB.CreatedAt = displayselection.CreatedAt
	displayselectionDB.DeletedAt = displayselection.DeletedAt
	displayselectionDB.ID = displayselection.ID

	// insertion point for basic fields copy operations
	displayselectionDB.Name = displayselection.Name

	// insertion point for pointer fields encoding
	displayselectionDB.DisplaySelectionPointersEncoding.XLFileID.Valid = true
	if (displayselection.XLFile != undefined) {
		displayselectionDB.DisplaySelectionPointersEncoding.XLFileID.Int64 = displayselection.XLFile.ID  
	} else {
		displayselectionDB.DisplaySelectionPointersEncoding.XLFileID.Int64 = 0 		
	}

	displayselectionDB.DisplaySelectionPointersEncoding.XLSheetID.Valid = true
	if (displayselection.XLSheet != undefined) {
		displayselectionDB.DisplaySelectionPointersEncoding.XLSheetID.Int64 = displayselection.XLSheet.ID  
	} else {
		displayselectionDB.DisplaySelectionPointersEncoding.XLSheetID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyDisplaySelectionDBToDisplaySelection update basic, pointers and slice of pointers fields of displayselection
// from respectively the basic fields and encoded fields of pointers and slices of pointers of displayselectionDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDisplaySelectionDBToDisplaySelection(displayselectionDB: DisplaySelectionDB, displayselection: DisplaySelection, frontRepo: FrontRepo) {

	displayselection.CreatedAt = displayselectionDB.CreatedAt
	displayselection.DeletedAt = displayselectionDB.DeletedAt
	displayselection.ID = displayselectionDB.ID

	// insertion point for basic fields copy operations
	displayselection.Name = displayselectionDB.Name

	// insertion point for pointer fields encoding
	displayselection.XLFile = frontRepo.map_ID_XLFile.get(displayselectionDB.DisplaySelectionPointersEncoding.XLFileID.Int64)
	displayselection.XLSheet = frontRepo.map_ID_XLSheet.get(displayselectionDB.DisplaySelectionPointersEncoding.XLSheetID.Int64)

	// insertion point for slice of pointers fields encoding
}
