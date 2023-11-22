package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongxlsx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.DisplaySelection": &(ref_models.DisplaySelection{}),

	"ref_models.DisplaySelection.Name": (ref_models.DisplaySelection{}).Name,

	"ref_models.DisplaySelection.XLFile": (ref_models.DisplaySelection{}).XLFile,

	"ref_models.DisplaySelection.XLSheet": (ref_models.DisplaySelection{}).XLSheet,

	"ref_models.XLCell": &(ref_models.XLCell{}),

	"ref_models.XLCell.Name": (ref_models.XLCell{}).Name,

	"ref_models.XLCell.X": (ref_models.XLCell{}).X,

	"ref_models.XLCell.Y": (ref_models.XLCell{}).Y,

	"ref_models.XLFile": &(ref_models.XLFile{}),

	"ref_models.XLFile.Name": (ref_models.XLFile{}).Name,

	"ref_models.XLFile.NbSheets": (ref_models.XLFile{}).NbSheets,

	"ref_models.XLFile.Sheets": (ref_models.XLFile{}).Sheets,

	"ref_models.XLRow": &(ref_models.XLRow{}),

	"ref_models.XLRow.Cells": (ref_models.XLRow{}).Cells,

	"ref_models.XLRow.Name": (ref_models.XLRow{}).Name,

	"ref_models.XLRow.RowIndex": (ref_models.XLRow{}).RowIndex,

	"ref_models.XLSheet": &(ref_models.XLSheet{}),

	"ref_models.XLSheet.MaxCol": (ref_models.XLSheet{}).MaxCol,

	"ref_models.XLSheet.MaxRow": (ref_models.XLSheet{}).MaxRow,

	"ref_models.XLSheet.Name": (ref_models.XLSheet{}).Name,

	"ref_models.XLSheet.NbRows": (ref_models.XLSheet{}).NbRows,

	"ref_models.XLSheet.Rows": (ref_models.XLSheet{}).Rows,

	"ref_models.XLSheet.SheetCells": (ref_models.XLSheet{}).SheetCells,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_MaxCol := (&models.Field{Name: `MaxCol`}).Stage(stage)
	__Field__000001_MaxRow := (&models.Field{Name: `MaxRow`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000007_NbRows := (&models.Field{Name: `NbRows`}).Stage(stage)
	__Field__000008_NbSheets := (&models.Field{Name: `NbSheets`}).Stage(stage)
	__Field__000009_RowIndex := (&models.Field{Name: `RowIndex`}).Stage(stage)
	__Field__000010_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000011_Y := (&models.Field{Name: `Y`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_DisplaySelection := (&models.GongStructShape{Name: `NewDiagram-DisplaySelection`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_XLCell := (&models.GongStructShape{Name: `NewDiagram-XLCell`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_XLFile := (&models.GongStructShape{Name: `NewDiagram-XLFile`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_XLRow := (&models.GongStructShape{Name: `NewDiagram-XLRow`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_XLSheet := (&models.GongStructShape{Name: `NewDiagram-XLSheet`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Cells := (&models.Link{Name: `Cells`}).Stage(stage)
	__Link__000001_Rows := (&models.Link{Name: `Rows`}).Stage(stage)
	__Link__000002_SheetCells := (&models.Link{Name: `SheetCells`}).Stage(stage)
	__Link__000003_Sheets := (&models.Link{Name: `Sheets`}).Stage(stage)
	__Link__000004_XLFile := (&models.Link{Name: `XLFile`}).Stage(stage)
	__Link__000005_XLSheet := (&models.Link{Name: `XLSheet`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_DisplaySelection := (&models.Position{Name: `Pos-NewDiagram-DisplaySelection`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_XLCell := (&models.Position{Name: `Pos-NewDiagram-XLCell`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_XLFile := (&models.Position{Name: `Pos-NewDiagram-XLFile`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_XLRow := (&models.Position{Name: `Pos-NewDiagram-XLRow`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_XLSheet := (&models.Position{Name: `Pos-NewDiagram-XLSheet`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLFile := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-DisplaySelection and NewDiagram-XLFile`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLSheet := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-DisplaySelection and NewDiagram-XLSheet`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLFile_and_NewDiagram_XLSheet := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-XLFile and NewDiagram-XLSheet`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLRow_and_NewDiagram_XLCell := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-XLRow and NewDiagram-XLCell`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLCell := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-XLSheet and NewDiagram-XLCell`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLRow := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-XLSheet and NewDiagram-XLRow`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_MaxCol.Name = `MaxCol`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.MaxCol]
	__Field__000000_MaxCol.Identifier = `ref_models.XLSheet.MaxCol`
	__Field__000000_MaxCol.FieldTypeAsString = ``
	__Field__000000_MaxCol.Structname = `XLSheet`
	__Field__000000_MaxCol.Fieldtypename = `int`

	// Field values setup
	__Field__000001_MaxRow.Name = `MaxRow`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.MaxRow]
	__Field__000001_MaxRow.Identifier = `ref_models.XLSheet.MaxRow`
	__Field__000001_MaxRow.FieldTypeAsString = ``
	__Field__000001_MaxRow.Structname = `XLSheet`
	__Field__000001_MaxRow.Fieldtypename = `int`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLRow.Name]
	__Field__000002_Name.Identifier = `ref_models.XLRow.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `XLRow`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.Name]
	__Field__000003_Name.Identifier = `ref_models.XLSheet.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `XLSheet`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplaySelection.Name]
	__Field__000004_Name.Identifier = `ref_models.DisplaySelection.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `DisplaySelection`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLFile.Name]
	__Field__000005_Name.Identifier = `ref_models.XLFile.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `XLFile`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell.Name]
	__Field__000006_Name.Identifier = `ref_models.XLCell.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `XLCell`
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_NbRows.Name = `NbRows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.NbRows]
	__Field__000007_NbRows.Identifier = `ref_models.XLSheet.NbRows`
	__Field__000007_NbRows.FieldTypeAsString = ``
	__Field__000007_NbRows.Structname = `XLSheet`
	__Field__000007_NbRows.Fieldtypename = `int`

	// Field values setup
	__Field__000008_NbSheets.Name = `NbSheets`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLFile.NbSheets]
	__Field__000008_NbSheets.Identifier = `ref_models.XLFile.NbSheets`
	__Field__000008_NbSheets.FieldTypeAsString = ``
	__Field__000008_NbSheets.Structname = `XLFile`
	__Field__000008_NbSheets.Fieldtypename = `int`

	// Field values setup
	__Field__000009_RowIndex.Name = `RowIndex`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLRow.RowIndex]
	__Field__000009_RowIndex.Identifier = `ref_models.XLRow.RowIndex`
	__Field__000009_RowIndex.FieldTypeAsString = ``
	__Field__000009_RowIndex.Structname = `XLRow`
	__Field__000009_RowIndex.Fieldtypename = `int`

	// Field values setup
	__Field__000010_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell.X]
	__Field__000010_X.Identifier = `ref_models.XLCell.X`
	__Field__000010_X.FieldTypeAsString = ``
	__Field__000010_X.Structname = `XLCell`
	__Field__000010_X.Fieldtypename = `int`

	// Field values setup
	__Field__000011_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell.Y]
	__Field__000011_Y.Identifier = `ref_models.XLCell.Y`
	__Field__000011_Y.FieldTypeAsString = ``
	__Field__000011_Y.Structname = `XLCell`
	__Field__000011_Y.Fieldtypename = `int`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_DisplaySelection.Name = `NewDiagram-DisplaySelection`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplaySelection]
	__GongStructShape__000000_NewDiagram_DisplaySelection.Identifier = `ref_models.DisplaySelection`
	__GongStructShape__000000_NewDiagram_DisplaySelection.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_DisplaySelection.NbInstances = 0
	__GongStructShape__000000_NewDiagram_DisplaySelection.Width = 240.000000
	__GongStructShape__000000_NewDiagram_DisplaySelection.Height = 78.000000
	__GongStructShape__000000_NewDiagram_DisplaySelection.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_XLCell.Name = `NewDiagram-XLCell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell]
	__GongStructShape__000001_NewDiagram_XLCell.Identifier = `ref_models.XLCell`
	__GongStructShape__000001_NewDiagram_XLCell.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_XLCell.NbInstances = 0
	__GongStructShape__000001_NewDiagram_XLCell.Width = 240.000000
	__GongStructShape__000001_NewDiagram_XLCell.Height = 108.000000
	__GongStructShape__000001_NewDiagram_XLCell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_XLFile.Name = `NewDiagram-XLFile`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLFile]
	__GongStructShape__000002_NewDiagram_XLFile.Identifier = `ref_models.XLFile`
	__GongStructShape__000002_NewDiagram_XLFile.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_XLFile.NbInstances = 0
	__GongStructShape__000002_NewDiagram_XLFile.Width = 240.000000
	__GongStructShape__000002_NewDiagram_XLFile.Height = 93.000000
	__GongStructShape__000002_NewDiagram_XLFile.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_XLRow.Name = `NewDiagram-XLRow`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLRow]
	__GongStructShape__000003_NewDiagram_XLRow.Identifier = `ref_models.XLRow`
	__GongStructShape__000003_NewDiagram_XLRow.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_XLRow.NbInstances = 0
	__GongStructShape__000003_NewDiagram_XLRow.Width = 240.000000
	__GongStructShape__000003_NewDiagram_XLRow.Height = 93.000000
	__GongStructShape__000003_NewDiagram_XLRow.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_XLSheet.Name = `NewDiagram-XLSheet`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet]
	__GongStructShape__000004_NewDiagram_XLSheet.Identifier = `ref_models.XLSheet`
	__GongStructShape__000004_NewDiagram_XLSheet.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_XLSheet.NbInstances = 0
	__GongStructShape__000004_NewDiagram_XLSheet.Width = 240.000000
	__GongStructShape__000004_NewDiagram_XLSheet.Height = 123.000000
	__GongStructShape__000004_NewDiagram_XLSheet.IsSelected = false

	// Link values setup
	__Link__000000_Cells.Name = `Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLRow.Cells]
	__Link__000000_Cells.Identifier = `ref_models.XLRow.Cells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell]
	__Link__000000_Cells.Fieldtypename = `ref_models.XLCell`
	__Link__000000_Cells.FieldOffsetX = 26.000000
	__Link__000000_Cells.FieldOffsetY = -20.000000
	__Link__000000_Cells.TargetMultiplicity = models.MANY
	__Link__000000_Cells.TargetMultiplicityOffsetX = 24.000000
	__Link__000000_Cells.TargetMultiplicityOffsetY = 38.000000
	__Link__000000_Cells.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Cells.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Cells.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Cells.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Cells.StartRatio = 0.788611
	__Link__000000_Cells.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Cells.EndRatio = 0.537037
	__Link__000000_Cells.CornerOffsetRatio = 1.010753

	// Link values setup
	__Link__000001_Rows.Name = `Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.Rows]
	__Link__000001_Rows.Identifier = `ref_models.XLSheet.Rows`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLRow]
	__Link__000001_Rows.Fieldtypename = `ref_models.XLRow`
	__Link__000001_Rows.FieldOffsetX = -62.000000
	__Link__000001_Rows.FieldOffsetY = -19.000000
	__Link__000001_Rows.TargetMultiplicity = models.MANY
	__Link__000001_Rows.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Rows.TargetMultiplicityOffsetY = 27.000000
	__Link__000001_Rows.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Rows.SourceMultiplicityOffsetX = -44.000000
	__Link__000001_Rows.SourceMultiplicityOffsetY = 23.000000
	__Link__000001_Rows.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_Rows.StartRatio = 0.738611
	__Link__000001_Rows.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Rows.EndRatio = 0.397849
	__Link__000001_Rows.CornerOffsetRatio = 1.008130

	// Link values setup
	__Link__000002_SheetCells.Name = `SheetCells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet.SheetCells]
	__Link__000002_SheetCells.Identifier = `ref_models.XLSheet.SheetCells`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLCell]
	__Link__000002_SheetCells.Fieldtypename = `ref_models.XLCell`
	__Link__000002_SheetCells.FieldOffsetX = 18.000000
	__Link__000002_SheetCells.FieldOffsetY = -18.000000
	__Link__000002_SheetCells.TargetMultiplicity = models.MANY
	__Link__000002_SheetCells.TargetMultiplicityOffsetX = -36.000000
	__Link__000002_SheetCells.TargetMultiplicityOffsetY = -10.000000
	__Link__000002_SheetCells.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_SheetCells.SourceMultiplicityOffsetX = -42.000000
	__Link__000002_SheetCells.SourceMultiplicityOffsetY = 27.000000
	__Link__000002_SheetCells.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_SheetCells.StartRatio = 0.230278
	__Link__000002_SheetCells.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_SheetCells.EndRatio = 0.280278
	__Link__000002_SheetCells.CornerOffsetRatio = 1.642276

	// Link values setup
	__Link__000003_Sheets.Name = `Sheets`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLFile.Sheets]
	__Link__000003_Sheets.Identifier = `ref_models.XLFile.Sheets`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet]
	__Link__000003_Sheets.Fieldtypename = `ref_models.XLSheet`
	__Link__000003_Sheets.FieldOffsetX = -83.000000
	__Link__000003_Sheets.FieldOffsetY = -17.000000
	__Link__000003_Sheets.TargetMultiplicity = models.MANY
	__Link__000003_Sheets.TargetMultiplicityOffsetX = 23.000000
	__Link__000003_Sheets.TargetMultiplicityOffsetY = -14.000000
	__Link__000003_Sheets.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Sheets.SourceMultiplicityOffsetX = 14.000000
	__Link__000003_Sheets.SourceMultiplicityOffsetY = 36.000000
	__Link__000003_Sheets.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Sheets.StartRatio = 0.713611
	__Link__000003_Sheets.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Sheets.EndRatio = 0.696944
	__Link__000003_Sheets.CornerOffsetRatio = 1.537634

	// Link values setup
	__Link__000004_XLFile.Name = `XLFile`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplaySelection.XLFile]
	__Link__000004_XLFile.Identifier = `ref_models.DisplaySelection.XLFile`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLFile]
	__Link__000004_XLFile.Fieldtypename = `ref_models.XLFile`
	__Link__000004_XLFile.FieldOffsetX = 8.000000
	__Link__000004_XLFile.FieldOffsetY = -14.000000
	__Link__000004_XLFile.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_XLFile.TargetMultiplicityOffsetX = 22.000000
	__Link__000004_XLFile.TargetMultiplicityOffsetY = 28.000000
	__Link__000004_XLFile.SourceMultiplicity = models.MANY
	__Link__000004_XLFile.SourceMultiplicityOffsetX = -29.000000
	__Link__000004_XLFile.SourceMultiplicityOffsetY = 24.000000
	__Link__000004_XLFile.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_XLFile.StartRatio = 0.500000
	__Link__000004_XLFile.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_XLFile.EndRatio = 0.602151
	__Link__000004_XLFile.CornerOffsetRatio = -0.753056

	// Link values setup
	__Link__000005_XLSheet.Name = `XLSheet`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DisplaySelection.XLSheet]
	__Link__000005_XLSheet.Identifier = `ref_models.DisplaySelection.XLSheet`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.XLSheet]
	__Link__000005_XLSheet.Fieldtypename = `ref_models.XLSheet`
	__Link__000005_XLSheet.FieldOffsetX = 11.000000
	__Link__000005_XLSheet.FieldOffsetY = -21.000000
	__Link__000005_XLSheet.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_XLSheet.TargetMultiplicityOffsetX = 22.000000
	__Link__000005_XLSheet.TargetMultiplicityOffsetY = 21.000000
	__Link__000005_XLSheet.SourceMultiplicity = models.MANY
	__Link__000005_XLSheet.SourceMultiplicityOffsetX = 20.000000
	__Link__000005_XLSheet.SourceMultiplicityOffsetY = 30.000000
	__Link__000005_XLSheet.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_XLSheet.StartRatio = 0.755278
	__Link__000005_XLSheet.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_XLSheet.EndRatio = 0.414634
	__Link__000005_XLSheet.CornerOffsetRatio = 1.025641

	// Position values setup
	__Position__000000_Pos_NewDiagram_DisplaySelection.X = 610.000000
	__Position__000000_Pos_NewDiagram_DisplaySelection.Y = 30.000000
	__Position__000000_Pos_NewDiagram_DisplaySelection.Name = `Pos-NewDiagram-DisplaySelection`

	// Position values setup
	__Position__000001_Pos_NewDiagram_XLCell.X = 57.000000
	__Position__000001_Pos_NewDiagram_XLCell.Y = 478.000000
	__Position__000001_Pos_NewDiagram_XLCell.Name = `Pos-NewDiagram-XLCell`

	// Position values setup
	__Position__000002_Pos_NewDiagram_XLFile.X = 69.000000
	__Position__000002_Pos_NewDiagram_XLFile.Y = 17.000000
	__Position__000002_Pos_NewDiagram_XLFile.Name = `Pos-NewDiagram-XLFile`

	// Position values setup
	__Position__000003_Pos_NewDiagram_XLRow.X = 595.000000
	__Position__000003_Pos_NewDiagram_XLRow.Y = 368.000000
	__Position__000003_Pos_NewDiagram_XLRow.Name = `Pos-NewDiagram-XLRow`

	// Position values setup
	__Position__000004_Pos_NewDiagram_XLSheet.X = 75.000000
	__Position__000004_Pos_NewDiagram_XLSheet.Y = 223.000000
	__Position__000004_Pos_NewDiagram_XLSheet.Name = `Pos-NewDiagram-XLSheet`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLFile.X = 700.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLFile.Y = 54.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLFile.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-DisplaySelection and NewDiagram-XLFile`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLSheet.X = 700.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLSheet.Y = 164.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLSheet.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-DisplaySelection and NewDiagram-XLSheet`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLFile_and_NewDiagram_XLSheet.X = 430.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLFile_and_NewDiagram_XLSheet.Y = 156.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLFile_and_NewDiagram_XLSheet.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-XLFile and NewDiagram-XLSheet`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLRow_and_NewDiagram_XLCell.X = 686.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLRow_and_NewDiagram_XLCell.Y = 469.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLRow_and_NewDiagram_XLCell.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-XLRow and NewDiagram-XLCell`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLCell.X = 426.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLCell.Y = 412.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLCell.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-XLSheet and NewDiagram-XLCell`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLRow.X = 692.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLRow.Y = 355.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLRow.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-XLSheet and NewDiagram-XLRow`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_XLFile)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_XLCell)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_XLSheet)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_DisplaySelection)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_XLRow)
	__GongStructShape__000000_NewDiagram_DisplaySelection.Position = __Position__000000_Pos_NewDiagram_DisplaySelection
	__GongStructShape__000000_NewDiagram_DisplaySelection.Fields = append(__GongStructShape__000000_NewDiagram_DisplaySelection.Fields, __Field__000004_Name)
	__GongStructShape__000000_NewDiagram_DisplaySelection.Links = append(__GongStructShape__000000_NewDiagram_DisplaySelection.Links, __Link__000004_XLFile)
	__GongStructShape__000000_NewDiagram_DisplaySelection.Links = append(__GongStructShape__000000_NewDiagram_DisplaySelection.Links, __Link__000005_XLSheet)
	__GongStructShape__000001_NewDiagram_XLCell.Position = __Position__000001_Pos_NewDiagram_XLCell
	__GongStructShape__000001_NewDiagram_XLCell.Fields = append(__GongStructShape__000001_NewDiagram_XLCell.Fields, __Field__000006_Name)
	__GongStructShape__000001_NewDiagram_XLCell.Fields = append(__GongStructShape__000001_NewDiagram_XLCell.Fields, __Field__000010_X)
	__GongStructShape__000001_NewDiagram_XLCell.Fields = append(__GongStructShape__000001_NewDiagram_XLCell.Fields, __Field__000011_Y)
	__GongStructShape__000002_NewDiagram_XLFile.Position = __Position__000002_Pos_NewDiagram_XLFile
	__GongStructShape__000002_NewDiagram_XLFile.Fields = append(__GongStructShape__000002_NewDiagram_XLFile.Fields, __Field__000005_Name)
	__GongStructShape__000002_NewDiagram_XLFile.Fields = append(__GongStructShape__000002_NewDiagram_XLFile.Fields, __Field__000008_NbSheets)
	__GongStructShape__000002_NewDiagram_XLFile.Links = append(__GongStructShape__000002_NewDiagram_XLFile.Links, __Link__000003_Sheets)
	__GongStructShape__000003_NewDiagram_XLRow.Position = __Position__000003_Pos_NewDiagram_XLRow
	__GongStructShape__000003_NewDiagram_XLRow.Fields = append(__GongStructShape__000003_NewDiagram_XLRow.Fields, __Field__000002_Name)
	__GongStructShape__000003_NewDiagram_XLRow.Fields = append(__GongStructShape__000003_NewDiagram_XLRow.Fields, __Field__000009_RowIndex)
	__GongStructShape__000003_NewDiagram_XLRow.Links = append(__GongStructShape__000003_NewDiagram_XLRow.Links, __Link__000000_Cells)
	__GongStructShape__000004_NewDiagram_XLSheet.Position = __Position__000004_Pos_NewDiagram_XLSheet
	__GongStructShape__000004_NewDiagram_XLSheet.Fields = append(__GongStructShape__000004_NewDiagram_XLSheet.Fields, __Field__000003_Name)
	__GongStructShape__000004_NewDiagram_XLSheet.Fields = append(__GongStructShape__000004_NewDiagram_XLSheet.Fields, __Field__000001_MaxRow)
	__GongStructShape__000004_NewDiagram_XLSheet.Fields = append(__GongStructShape__000004_NewDiagram_XLSheet.Fields, __Field__000000_MaxCol)
	__GongStructShape__000004_NewDiagram_XLSheet.Fields = append(__GongStructShape__000004_NewDiagram_XLSheet.Fields, __Field__000007_NbRows)
	__GongStructShape__000004_NewDiagram_XLSheet.Links = append(__GongStructShape__000004_NewDiagram_XLSheet.Links, __Link__000001_Rows)
	__GongStructShape__000004_NewDiagram_XLSheet.Links = append(__GongStructShape__000004_NewDiagram_XLSheet.Links, __Link__000002_SheetCells)
	__Link__000000_Cells.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLRow_and_NewDiagram_XLCell
	__Link__000001_Rows.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLRow
	__Link__000002_SheetCells.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLSheet_and_NewDiagram_XLCell
	__Link__000003_Sheets.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_XLFile_and_NewDiagram_XLSheet
	__Link__000004_XLFile.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLFile
	__Link__000005_XLSheet.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_DisplaySelection_and_NewDiagram_XLSheet
}
