// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
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
			OnSave: __gong__New__DisplaySelectionFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, formGroup, probe)
	case "XLCell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLCell Form",
			OnSave: __gong__New__XLCellFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, formGroup, probe)
	case "XLFile":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLFile Form",
			OnSave: __gong__New__XLFileFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, formGroup, probe)
	case "XLRow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLRow Form",
			OnSave: __gong__New__XLRowFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, formGroup, probe)
	case "XLSheet":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " XLSheet Form",
			OnSave: __gong__New__XLSheetFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, formGroup, probe)
	}
	formStage.Commit()
}
