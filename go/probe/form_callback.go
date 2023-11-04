// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
	"github.com/fullstack-lang/gongxlsx/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__DisplaySelectionFormCallback(
	displayselection *models.DisplaySelection,
	probe *Probe,
) (displayselectionFormCallback *DisplaySelectionFormCallback) {
	displayselectionFormCallback = new(DisplaySelectionFormCallback)
	displayselectionFormCallback.probe = probe
	displayselectionFormCallback.displayselection = displayselection

	displayselectionFormCallback.CreationMode = (displayselection == nil)

	return
}

type DisplaySelectionFormCallback struct {
	displayselection *models.DisplaySelection

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (displayselectionFormCallback *DisplaySelectionFormCallback) OnSave() {

	log.Println("DisplaySelectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayselectionFormCallback.probe.formStage.Checkout()

	if displayselectionFormCallback.displayselection == nil {
		displayselectionFormCallback.displayselection = new(models.DisplaySelection).Stage(displayselectionFormCallback.probe.stageOfInterest)
	}
	displayselection_ := displayselectionFormCallback.displayselection
	_ = displayselection_

	// get the formGroup
	formGroup := displayselectionFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayselection_.Name), formDiv)
		case "XLFile":
			FormDivSelectFieldToField(&(displayselection_.XLFile), displayselectionFormCallback.probe.stageOfInterest, formDiv)
		case "XLSheet":
			FormDivSelectFieldToField(&(displayselection_.XLSheet), displayselectionFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	displayselectionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DisplaySelection](
		displayselectionFormCallback.probe,
	)
	displayselectionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if displayselectionFormCallback.CreationMode {
		displayselectionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DisplaySelectionFormCallback(
				nil,
				displayselectionFormCallback.probe,
			),
		}).Stage(displayselectionFormCallback.probe.formStage)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, newFormGroup, displayselectionFormCallback.probe)
		displayselectionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(displayselectionFormCallback.probe)
}
func __gong__New__XLCellFormCallback(
	xlcell *models.XLCell,
	probe *Probe,
) (xlcellFormCallback *XLCellFormCallback) {
	xlcellFormCallback = new(XLCellFormCallback)
	xlcellFormCallback.probe = probe
	xlcellFormCallback.xlcell = xlcell

	xlcellFormCallback.CreationMode = (xlcell == nil)

	return
}

type XLCellFormCallback struct {
	xlcell *models.XLCell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (xlcellFormCallback *XLCellFormCallback) OnSave() {

	log.Println("XLCellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlcellFormCallback.probe.formStage.Checkout()

	if xlcellFormCallback.xlcell == nil {
		xlcellFormCallback.xlcell = new(models.XLCell).Stage(xlcellFormCallback.probe.stageOfInterest)
	}
	xlcell_ := xlcellFormCallback.xlcell
	_ = xlcell_

	// get the formGroup
	formGroup := xlcellFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlcell_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(xlcell_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(xlcell_.Y), formDiv)
		case "XLRow:Cells":
			// we need to retrieve the field owner before the change
			var pastXLRowOwner *models.XLRow
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLRow"
			rf.Fieldname = "Cells"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xlcellFormCallback.probe.stageOfInterest,
				xlcellFormCallback.probe.backRepoOfInterest,
				xlcell_,
				&rf)

			if reverseFieldOwner != nil {
				pastXLRowOwner = reverseFieldOwner.(*models.XLRow)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastXLRowOwner != nil {
					idx := slices.Index(pastXLRowOwner.Cells, xlcell_)
					pastXLRowOwner.Cells = slices.Delete(pastXLRowOwner.Cells, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _xlrow := range *models.GetGongstructInstancesSet[models.XLRow](xlcellFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _xlrow.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newXLRowOwner := _xlrow // we have a match
						if pastXLRowOwner != nil {
							if newXLRowOwner != pastXLRowOwner {
								idx := slices.Index(pastXLRowOwner.Cells, xlcell_)
								pastXLRowOwner.Cells = slices.Delete(pastXLRowOwner.Cells, idx, idx+1)
								newXLRowOwner.Cells = append(newXLRowOwner.Cells, xlcell_)
							}
						} else {
							newXLRowOwner.Cells = append(newXLRowOwner.Cells, xlcell_)
						}
					}
				}
			}
		case "XLSheet:SheetCells":
			// we need to retrieve the field owner before the change
			var pastXLSheetOwner *models.XLSheet
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "SheetCells"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xlcellFormCallback.probe.stageOfInterest,
				xlcellFormCallback.probe.backRepoOfInterest,
				xlcell_,
				&rf)

			if reverseFieldOwner != nil {
				pastXLSheetOwner = reverseFieldOwner.(*models.XLSheet)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastXLSheetOwner != nil {
					idx := slices.Index(pastXLSheetOwner.SheetCells, xlcell_)
					pastXLSheetOwner.SheetCells = slices.Delete(pastXLSheetOwner.SheetCells, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlcellFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _xlsheet.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newXLSheetOwner := _xlsheet // we have a match
						if pastXLSheetOwner != nil {
							if newXLSheetOwner != pastXLSheetOwner {
								idx := slices.Index(pastXLSheetOwner.SheetCells, xlcell_)
								pastXLSheetOwner.SheetCells = slices.Delete(pastXLSheetOwner.SheetCells, idx, idx+1)
								newXLSheetOwner.SheetCells = append(newXLSheetOwner.SheetCells, xlcell_)
							}
						} else {
							newXLSheetOwner.SheetCells = append(newXLSheetOwner.SheetCells, xlcell_)
						}
					}
				}
			}
		}
	}

	xlcellFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XLCell](
		xlcellFormCallback.probe,
	)
	xlcellFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlcellFormCallback.CreationMode {
		xlcellFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLCellFormCallback(
				nil,
				xlcellFormCallback.probe,
			),
		}).Stage(xlcellFormCallback.probe.formStage)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, newFormGroup, xlcellFormCallback.probe)
		xlcellFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xlcellFormCallback.probe)
}
func __gong__New__XLFileFormCallback(
	xlfile *models.XLFile,
	probe *Probe,
) (xlfileFormCallback *XLFileFormCallback) {
	xlfileFormCallback = new(XLFileFormCallback)
	xlfileFormCallback.probe = probe
	xlfileFormCallback.xlfile = xlfile

	xlfileFormCallback.CreationMode = (xlfile == nil)

	return
}

