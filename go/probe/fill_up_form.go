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
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("XLFile", instanceWithInferedType.XLFile, formGroup, probe)
		AssociationFieldToForm("XLSheet", instanceWithInferedType.XLSheet, formGroup, probe)

	case *models.XLCell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLRow"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLRow),
					"Cells",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.XLRow, *models.XLCell](
					nil,
					"Cells",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "SheetCells"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLSheet),
					"SheetCells",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.XLSheet, *models.XLCell](
					nil,
					"SheetCells",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.XLFile:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NbSheets", instanceWithInferedType.NbSheets, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("Sheets", instanceWithInferedType, &instanceWithInferedType.Sheets, formGroup, probe)

	case *models.XLRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("RowIndex", instanceWithInferedType.RowIndex, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLSheet),
					"Rows",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.XLSheet, *models.XLRow](
					nil,
					"Rows",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.XLSheet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("MaxRow", instanceWithInferedType.MaxRow, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("MaxCol", instanceWithInferedType.MaxCol, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("NbRows", instanceWithInferedType.NbRows, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, probe)
		AssociationSliceToForm("SheetCells", instanceWithInferedType, &instanceWithInferedType.SheetCells, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLFile"
			rf.Fieldname = "Sheets"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.XLFile),
					"Sheets",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.XLFile, *models.XLSheet](
					nil,
					"Sheets",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
