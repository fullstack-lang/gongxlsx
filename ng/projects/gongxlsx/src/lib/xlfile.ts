// generated code - do not edit

import { XLFileDB } from './xlfile-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { XLSheet } from './xlsheet'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class XLFile {

	static GONGSTRUCT_NAME = "XLFile"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NbSheets: number = 0

	// insertion point for pointers and slices of pointers declarations
	Sheets: Array<XLSheet> = []
}

export function CopyXLFileToXLFileDB(xlfile: XLFile, xlfileDB: XLFileDB) {

	xlfileDB.CreatedAt = xlfile.CreatedAt
	xlfileDB.DeletedAt = xlfile.DeletedAt
	xlfileDB.ID = xlfile.ID

	// insertion point for basic fields copy operations
	xlfileDB.Name = xlfile.Name
	xlfileDB.NbSheets = xlfile.NbSheets

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlfileDB.XLFilePointersEncoding.Sheets = []
	for (let _xlsheet of xlfile.Sheets) {
		xlfileDB.XLFilePointersEncoding.Sheets.push(_xlsheet.ID)
	}

}

// CopyXLFileDBToXLFile update basic, pointers and slice of pointers fields of xlfile
// from respectively the basic fields and encoded fields of pointers and slices of pointers of xlfileDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyXLFileDBToXLFile(xlfileDB: XLFileDB, xlfile: XLFile, frontRepo: FrontRepo) {

	xlfile.CreatedAt = xlfileDB.CreatedAt
	xlfile.DeletedAt = xlfileDB.DeletedAt
	xlfile.ID = xlfileDB.ID

	// insertion point for basic fields copy operations
	xlfile.Name = xlfileDB.Name
	xlfile.NbSheets = xlfileDB.NbSheets

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	xlfile.Sheets = new Array<XLSheet>()
	for (let _id of xlfileDB.XLFilePointersEncoding.Sheets) {
		let _xlsheet = frontRepo.map_ID_XLSheet.get(_id)
		if (_xlsheet != undefined) {
			xlfile.Sheets.push(_xlsheet!)
		}
	}
}