type XLFileFormCallback struct {
	xlfile *models.XLFile

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (xlfileFormCallback *XLFileFormCallback) OnSave() {

	log.Println("XLFileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlfileFormCallback.probe.formStage.Checkout()

	if xlfileFormCallback.xlfile == nil {
		xlfileFormCallback.xlfile = new(models.XLFile).Stage(xlfileFormCallback.probe.stageOfInterest)
	}
	xlfile_ := xlfileFormCallback.xlfile
	_ = xlfile_

	// get the formGroup
	formGroup := xlfileFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlfile_.Name), formDiv)
		case "NbSheets":
			FormDivBasicFieldToField(&(xlfile_.NbSheets), formDiv)
		}
	}

	xlfileFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XLFile](
		xlfileFormCallback.probe,
	)
	xlfileFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlfileFormCallback.CreationMode {
		xlfileFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLFileFormCallback(
				nil,
				xlfileFormCallback.probe,
			),
		}).Stage(xlfileFormCallback.probe.formStage)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, newFormGroup, xlfileFormCallback.probe)
		xlfileFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xlfileFormCallback.probe)
}
func __gong__New__XLRowFormCallback(
	xlrow *models.XLRow,
	probe *Probe,
) (xlrowFormCallback *XLRowFormCallback) {
	xlrowFormCallback = new(XLRowFormCallback)
	xlrowFormCallback.probe = probe
	xlrowFormCallback.xlrow = xlrow

	xlrowFormCallback.CreationMode = (xlrow == nil)

	return
}

