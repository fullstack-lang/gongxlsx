// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *DisplaySelection:
		// insertion point per field

	case *XLCell:
		// insertion point per field

	case *XLFile:
		// insertion point per field
		if fieldName == "Sheets" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*XLFile) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*XLFile)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Sheets).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Sheets = _inferedTypeInstance.Sheets[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Sheets =
								append(_inferedTypeInstance.Sheets, any(fieldInstance).(*XLSheet))
						}
					}
				}
			}
		}

	case *XLRow:
		// insertion point per field
		if fieldName == "Cells" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*XLRow) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*XLRow)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Cells).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Cells = _inferedTypeInstance.Cells[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Cells =
								append(_inferedTypeInstance.Cells, any(fieldInstance).(*XLCell))
						}
					}
				}
			}
		}

	case *XLSheet:
		// insertion point per field
		if fieldName == "Rows" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*XLSheet) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*XLSheet)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rows).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rows = _inferedTypeInstance.Rows[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rows =
								append(_inferedTypeInstance.Rows, any(fieldInstance).(*XLRow))
						}
					}
				}
			}
		}
		if fieldName == "SheetCells" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*XLSheet) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*XLSheet)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SheetCells).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SheetCells = _inferedTypeInstance.SheetCells[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SheetCells =
								append(_inferedTypeInstance.SheetCells, any(fieldInstance).(*XLCell))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct DisplaySelection
	// insertion point per field

	// Compute reverse map for named struct XLCell
	// insertion point per field

	// Compute reverse map for named struct XLFile
	// insertion point per field
	clear(stage.XLFile_Sheets_reverseMap)
	stage.XLFile_Sheets_reverseMap = make(map[*XLSheet]*XLFile)
	for xlfile := range stage.XLFiles {
		_ = xlfile
		for _, _xlsheet := range xlfile.Sheets {
			stage.XLFile_Sheets_reverseMap[_xlsheet] = xlfile
		}
	}

	// Compute reverse map for named struct XLRow
	// insertion point per field
	clear(stage.XLRow_Cells_reverseMap)
	stage.XLRow_Cells_reverseMap = make(map[*XLCell]*XLRow)
	for xlrow := range stage.XLRows {
		_ = xlrow
		for _, _xlcell := range xlrow.Cells {
			stage.XLRow_Cells_reverseMap[_xlcell] = xlrow
		}
	}

	// Compute reverse map for named struct XLSheet
	// insertion point per field
	clear(stage.XLSheet_Rows_reverseMap)
	stage.XLSheet_Rows_reverseMap = make(map[*XLRow]*XLSheet)
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		for _, _xlrow := range xlsheet.Rows {
			stage.XLSheet_Rows_reverseMap[_xlrow] = xlsheet
		}
	}
	clear(stage.XLSheet_SheetCells_reverseMap)
	stage.XLSheet_SheetCells_reverseMap = make(map[*XLCell]*XLSheet)
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		for _, _xlcell := range xlsheet.SheetCells {
			stage.XLSheet_SheetCells_reverseMap[_xlcell] = xlsheet
		}
	}

}
