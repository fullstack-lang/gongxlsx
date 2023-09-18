// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "DisplaySelection":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " DisplaySelection Form",
			OnSave: NewDisplaySelectionFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, formGroup, playground)
	case "XLCell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLCell Form",
			OnSave: NewXLCellFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, formGroup, playground)
	case "XLFile":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLFile Form",
			OnSave: NewXLFileFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, formGroup, playground)
	case "XLRow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLRow Form",
			OnSave: NewXLRowFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, formGroup, playground)
	case "XLSheet":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLSheet Form",
			OnSave: NewXLSheetFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, formGroup, playground)
	}
	formStage.Commit()
}
