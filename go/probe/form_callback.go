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
	playground *Playground,
) (displayselectionFormCallback *DisplaySelectionFormCallback) {
	displayselectionFormCallback = new(DisplaySelectionFormCallback)
	displayselectionFormCallback.playground = playground
	displayselectionFormCallback.displayselection = displayselection

	displayselectionFormCallback.CreationMode = (displayselection == nil)

	return
}

type DisplaySelectionFormCallback struct {
	displayselection *models.DisplaySelection

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (displayselectionFormCallback *DisplaySelectionFormCallback) OnSave() {

	log.Println("DisplaySelectionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	displayselectionFormCallback.playground.formStage.Checkout()

	if displayselectionFormCallback.displayselection == nil {
		displayselectionFormCallback.displayselection = new(models.DisplaySelection).Stage(displayselectionFormCallback.playground.stageOfInterest)
	}
	displayselection_ := displayselectionFormCallback.displayselection
	_ = displayselection_

	// get the formGroup
	formGroup := displayselectionFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(displayselection_.Name), formDiv)
		case "XLFile":
			FormDivSelectFieldToField(&(displayselection_.XLFile), displayselectionFormCallback.playground.stageOfInterest, formDiv)
		case "XLSheet":
			FormDivSelectFieldToField(&(displayselection_.XLSheet), displayselectionFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	displayselectionFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.DisplaySelection](
		displayselectionFormCallback.playground,
	)
	displayselectionFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if displayselectionFormCallback.CreationMode {
		displayselectionFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__DisplaySelectionFormCallback(
				nil,
				displayselectionFormCallback.playground,
			),
		}).Stage(displayselectionFormCallback.playground.formStage)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, newFormGroup, displayselectionFormCallback.playground)
		displayselectionFormCallback.playground.formStage.Commit()
	}

	fillUpTree(displayselectionFormCallback.playground)
}
func __gong__New__XLCellFormCallback(
	xlcell *models.XLCell,
	playground *Playground,
) (xlcellFormCallback *XLCellFormCallback) {
	xlcellFormCallback = new(XLCellFormCallback)
	xlcellFormCallback.playground = playground
	xlcellFormCallback.xlcell = xlcell

	xlcellFormCallback.CreationMode = (xlcell == nil)

	return
}

type XLCellFormCallback struct {
	xlcell *models.XLCell

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (xlcellFormCallback *XLCellFormCallback) OnSave() {

	log.Println("XLCellFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlcellFormCallback.playground.formStage.Checkout()

	if xlcellFormCallback.xlcell == nil {
		xlcellFormCallback.xlcell = new(models.XLCell).Stage(xlcellFormCallback.playground.stageOfInterest)
	}
	xlcell_ := xlcellFormCallback.xlcell
	_ = xlcell_

	// get the formGroup
	formGroup := xlcellFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				xlcellFormCallback.playground.stageOfInterest,
				xlcellFormCallback.playground.backRepoOfInterest,
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
				for _xlrow := range *models.GetGongstructInstancesSet[models.XLRow](xlcellFormCallback.playground.stageOfInterest) {

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
				xlcellFormCallback.playground.stageOfInterest,
				xlcellFormCallback.playground.backRepoOfInterest,
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
				for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlcellFormCallback.playground.stageOfInterest) {

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

	xlcellFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.XLCell](
		xlcellFormCallback.playground,
	)
	xlcellFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if xlcellFormCallback.CreationMode {
		xlcellFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLCellFormCallback(
				nil,
				xlcellFormCallback.playground,
			),
		}).Stage(xlcellFormCallback.playground.formStage)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, newFormGroup, xlcellFormCallback.playground)
		xlcellFormCallback.playground.formStage.Commit()
	}

	fillUpTree(xlcellFormCallback.playground)
}
func __gong__New__XLFileFormCallback(
	xlfile *models.XLFile,
	playground *Playground,
) (xlfileFormCallback *XLFileFormCallback) {
	xlfileFormCallback = new(XLFileFormCallback)
	xlfileFormCallback.playground = playground
	xlfileFormCallback.xlfile = xlfile

	xlfileFormCallback.CreationMode = (xlfile == nil)

	return
}

