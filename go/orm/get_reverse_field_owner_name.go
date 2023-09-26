// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxlsx/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		tmp := GetInstanceDBFromInstance[models.DisplaySelection, DisplaySelectionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}

	case *models.XLCell:
		tmp := GetInstanceDBFromInstance[models.XLCell, XLCellDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			case "Cells":
				if tmp != nil && tmp.XLRow_CellsDBID.Int64 != 0 {
					id := uint(tmp.XLRow_CellsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				if tmp != nil && tmp.XLSheet_SheetCellsDBID.Int64 != 0 {
					id := uint(tmp.XLSheet_SheetCellsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.XLFile:
		tmp := GetInstanceDBFromInstance[models.XLFile, XLFileDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}

	case *models.XLRow:
		tmp := GetInstanceDBFromInstance[models.XLRow, XLRowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "Rows":
				if tmp != nil && tmp.XLSheet_RowsDBID.Int64 != 0 {
					id := uint(tmp.XLSheet_RowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.XLSheet:
		tmp := GetInstanceDBFromInstance[models.XLSheet, XLSheetDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			case "Sheets":
				if tmp != nil && tmp.XLFile_SheetsDBID.Int64 != 0 {
					id := uint(tmp.XLFile_SheetsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		tmp := GetInstanceDBFromInstance[models.DisplaySelection, DisplaySelectionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.XLCell:
		tmp := GetInstanceDBFromInstance[models.XLCell, XLCellDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			case "Cells":
				if tmp != nil && tmp.XLRow_CellsDBID.Int64 != 0 {
					id := uint(tmp.XLRow_CellsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[id]
					res = reservePointerTarget
				}
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				if tmp != nil && tmp.XLSheet_SheetCellsDBID.Int64 != 0 {
					id := uint(tmp.XLSheet_SheetCellsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
					res = reservePointerTarget
				}
			}
		}
	
	case *models.XLFile:
		tmp := GetInstanceDBFromInstance[models.XLFile, XLFileDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}
	
	case *models.XLRow:
		tmp := GetInstanceDBFromInstance[models.XLRow, XLRowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "Rows":
				if tmp != nil && tmp.XLSheet_RowsDBID.Int64 != 0 {
					id := uint(tmp.XLSheet_RowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
					res = reservePointerTarget
				}
			}
		}
	
	case *models.XLSheet:
		tmp := GetInstanceDBFromInstance[models.XLSheet, XLSheetDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "DisplaySelection":
			switch reverseField.Fieldname {
			}
		case "XLCell":
			switch reverseField.Fieldname {
			}
		case "XLFile":
			switch reverseField.Fieldname {
			case "Sheets":
				if tmp != nil && tmp.XLFile_SheetsDBID.Int64 != 0 {
					id := uint(tmp.XLFile_SheetsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[id]
					res = reservePointerTarget
				}
			}
		case "XLRow":
			switch reverseField.Fieldname {
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
