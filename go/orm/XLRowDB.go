// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongxlsx/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_XLRow_sql sql.NullBool
var dummy_XLRow_time time.Duration
var dummy_XLRow_sort sort.Float64Slice

// XLRowAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xlrowAPI
type XLRowAPI struct {
	gorm.Model

	models.XLRow

	// encoding of pointers
	XLRowPointersEnconding
}

// XLRowPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type XLRowPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field XLSheet{}.Rows []*XLRow
	XLSheet_RowsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	XLSheet_RowsDBID_Index sql.NullInt64
}

// XLRowDB describes a xlrow in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xlrowDB
type XLRowDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xlrowDB.Name
	Name_Data sql.NullString

	// Declation for basic field xlrowDB.RowIndex
	RowIndex_Data sql.NullInt64
	// encoding of pointers
	XLRowPointersEnconding
}

// XLRowDBs arrays xlrowDBs
// swagger:response xlrowDBsResponse
type XLRowDBs []XLRowDB

// XLRowDBResponse provides response
// swagger:response xlrowDBResponse
type XLRowDBResponse struct {
	XLRowDB
}

// XLRowWOP is a XLRow without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type XLRowWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	RowIndex int `xlsx:"2"`
	// insertion for WOP pointer fields
}

var XLRow_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"RowIndex",
}

type BackRepoXLRowStruct struct {
	// stores XLRowDB according to their gorm ID
	Map_XLRowDBID_XLRowDB *map[uint]*XLRowDB

	// stores XLRowDB ID according to XLRow address
	Map_XLRowPtr_XLRowDBID *map[*models.XLRow]uint

	// stores XLRow according to their gorm ID
	Map_XLRowDBID_XLRowPtr *map[uint]*models.XLRow

	db *gorm.DB
}

func (backRepoXLRow *BackRepoXLRowStruct) GetDB() *gorm.DB {
	return backRepoXLRow.db
}

// GetXLRowDBFromXLRowPtr is a handy function to access the back repo instance from the stage instance
func (backRepoXLRow *BackRepoXLRowStruct) GetXLRowDBFromXLRowPtr(xlrow *models.XLRow) (xlrowDB *XLRowDB) {
	id := (*backRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]
	xlrowDB = (*backRepoXLRow.Map_XLRowDBID_XLRowDB)[id]
	return
}

