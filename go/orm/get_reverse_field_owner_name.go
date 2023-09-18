// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongxlsx/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseFieldName string) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		tmp := GetInstanceDBFromInstance[models.DisplaySelection, DisplaySelectionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.XLCell:
		tmp := GetInstanceDBFromInstance[models.XLCell, XLCellDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Cells":
			if tmp.XLRow_CellsDBID.Int64 != 0 {
				id := uint(tmp.XLRow_CellsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoXLRow.Map_XLRowDBID_XLRowPtr[id]
				res = reservePointerTarget.Name
			}
		case "SheetCells":
			if tmp.XLSheet_SheetCellsDBID.Int64 != 0 {
				id := uint(tmp.XLSheet_SheetCellsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.XLFile:
		tmp := GetInstanceDBFromInstance[models.XLFile, XLFileDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		}

	case *models.XLRow:
		tmp := GetInstanceDBFromInstance[models.XLRow, XLRowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Rows":
			if tmp.XLSheet_RowsDBID.Int64 != 0 {
				id := uint(tmp.XLSheet_RowsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[id]
				res = reservePointerTarget.Name
			}
		}

	case *models.XLSheet:
		tmp := GetInstanceDBFromInstance[models.XLSheet, XLSheetDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseFieldName {
		// insertion point
		case "Sheets":
			if tmp.XLFile_SheetsDBID.Int64 != 0 {
				id := uint(tmp.XLFile_SheetsDBID.Int64)
				reservePointerTarget := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[id]
				res = reservePointerTarget.Name
			}
		}

	default:
		_ = inst
	}
	return
}