type XLRowFormCallback struct {
	xlrow *models.XLRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (xlrowFormCallback *XLRowFormCallback) OnSave() {

	log.Println("XLRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlrowFormCallback.probe.formStage.Checkout()

	if xlrowFormCallback.xlrow == nil {
		xlrowFormCallback.xlrow = new(models.XLRow).Stage(xlrowFormCallback.probe.stageOfInterest)
	}
	xlrow_ := xlrowFormCallback.xlrow
	_ = xlrow_

	// get the formGroup
	formGroup := xlrowFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlrow_.Name), formDiv)
		case "RowIndex":
			FormDivBasicFieldToField(&(xlrow_.RowIndex), formDiv)
		case "XLSheet:Rows":
			// we need to retrieve the field owner before the change
			var pastXLSheetOwner *models.XLSheet
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLSheet"
			rf.Fieldname = "Rows"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xlrowFormCallback.probe.stageOfInterest,
				xlrowFormCallback.probe.backRepoOfInterest,
				xlrow_,
				&rf)

			if reverseFieldOwner != nil {
				pastXLSheetOwner = reverseFieldOwner.(*models.XLSheet)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastXLSheetOwner != nil {
					idx := slices.Index(pastXLSheetOwner.Rows, xlrow_)
					pastXLSheetOwner.Rows = slices.Delete(pastXLSheetOwner.Rows, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlrowFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _xlsheet.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newXLSheetOwner := _xlsheet // we have a match
						if pastXLSheetOwner != nil {
							if newXLSheetOwner != pastXLSheetOwner {
								idx := slices.Index(pastXLSheetOwner.Rows, xlrow_)
								pastXLSheetOwner.Rows = slices.Delete(pastXLSheetOwner.Rows, idx, idx+1)
								newXLSheetOwner.Rows = append(newXLSheetOwner.Rows, xlrow_)
							}
						} else {
							newXLSheetOwner.Rows = append(newXLSheetOwner.Rows, xlrow_)
						}
					}
				}
			}
		}
	}

	xlrowFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XLRow](
		xlrowFormCallback.probe,
	)
	xlrowFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlrowFormCallback.CreationMode {
		xlrowFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLRowFormCallback(
				nil,
				xlrowFormCallback.probe,
			),
		}).Stage(xlrowFormCallback.probe.formStage)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, newFormGroup, xlrowFormCallback.probe)
		xlrowFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xlrowFormCallback.probe)
}
func __gong__New__XLSheetFormCallback(
	xlsheet *models.XLSheet,
	probe *Probe,
) (xlsheetFormCallback *XLSheetFormCallback) {
	xlsheetFormCallback = new(XLSheetFormCallback)
	xlsheetFormCallback.probe = probe
	xlsheetFormCallback.xlsheet = xlsheet

	xlsheetFormCallback.CreationMode = (xlsheet == nil)

	return
}

type XLSheetFormCallback struct {
	xlsheet *models.XLSheet

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (xlsheetFormCallback *XLSheetFormCallback) OnSave() {

	log.Println("XLSheetFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlsheetFormCallback.probe.formStage.Checkout()

	if xlsheetFormCallback.xlsheet == nil {
		xlsheetFormCallback.xlsheet = new(models.XLSheet).Stage(xlsheetFormCallback.probe.stageOfInterest)
	}
	xlsheet_ := xlsheetFormCallback.xlsheet
	_ = xlsheet_

	// get the formGroup
	formGroup := xlsheetFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlsheet_.Name), formDiv)
		case "MaxRow":
			FormDivBasicFieldToField(&(xlsheet_.MaxRow), formDiv)
		case "MaxCol":
			FormDivBasicFieldToField(&(xlsheet_.MaxCol), formDiv)
		case "NbRows":
			FormDivBasicFieldToField(&(xlsheet_.NbRows), formDiv)
		case "XLFile:Sheets":
			// we need to retrieve the field owner before the change
			var pastXLFileOwner *models.XLFile
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "XLFile"
			rf.Fieldname = "Sheets"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				xlsheetFormCallback.probe.stageOfInterest,
				xlsheetFormCallback.probe.backRepoOfInterest,
				xlsheet_,
				&rf)

			if reverseFieldOwner != nil {
				pastXLFileOwner = reverseFieldOwner.(*models.XLFile)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastXLFileOwner != nil {
					idx := slices.Index(pastXLFileOwner.Sheets, xlsheet_)
					pastXLFileOwner.Sheets = slices.Delete(pastXLFileOwner.Sheets, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _xlfile := range *models.GetGongstructInstancesSet[models.XLFile](xlsheetFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _xlfile.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newXLFileOwner := _xlfile // we have a match
						if pastXLFileOwner != nil {
							if newXLFileOwner != pastXLFileOwner {
								idx := slices.Index(pastXLFileOwner.Sheets, xlsheet_)
								pastXLFileOwner.Sheets = slices.Delete(pastXLFileOwner.Sheets, idx, idx+1)
								newXLFileOwner.Sheets = append(newXLFileOwner.Sheets, xlsheet_)
							}
						} else {
							newXLFileOwner.Sheets = append(newXLFileOwner.Sheets, xlsheet_)
						}
					}
				}
			}
		}
	}

	xlsheetFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.XLSheet](
		xlsheetFormCallback.probe,
	)
	xlsheetFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if xlsheetFormCallback.CreationMode {
		xlsheetFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLSheetFormCallback(
				nil,
				xlsheetFormCallback.probe,
			),
		}).Stage(xlsheetFormCallback.probe.formStage)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, newFormGroup, xlsheetFormCallback.probe)
		xlsheetFormCallback.probe.formStage.Commit()
	}

	fillUpTree(xlsheetFormCallback.probe)
}
