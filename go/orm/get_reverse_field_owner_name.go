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
				if _xlrow, ok := stage.XLRow_Cells_reverseMap[inst]; ok {
					res = _xlrow.Name
				}
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				if _xlsheet, ok := stage.XLSheet_SheetCells_reverseMap[inst]; ok {
					res = _xlsheet.Name
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
				if _xlsheet, ok := stage.XLSheet_Rows_reverseMap[inst]; ok {
					res = _xlsheet.Name
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
				if _xlfile, ok := stage.XLFile_Sheets_reverseMap[inst]; ok {
					res = _xlfile.Name
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
				res = stage.XLRow_Cells_reverseMap[inst]
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				res = stage.XLSheet_SheetCells_reverseMap[inst]
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
				res = stage.XLSheet_Rows_reverseMap[inst]
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
				res = stage.XLFile_Sheets_reverseMap[inst]
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