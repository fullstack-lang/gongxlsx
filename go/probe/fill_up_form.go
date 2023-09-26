// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
	"github.com/fullstack-lang/gongxlsx/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("XLFile", instanceWithInferedType.XLFile, formGroup, playground)
		AssociationFieldToForm("XLSheet", instanceWithInferedType.XLSheet, formGroup, playground)

	case *models.XLCell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLRow"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLRow),
					"Cells",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.XLRow, *models.XLCell](
					nil,
					"Cells",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "SheetCells"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLSheet),
					"SheetCells",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.XLSheet, *models.XLCell](
					nil,
					"SheetCells",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.XLFile:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbSheets", instanceWithInferedType.NbSheets, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Sheets", instanceWithInferedType, &instanceWithInferedType.Sheets, formGroup, playground)

	case *models.XLRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RowIndex", instanceWithInferedType.RowIndex, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLSheet),
					"Rows",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.XLSheet, *models.XLRow](
					nil,
					"Rows",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.XLSheet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxRow", instanceWithInferedType.MaxRow, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxCol", instanceWithInferedType.MaxCol, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbRows", instanceWithInferedType.NbRows, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, playground)
		AssociationSliceToForm("SheetCells", instanceWithInferedType, &instanceWithInferedType.SheetCells, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLFile"
			rf.Fieldname = "Sheets"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLFile),
					"Sheets",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.XLFile, *models.XLSheet](
					nil,
					"Sheets",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
