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
var dummy_DisplaySelection_sql sql.NullBool
var dummy_DisplaySelection_time time.Duration
var dummy_DisplaySelection_sort sort.Float64Slice

// DisplaySelectionAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model displayselectionAPI
type DisplaySelectionAPI struct {
	gorm.Model

	models.DisplaySelection

	// encoding of pointers
	DisplaySelectionPointersEnconding
}

// DisplaySelectionPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DisplaySelectionPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field XLFile is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	XLFileID sql.NullInt64

	// field XLSheet is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	XLSheetID sql.NullInt64
}

// DisplaySelectionDB describes a displayselection in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model displayselectionDB
type DisplaySelectionDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field displayselectionDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString
	// encoding of pointers
	DisplaySelectionPointersEnconding
}

// DisplaySelectionDBs arrays displayselectionDBs
// swagger:response displayselectionDBsResponse
type DisplaySelectionDBs []DisplaySelectionDB

// DisplaySelectionDBResponse provides response
// swagger:response displayselectionDBResponse
type DisplaySelectionDBResponse struct {
	DisplaySelectionDB
}

// DisplaySelectionWOP is a DisplaySelection without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DisplaySelectionWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var DisplaySelection_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoDisplaySelectionStruct struct {
	// stores DisplaySelectionDB according to their gorm ID
	Map_DisplaySelectionDBID_DisplaySelectionDB *map[uint]*DisplaySelectionDB

	// stores DisplaySelectionDB ID according to DisplaySelection address
	Map_DisplaySelectionPtr_DisplaySelectionDBID *map[*models.DisplaySelection]uint

	// stores DisplaySelection according to their gorm ID
	Map_DisplaySelectionDBID_DisplaySelectionPtr *map[uint]*models.DisplaySelection

	db *gorm.DB
}

func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) GetDB() *gorm.DB {
	return backRepoDisplaySelection.db
}

// GetDisplaySelectionDBFromDisplaySelectionPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) GetDisplaySelectionDBFromDisplaySelectionPtr(displayselection *models.DisplaySelection) (displayselectionDB *DisplaySelectionDB) {
	id := (*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]
	displayselectionDB = (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[id]
	return
}

