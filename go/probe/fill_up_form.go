// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("XLFile", instanceWithInferedType.XLFile, formGroup, playground)
		AssociationFieldToForm("XLSheet", instanceWithInferedType.XLSheet, formGroup, playground)

	case *models.XLCell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, playground.formStage, formGroup)

	case *models.XLFile:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("NbSheets", instanceWithInferedType.NbSheets, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Sheets", instanceWithInferedType, &instanceWithInferedType.Sheets, formGroup, playground)

	case *models.XLRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("RowIndex", instanceWithInferedType.RowIndex, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, playground)

	case *models.XLSheet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxRow", instanceWithInferedType.MaxRow, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxCol", instanceWithInferedType.MaxCol, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("NbRows", instanceWithInferedType.NbRows, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, playground)
		AssociationSliceToForm("SheetCells", instanceWithInferedType, &instanceWithInferedType.SheetCells, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
