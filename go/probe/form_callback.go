// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxlsx/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewDisplaySelectionFormCallback(
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
			OnSave: NewDisplaySelectionFormCallback(
				nil,
				displayselectionFormCallback.playground,
			),
		}).Stage(displayselectionFormCallback.playground.formStage)
		displayselection := new(models.DisplaySelection)
		FillUpForm(displayselection, newFormGroup, displayselectionFormCallback.playground)
		displayselectionFormCallback.playground.formStage.Commit()
	}

}
func NewXLCellFormCallback(
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
			OnSave: NewXLCellFormCallback(
				nil,
				xlcellFormCallback.playground,
			),
		}).Stage(xlcellFormCallback.playground.formStage)
		xlcell := new(models.XLCell)
		FillUpForm(xlcell, newFormGroup, xlcellFormCallback.playground)
		xlcellFormCallback.playground.formStage.Commit()
	}

}
func NewXLFileFormCallback(
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
			OnSave: NewXLFileFormCallback(
				nil,
				xlfileFormCallback.playground,
			),
		}).Stage(xlfileFormCallback.playground.formStage)
		xlfile := new(models.XLFile)
		FillUpForm(xlfile, newFormGroup, xlfileFormCallback.playground)
		xlfileFormCallback.playground.formStage.Commit()
	}

}
func NewXLRowFormCallback(
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
			OnSave: NewXLRowFormCallback(
				nil,
				xlrowFormCallback.playground,
			),
		}).Stage(xlrowFormCallback.playground.formStage)
		xlrow := new(models.XLRow)
		FillUpForm(xlrow, newFormGroup, xlrowFormCallback.playground)
		xlrowFormCallback.playground.formStage.Commit()
	}

}
func NewXLSheetFormCallback(
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
			OnSave: NewXLSheetFormCallback(
				nil,
				xlsheetFormCallback.playground,
			),
		}).Stage(xlsheetFormCallback.playground.formStage)
		xlsheet := new(models.XLSheet)
		FillUpForm(xlsheet, newFormGroup, xlsheetFormCallback.playground)
		xlsheetFormCallback.playground.formStage.Commit()
	}

}