// BackRepoXLRow.Init set up the BackRepo of the XLRow
func (backRepoXLRow *BackRepoXLRowStruct) Init(db *gorm.DB) (Error error) {

	if backRepoXLRow.Map_XLRowDBID_XLRowPtr != nil {
		err := errors.New("In Init, backRepoXLRow.Map_XLRowDBID_XLRowPtr should be nil")
		return err
	}

	if backRepoXLRow.Map_XLRowDBID_XLRowDB != nil {
		err := errors.New("In Init, backRepoXLRow.Map_XLRowDBID_XLRowDB should be nil")
		return err
	}

	if backRepoXLRow.Map_XLRowPtr_XLRowDBID != nil {
		err := errors.New("In Init, backRepoXLRow.Map_XLRowPtr_XLRowDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.XLRow, 0)
	backRepoXLRow.Map_XLRowDBID_XLRowPtr = &tmp

	tmpDB := make(map[uint]*XLRowDB, 0)
	backRepoXLRow.Map_XLRowDBID_XLRowDB = &tmpDB

	tmpID := make(map[*models.XLRow]uint, 0)
	backRepoXLRow.Map_XLRowPtr_XLRowDBID = &tmpID

	backRepoXLRow.db = db
	return
}

// BackRepoXLRow.CommitPhaseOne commits all staged instances of XLRow to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXLRow *BackRepoXLRowStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xlrow := range stage.XLRows {
		backRepoXLRow.CommitPhaseOneInstance(xlrow)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xlrow := range *backRepoXLRow.Map_XLRowDBID_XLRowPtr {
		if _, ok := stage.XLRows[xlrow]; !ok {
			backRepoXLRow.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXLRow.CommitDeleteInstance commits deletion of XLRow to the BackRepo
func (backRepoXLRow *BackRepoXLRowStruct) CommitDeleteInstance(id uint) (Error error) {

	xlrow := (*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[id]

	// xlrow is not staged anymore, remove xlrowDB
	xlrowDB := (*backRepoXLRow.Map_XLRowDBID_XLRowDB)[id]
	query := backRepoXLRow.db.Unscoped().Delete(&xlrowDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoXLRow.Map_XLRowPtr_XLRowDBID), xlrow)
	delete((*backRepoXLRow.Map_XLRowDBID_XLRowPtr), id)
	delete((*backRepoXLRow.Map_XLRowDBID_XLRowDB), id)

	return
}

// BackRepoXLRow.CommitPhaseOneInstance commits xlrow staged instances of XLRow to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXLRow *BackRepoXLRowStruct) CommitPhaseOneInstance(xlrow *models.XLRow) (Error error) {

	// check if the xlrow is not commited yet
	if _, ok := (*backRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]; ok {
		return
	}

	// initiate xlrow
	var xlrowDB XLRowDB
	xlrowDB.CopyBasicFieldsFromXLRow(xlrow)

	query := backRepoXLRow.db.Create(&xlrowDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow] = xlrowDB.ID
	(*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[xlrowDB.ID] = xlrow
	(*backRepoXLRow.Map_XLRowDBID_XLRowDB)[xlrowDB.ID] = &xlrowDB

	return
}

// BackRepoXLRow.CommitPhaseTwo commits all staged instances of XLRow to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLRow *BackRepoXLRowStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xlrow := range *backRepoXLRow.Map_XLRowDBID_XLRowPtr {
		backRepoXLRow.CommitPhaseTwoInstance(backRepo, idx, xlrow)
	}

	return
}

// BackRepoXLRow.CommitPhaseTwoInstance commits {{structname }} of models.XLRow to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLRow *BackRepoXLRowStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xlrow *models.XLRow) (Error error) {

	// fetch matching xlrowDB
	if xlrowDB, ok := (*backRepoXLRow.Map_XLRowDBID_XLRowDB)[idx]; ok {

		xlrowDB.CopyBasicFieldsFromXLRow(xlrow)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers xlrow.Cells into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, xlcellAssocEnd := range xlrow.Cells {

			// get the back repo instance at the association end
			xlcellAssocEnd_DB :=
				backRepo.BackRepoXLCell.GetXLCellDBFromXLCellPtr(xlcellAssocEnd)

			// encode reverse pointer in the association end back repo instance
			xlcellAssocEnd_DB.XLRow_CellsDBID.Int64 = int64(xlrowDB.ID)
			xlcellAssocEnd_DB.XLRow_CellsDBID.Valid = true
			xlcellAssocEnd_DB.XLRow_CellsDBID_Index.Int64 = int64(idx)
			xlcellAssocEnd_DB.XLRow_CellsDBID_Index.Valid = true
			if q := backRepoXLRow.db.Save(xlcellAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoXLRow.db.Save(&xlrowDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown XLRow intance %s", xlrow.Name))
		return err
	}

	return
}

// BackRepoXLRow.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoXLRow *BackRepoXLRowStruct) CheckoutPhaseOne() (Error error) {

	xlrowDBArray := make([]XLRowDB, 0)
	query := backRepoXLRow.db.Find(&xlrowDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xlrowInstancesToBeRemovedFromTheStage := make(map[*models.XLRow]any)
	for key, value := range models.Stage.XLRows {
		xlrowInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xlrowDB := range xlrowDBArray {
		backRepoXLRow.CheckoutPhaseOneInstance(&xlrowDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xlrow, ok := (*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[xlrowDB.ID]
		if ok {
			delete(xlrowInstancesToBeRemovedFromTheStage, xlrow)
		}
	}

	// remove from stage and back repo's 3 maps all xlrows that are not in the checkout
	for xlrow := range xlrowInstancesToBeRemovedFromTheStage {
		xlrow.Unstage()

		// remove instance from the back repo 3 maps
		xlrowID := (*backRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]
		delete((*backRepoXLRow.Map_XLRowPtr_XLRowDBID), xlrow)
		delete((*backRepoXLRow.Map_XLRowDBID_XLRowDB), xlrowID)
		delete((*backRepoXLRow.Map_XLRowDBID_XLRowPtr), xlrowID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xlrowDB that has been found in the DB, updates the backRepo and stages the
// models version of the xlrowDB
func (backRepoXLRow *BackRepoXLRowStruct) CheckoutPhaseOneInstance(xlrowDB *XLRowDB) (Error error) {

	xlrow, ok := (*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[xlrowDB.ID]
	if !ok {
		xlrow = new(models.XLRow)

		(*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[xlrowDB.ID] = xlrow
		(*backRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow] = xlrowDB.ID

		// append model store with the new element
		xlrow.Name = xlrowDB.Name_Data.String
		xlrow.Stage()
	}
	xlrowDB.CopyBasicFieldsToXLRow(xlrow)

	// preserve pointer to xlrowDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_XLRowDBID_XLRowDB)[xlrowDB hold variable pointers
	xlrowDB_Data := *xlrowDB
	preservedPtrToXLRow := &xlrowDB_Data
	(*backRepoXLRow.Map_XLRowDBID_XLRowDB)[xlrowDB.ID] = preservedPtrToXLRow

	return
}

// BackRepoXLRow.CheckoutPhaseTwo Checkouts all staged instances of XLRow to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLRow *BackRepoXLRowStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xlrowDB := range *backRepoXLRow.Map_XLRowDBID_XLRowDB {
		backRepoXLRow.CheckoutPhaseTwoInstance(backRepo, xlrowDB)
	}
	return
}

// BackRepoXLRow.CheckoutPhaseTwoInstance Checkouts staged instances of XLRow to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLRow *BackRepoXLRowStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xlrowDB *XLRowDB) (Error error) {

	xlrow := (*backRepoXLRow.Map_XLRowDBID_XLRowPtr)[xlrowDB.ID]
	_ = xlrow // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem xlrow.Cells in the stage from the encode in the back repo
	// It parses all XLCellDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	xlrow.Cells = xlrow.Cells[:0]
	// 2. loop all instances in the type in the association end
	for _, xlcellDB_AssocEnd := range *backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if xlcellDB_AssocEnd.XLRow_CellsDBID.Int64 == int64(xlrowDB.ID) {
			// 4. fetch the associated instance in the stage
			xlcell_AssocEnd := (*backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellPtr)[xlcellDB_AssocEnd.ID]
			// 5. append it the association slice
			xlrow.Cells = append(xlrow.Cells, xlcell_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(xlrow.Cells, func(i, j int) bool {
		xlcellDB_i_ID := (*backRepo.BackRepoXLCell.Map_XLCellPtr_XLCellDBID)[xlrow.Cells[i]]
		xlcellDB_j_ID := (*backRepo.BackRepoXLCell.Map_XLCellPtr_XLCellDBID)[xlrow.Cells[j]]

		xlcellDB_i := (*backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellDB)[xlcellDB_i_ID]
		xlcellDB_j := (*backRepo.BackRepoXLCell.Map_XLCellDBID_XLCellDB)[xlcellDB_j_ID]

		return xlcellDB_i.XLRow_CellsDBID_Index.Int64 < xlcellDB_j.XLRow_CellsDBID_Index.Int64
	})

	return
}

// CommitXLRow allows commit of a single xlrow (if already staged)
func (backRepo *BackRepoStruct) CommitXLRow(xlrow *models.XLRow) {
	backRepo.BackRepoXLRow.CommitPhaseOneInstance(xlrow)
	if id, ok := (*backRepo.BackRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]; ok {
		backRepo.BackRepoXLRow.CommitPhaseTwoInstance(backRepo, id, xlrow)
	}
}

// CommitXLRow allows checkout of a single xlrow (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXLRow(xlrow *models.XLRow) {
	// check if the xlrow is staged
	if _, ok := (*backRepo.BackRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]; ok {

		if id, ok := (*backRepo.BackRepoXLRow.Map_XLRowPtr_XLRowDBID)[xlrow]; ok {
			var xlrowDB XLRowDB
			xlrowDB.ID = id

			if err := backRepo.BackRepoXLRow.db.First(&xlrowDB, id).Error; err != nil {
				log.Panicln("CheckoutXLRow : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXLRow.CheckoutPhaseOneInstance(&xlrowDB)
			backRepo.BackRepoXLRow.CheckoutPhaseTwoInstance(backRepo, &xlrowDB)
		}
	}
}

// CopyBasicFieldsFromXLRow
func (xlrowDB *XLRowDB) CopyBasicFieldsFromXLRow(xlrow *models.XLRow) {
	// insertion point for fields commit

	xlrowDB.Name_Data.String = xlrow.Name
	xlrowDB.Name_Data.Valid = true

	xlrowDB.RowIndex_Data.Int64 = int64(xlrow.RowIndex)
	xlrowDB.RowIndex_Data.Valid = true
}

// CopyBasicFieldsFromXLRowWOP
func (xlrowDB *XLRowDB) CopyBasicFieldsFromXLRowWOP(xlrow *XLRowWOP) {
	// insertion point for fields commit

	xlrowDB.Name_Data.String = xlrow.Name
	xlrowDB.Name_Data.Valid = true

	xlrowDB.RowIndex_Data.Int64 = int64(xlrow.RowIndex)
	xlrowDB.RowIndex_Data.Valid = true
}

// CopyBasicFieldsToXLRow
func (xlrowDB *XLRowDB) CopyBasicFieldsToXLRow(xlrow *models.XLRow) {
	// insertion point for checkout of basic fields (back repo to stage)
	xlrow.Name = xlrowDB.Name_Data.String
	xlrow.RowIndex = int(xlrowDB.RowIndex_Data.Int64)
}

// CopyBasicFieldsToXLRowWOP
func (xlrowDB *XLRowDB) CopyBasicFieldsToXLRowWOP(xlrow *XLRowWOP) {
	xlrow.ID = int(xlrowDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xlrow.Name = xlrowDB.Name_Data.String
	xlrow.RowIndex = int(xlrowDB.RowIndex_Data.Int64)
}

// Backup generates a json file from a slice of all XLRowDB instances in the backrepo
func (backRepoXLRow *BackRepoXLRowStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "XLRowDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*XLRowDB, 0)
	for _, xlrowDB := range *backRepoXLRow.Map_XLRowDBID_XLRowDB {
		forBackup = append(forBackup, xlrowDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json XLRow ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json XLRow file", err.Error())
	}
}

// Backup generates a json file from a slice of all XLRowDB instances in the backrepo
func (backRepoXLRow *BackRepoXLRowStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*XLRowDB, 0)
	for _, xlrowDB := range *backRepoXLRow.Map_XLRowDBID_XLRowDB {
		forBackup = append(forBackup, xlrowDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("XLRow")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&XLRow_Fields, -1)
	for _, xlrowDB := range forBackup {

		var xlrowWOP XLRowWOP
		xlrowDB.CopyBasicFieldsToXLRowWOP(&xlrowWOP)

		row := sh.AddRow()
		row.WriteStruct(&xlrowWOP, -1)
	}
}

// RestoreXL from the "XLRow" sheet all XLRowDB instances
func (backRepoXLRow *BackRepoXLRowStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXLRowid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["XLRow"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXLRow.rowVisitorXLRow)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoXLRow *BackRepoXLRowStruct) rowVisitorXLRow(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xlrowWOP XLRowWOP
		row.ReadStruct(&xlrowWOP)

		// add the unmarshalled struct to the stage
		xlrowDB := new(XLRowDB)
		xlrowDB.CopyBasicFieldsFromXLRowWOP(&xlrowWOP)

		xlrowDB_ID_atBackupTime := xlrowDB.ID
		xlrowDB.ID = 0
		query := backRepoXLRow.db.Create(xlrowDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoXLRow.Map_XLRowDBID_XLRowDB)[xlrowDB.ID] = xlrowDB
		BackRepoXLRowid_atBckpTime_newID[xlrowDB_ID_atBackupTime] = xlrowDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "XLRowDB.json" in dirPath that stores an array
// of XLRowDB and stores it in the database
// the map BackRepoXLRowid_atBckpTime_newID is updated accordingly
func (backRepoXLRow *BackRepoXLRowStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXLRowid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "XLRowDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json XLRow file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*XLRowDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_XLRowDBID_XLRowDB
	for _, xlrowDB := range forRestore {

		xlrowDB_ID_atBackupTime := xlrowDB.ID
		xlrowDB.ID = 0
		query := backRepoXLRow.db.Create(xlrowDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoXLRow.Map_XLRowDBID_XLRowDB)[xlrowDB.ID] = xlrowDB
		BackRepoXLRowid_atBckpTime_newID[xlrowDB_ID_atBackupTime] = xlrowDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json XLRow file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<XLRow>id_atBckpTime_newID
// to compute new index
func (backRepoXLRow *BackRepoXLRowStruct) RestorePhaseTwo() {

	for _, xlrowDB := range *backRepoXLRow.Map_XLRowDBID_XLRowDB {

		// next line of code is to avert unused variable compilation error
		_ = xlrowDB

		// insertion point for reindexing pointers encoding
		// This reindex xlrow.Rows
		if xlrowDB.XLSheet_RowsDBID.Int64 != 0 {
			xlrowDB.XLSheet_RowsDBID.Int64 =
				int64(BackRepoXLSheetid_atBckpTime_newID[uint(xlrowDB.XLSheet_RowsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoXLRow.db.Model(xlrowDB).Updates(*xlrowDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXLRowid_atBckpTime_newID map[uint]uint