// BackRepoDisplaySelection.Init set up the BackRepo of the DisplaySelection
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) Init(db *gorm.DB) (Error error) {

	if backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr != nil {
		err := errors.New("In Init, backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr should be nil")
		return err
	}

	if backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB != nil {
		err := errors.New("In Init, backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB should be nil")
		return err
	}

	if backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID != nil {
		err := errors.New("In Init, backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.DisplaySelection, 0)
	backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr = &tmp

	tmpDB := make(map[uint]*DisplaySelectionDB, 0)
	backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB = &tmpDB

	tmpID := make(map[*models.DisplaySelection]uint, 0)
	backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID = &tmpID

	backRepoDisplaySelection.db = db
	return
}

// BackRepoDisplaySelection.CommitPhaseOne commits all staged instances of DisplaySelection to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for displayselection := range stage.DisplaySelections {
		backRepoDisplaySelection.CommitPhaseOneInstance(displayselection)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, displayselection := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr {
		if _, ok := stage.DisplaySelections[displayselection]; !ok {
			backRepoDisplaySelection.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDisplaySelection.CommitDeleteInstance commits deletion of DisplaySelection to the BackRepo
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CommitDeleteInstance(id uint) (Error error) {

	displayselection := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[id]

	// displayselection is not staged anymore, remove displayselectionDB
	displayselectionDB := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[id]
	query := backRepoDisplaySelection.db.Unscoped().Delete(&displayselectionDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID), displayselection)
	delete((*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr), id)
	delete((*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB), id)

	return
}

// BackRepoDisplaySelection.CommitPhaseOneInstance commits displayselection staged instances of DisplaySelection to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CommitPhaseOneInstance(displayselection *models.DisplaySelection) (Error error) {

	// check if the displayselection is not commited yet
	if _, ok := (*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]; ok {
		return
	}

	// initiate displayselection
	var displayselectionDB DisplaySelectionDB
	displayselectionDB.CopyBasicFieldsFromDisplaySelection(displayselection)

	query := backRepoDisplaySelection.db.Create(&displayselectionDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection] = displayselectionDB.ID
	(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[displayselectionDB.ID] = displayselection
	(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[displayselectionDB.ID] = &displayselectionDB

	return
}

// BackRepoDisplaySelection.CommitPhaseTwo commits all staged instances of DisplaySelection to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, displayselection := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr {
		backRepoDisplaySelection.CommitPhaseTwoInstance(backRepo, idx, displayselection)
	}

	return
}

// BackRepoDisplaySelection.CommitPhaseTwoInstance commits {{structname }} of models.DisplaySelection to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, displayselection *models.DisplaySelection) (Error error) {

	// fetch matching displayselectionDB
	if displayselectionDB, ok := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[idx]; ok {

		displayselectionDB.CopyBasicFieldsFromDisplaySelection(displayselection)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value displayselection.XLFile translates to updating the displayselection.XLFileID
		displayselectionDB.XLFileID.Valid = true // allow for a 0 value (nil association)
		if displayselection.XLFile != nil {
			if XLFileId, ok := (*backRepo.BackRepoXLFile.Map_XLFilePtr_XLFileDBID)[displayselection.XLFile]; ok {
				displayselectionDB.XLFileID.Int64 = int64(XLFileId)
				displayselectionDB.XLFileID.Valid = true
			}
		}

		// commit pointer value displayselection.XLSheet translates to updating the displayselection.XLSheetID
		displayselectionDB.XLSheetID.Valid = true // allow for a 0 value (nil association)
		if displayselection.XLSheet != nil {
			if XLSheetId, ok := (*backRepo.BackRepoXLSheet.Map_XLSheetPtr_XLSheetDBID)[displayselection.XLSheet]; ok {
				displayselectionDB.XLSheetID.Int64 = int64(XLSheetId)
				displayselectionDB.XLSheetID.Valid = true
			}
		}

		query := backRepoDisplaySelection.db.Save(&displayselectionDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DisplaySelection intance %s", displayselection.Name))
		return err
	}

	return
}

// BackRepoDisplaySelection.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CheckoutPhaseOne() (Error error) {

	displayselectionDBArray := make([]DisplaySelectionDB, 0)
	query := backRepoDisplaySelection.db.Find(&displayselectionDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	displayselectionInstancesToBeRemovedFromTheStage := make(map[*models.DisplaySelection]struct{})
	for key, value := range models.Stage.DisplaySelections {
		displayselectionInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, displayselectionDB := range displayselectionDBArray {
		backRepoDisplaySelection.CheckoutPhaseOneInstance(&displayselectionDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		displayselection, ok := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[displayselectionDB.ID]
		if ok {
			delete(displayselectionInstancesToBeRemovedFromTheStage, displayselection)
		}
	}

	// remove from stage and back repo's 3 maps all displayselections that are not in the checkout
	for displayselection := range displayselectionInstancesToBeRemovedFromTheStage {
		displayselection.Unstage()

		// remove instance from the back repo 3 maps
		displayselectionID := (*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]
		delete((*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID), displayselection)
		delete((*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB), displayselectionID)
		delete((*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr), displayselectionID)
	}

	return
}

// CheckoutPhaseOneInstance takes a displayselectionDB that has been found in the DB, updates the backRepo and stages the
// models version of the displayselectionDB
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CheckoutPhaseOneInstance(displayselectionDB *DisplaySelectionDB) (Error error) {

	displayselection, ok := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[displayselectionDB.ID]
	if !ok {
		displayselection = new(models.DisplaySelection)

		(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[displayselectionDB.ID] = displayselection
		(*backRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection] = displayselectionDB.ID

		// append model store with the new element
		displayselection.Name = displayselectionDB.Name_Data.String
		displayselection.Stage()
	}
	displayselectionDB.CopyBasicFieldsToDisplaySelection(displayselection)

	// preserve pointer to displayselectionDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DisplaySelectionDBID_DisplaySelectionDB)[displayselectionDB hold variable pointers
	displayselectionDB_Data := *displayselectionDB
	preservedPtrToDisplaySelection := &displayselectionDB_Data
	(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[displayselectionDB.ID] = preservedPtrToDisplaySelection

	return
}

// BackRepoDisplaySelection.CheckoutPhaseTwo Checkouts all staged instances of DisplaySelection to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, displayselectionDB := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB {
		backRepoDisplaySelection.CheckoutPhaseTwoInstance(backRepo, displayselectionDB)
	}
	return
}

// BackRepoDisplaySelection.CheckoutPhaseTwoInstance Checkouts staged instances of DisplaySelection to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, displayselectionDB *DisplaySelectionDB) (Error error) {

	displayselection := (*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionPtr)[displayselectionDB.ID]
	_ = displayselection // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// XLFile field
	if displayselectionDB.XLFileID.Int64 != 0 {
		displayselection.XLFile = (*backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr)[uint(displayselectionDB.XLFileID.Int64)]
	}
	// XLSheet field
	if displayselectionDB.XLSheetID.Int64 != 0 {
		displayselection.XLSheet = (*backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr)[uint(displayselectionDB.XLSheetID.Int64)]
	}
	return
}

// CommitDisplaySelection allows commit of a single displayselection (if already staged)
func (backRepo *BackRepoStruct) CommitDisplaySelection(displayselection *models.DisplaySelection) {
	backRepo.BackRepoDisplaySelection.CommitPhaseOneInstance(displayselection)
	if id, ok := (*backRepo.BackRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]; ok {
		backRepo.BackRepoDisplaySelection.CommitPhaseTwoInstance(backRepo, id, displayselection)
	}
}

// CommitDisplaySelection allows checkout of a single displayselection (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDisplaySelection(displayselection *models.DisplaySelection) {
	// check if the displayselection is staged
	if _, ok := (*backRepo.BackRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]; ok {

		if id, ok := (*backRepo.BackRepoDisplaySelection.Map_DisplaySelectionPtr_DisplaySelectionDBID)[displayselection]; ok {
			var displayselectionDB DisplaySelectionDB
			displayselectionDB.ID = id

			if err := backRepo.BackRepoDisplaySelection.db.First(&displayselectionDB, id).Error; err != nil {
				log.Panicln("CheckoutDisplaySelection : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDisplaySelection.CheckoutPhaseOneInstance(&displayselectionDB)
			backRepo.BackRepoDisplaySelection.CheckoutPhaseTwoInstance(backRepo, &displayselectionDB)
		}
	}
}

// CopyBasicFieldsFromDisplaySelection
func (displayselectionDB *DisplaySelectionDB) CopyBasicFieldsFromDisplaySelection(displayselection *models.DisplaySelection) {
	// insertion point for fields commit

	displayselectionDB.Name_Data.String = displayselection.Name
	displayselectionDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromDisplaySelectionWOP
func (displayselectionDB *DisplaySelectionDB) CopyBasicFieldsFromDisplaySelectionWOP(displayselection *DisplaySelectionWOP) {
	// insertion point for fields commit

	displayselectionDB.Name_Data.String = displayselection.Name
	displayselectionDB.Name_Data.Valid = true
}

// CopyBasicFieldsToDisplaySelection
func (displayselectionDB *DisplaySelectionDB) CopyBasicFieldsToDisplaySelection(displayselection *models.DisplaySelection) {
	// insertion point for checkout of basic fields (back repo to stage)
	displayselection.Name = displayselectionDB.Name_Data.String
}

// CopyBasicFieldsToDisplaySelectionWOP
func (displayselectionDB *DisplaySelectionDB) CopyBasicFieldsToDisplaySelectionWOP(displayselection *DisplaySelectionWOP) {
	displayselection.ID = int(displayselectionDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	displayselection.Name = displayselectionDB.Name_Data.String
}

// Backup generates a json file from a slice of all DisplaySelectionDB instances in the backrepo
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DisplaySelectionDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DisplaySelectionDB, 0)
	for _, displayselectionDB := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB {
		forBackup = append(forBackup, displayselectionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json DisplaySelection ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json DisplaySelection file", err.Error())
	}
}

// Backup generates a json file from a slice of all DisplaySelectionDB instances in the backrepo
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DisplaySelectionDB, 0)
	for _, displayselectionDB := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB {
		forBackup = append(forBackup, displayselectionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DisplaySelection")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DisplaySelection_Fields, -1)
	for _, displayselectionDB := range forBackup {

		var displayselectionWOP DisplaySelectionWOP
		displayselectionDB.CopyBasicFieldsToDisplaySelectionWOP(&displayselectionWOP)

		row := sh.AddRow()
		row.WriteStruct(&displayselectionWOP, -1)
	}
}

// RestoreXL from the "DisplaySelection" sheet all DisplaySelectionDB instances
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDisplaySelectionid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DisplaySelection"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDisplaySelection.rowVisitorDisplaySelection)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) rowVisitorDisplaySelection(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var displayselectionWOP DisplaySelectionWOP
		row.ReadStruct(&displayselectionWOP)

		// add the unmarshalled struct to the stage
		displayselectionDB := new(DisplaySelectionDB)
		displayselectionDB.CopyBasicFieldsFromDisplaySelectionWOP(&displayselectionWOP)

		displayselectionDB_ID_atBackupTime := displayselectionDB.ID
		displayselectionDB.ID = 0
		query := backRepoDisplaySelection.db.Create(displayselectionDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[displayselectionDB.ID] = displayselectionDB
		BackRepoDisplaySelectionid_atBckpTime_newID[displayselectionDB_ID_atBackupTime] = displayselectionDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DisplaySelectionDB.json" in dirPath that stores an array
// of DisplaySelectionDB and stores it in the database
// the map BackRepoDisplaySelectionid_atBckpTime_newID is updated accordingly
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDisplaySelectionid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DisplaySelectionDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json DisplaySelection file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DisplaySelectionDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DisplaySelectionDBID_DisplaySelectionDB
	for _, displayselectionDB := range forRestore {

		displayselectionDB_ID_atBackupTime := displayselectionDB.ID
		displayselectionDB.ID = 0
		query := backRepoDisplaySelection.db.Create(displayselectionDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB)[displayselectionDB.ID] = displayselectionDB
		BackRepoDisplaySelectionid_atBckpTime_newID[displayselectionDB_ID_atBackupTime] = displayselectionDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json DisplaySelection file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DisplaySelection>id_atBckpTime_newID
// to compute new index
func (backRepoDisplaySelection *BackRepoDisplaySelectionStruct) RestorePhaseTwo() {

	for _, displayselectionDB := range *backRepoDisplaySelection.Map_DisplaySelectionDBID_DisplaySelectionDB {

		// next line of code is to avert unused variable compilation error
		_ = displayselectionDB

		// insertion point for reindexing pointers encoding
		// reindexing XLFile field
		if displayselectionDB.XLFileID.Int64 != 0 {
			displayselectionDB.XLFileID.Int64 = int64(BackRepoXLFileid_atBckpTime_newID[uint(displayselectionDB.XLFileID.Int64)])
			displayselectionDB.XLFileID.Valid = true
		}

		// reindexing XLSheet field
		if displayselectionDB.XLSheetID.Int64 != 0 {
			displayselectionDB.XLSheetID.Int64 = int64(BackRepoXLSheetid_atBckpTime_newID[uint(displayselectionDB.XLSheetID.Int64)])
			displayselectionDB.XLSheetID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoDisplaySelection.db.Model(displayselectionDB).Updates(*displayselectionDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDisplaySelectionid_atBckpTime_newID map[uint]uint