type XLFileFormCallback struct {
	xlfile *models.XLFile

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (xlfileFormCallback *XLFileFormCallback) OnSave() {

	log.Println("XLFileFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlfileFormCallback.playground.formStage.Checkout()

	if xlfileFormCallback.xlfile == nil {
		xlfileFormCallback.xlfile = new(models.XLFile).Stage(xlfileFormCallback.playground.stageOfInterest)
	}
	xlfile_ := xlfileFormCallback.xlfile
	_ = xlfile_

	// get the formGroup
	formGroup := xlfileFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(xlfile_.Name), formDiv)
		case "NbSheets":
			FormDivBasicFieldToField(&(xlfile_.NbSheets), formDiv)
		}
	}

	xlfileFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.XLFile](
		xlfileFormCallback.playground,
	)
	xlfileFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if xlfileFormCallback.CreationMode {
		xlfileFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLFileFormCallback(
				nil,
				xlfileFormCallback.playground,
			),
		}).Stage(xlfileFormCallback.playground.formStage)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, newFormGroup, xlfileFormCallback.playground)
		xlfileFormCallback.playground.formStage.Commit()
	}

	fillUpTree(xlfileFormCallback.playground)
}
func __gong__New__XLRowFormCallback(
	xlrow *models.XLRow,
	playground *Playground,
) (xlrowFormCallback *XLRowFormCallback) {
	xlrowFormCallback = new(XLRowFormCallback)
	xlrowFormCallback.playground = playground
	xlrowFormCallback.xlrow = xlrow

	xlrowFormCallback.CreationMode = (xlrow == nil)

	return
}

type XLRowFormCallback struct {
	xlrow *models.XLRow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (xlrowFormCallback *XLRowFormCallback) OnSave() {

	log.Println("XLRowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlrowFormCallback.playground.formStage.Checkout()

	if xlrowFormCallback.xlrow == nil {
		xlrowFormCallback.xlrow = new(models.XLRow).Stage(xlrowFormCallback.playground.stageOfInterest)
	}
	xlrow_ := xlrowFormCallback.xlrow
	_ = xlrow_

	// get the formGroup
	formGroup := xlrowFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				xlrowFormCallback.playground.stageOfInterest,
				xlrowFormCallback.playground.backRepoOfInterest,
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
				for _xlsheet := range *models.GetGongstructInstancesSet[models.XLSheet](xlrowFormCallback.playground.stageOfInterest) {

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

	xlrowFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.XLRow](
		xlrowFormCallback.playground,
	)
	xlrowFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if xlrowFormCallback.CreationMode {
		xlrowFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLRowFormCallback(
				nil,
				xlrowFormCallback.playground,
			),
		}).Stage(xlrowFormCallback.playground.formStage)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, newFormGroup, xlrowFormCallback.playground)
		xlrowFormCallback.playground.formStage.Commit()
	}

	fillUpTree(xlrowFormCallback.playground)
}
func __gong__New__XLSheetFormCallback(
	xlsheet *models.XLSheet,
	playground *Playground,
) (xlsheetFormCallback *XLSheetFormCallback) {
	xlsheetFormCallback = new(XLSheetFormCallback)
	xlsheetFormCallback.playground = playground
	xlsheetFormCallback.xlsheet = xlsheet

	xlsheetFormCallback.CreationMode = (xlsheet == nil)

	return
}

type XLSheetFormCallback struct {
	xlsheet *models.XLSheet

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (xlsheetFormCallback *XLSheetFormCallback) OnSave() {

	log.Println("XLSheetFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	xlsheetFormCallback.playground.formStage.Checkout()

	if xlsheetFormCallback.xlsheet == nil {
		xlsheetFormCallback.xlsheet = new(models.XLSheet).Stage(xlsheetFormCallback.playground.stageOfInterest)
	}
	xlsheet_ := xlsheetFormCallback.xlsheet
	_ = xlsheet_

	// get the formGroup
	formGroup := xlsheetFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				xlsheetFormCallback.playground.stageOfInterest,
				xlsheetFormCallback.playground.backRepoOfInterest,
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
				for _xlfile := range *models.GetGongstructInstancesSet[models.XLFile](xlsheetFormCallback.playground.stageOfInterest) {

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

	xlsheetFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.XLSheet](
		xlsheetFormCallback.playground,
	)
	xlsheetFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if xlsheetFormCallback.CreationMode {
		xlsheetFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__XLSheetFormCallback(
				nil,
				xlsheetFormCallback.playground,
			),
		}).Stage(xlsheetFormCallback.playground.formStage)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, newFormGroup, xlsheetFormCallback.playground)
		xlsheetFormCallback.playground.formStage.Commit()
	}

	fillUpTree(xlsheetFormCallback.playground)
}
