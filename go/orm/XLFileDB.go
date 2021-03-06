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
var dummy_XLFile_sql sql.NullBool
var dummy_XLFile_time time.Duration
var dummy_XLFile_sort sort.Float64Slice

// XLFileAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model xlfileAPI
type XLFileAPI struct {
	gorm.Model

	models.XLFile

	// encoding of pointers
	XLFilePointersEnconding
}

// XLFilePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type XLFilePointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// XLFileDB describes a xlfile in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model xlfileDB
type XLFileDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field xlfileDB.Name
	Name_Data sql.NullString

	// Declation for basic field xlfileDB.NbSheets
	NbSheets_Data sql.NullInt64
	// encoding of pointers
	XLFilePointersEnconding
}

// XLFileDBs arrays xlfileDBs
// swagger:response xlfileDBsResponse
type XLFileDBs []XLFileDB

// XLFileDBResponse provides response
// swagger:response xlfileDBResponse
type XLFileDBResponse struct {
	XLFileDB
}

// XLFileWOP is a XLFile without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type XLFileWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	NbSheets int `xlsx:"2"`
	// insertion for WOP pointer fields
}

var XLFile_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"NbSheets",
}

type BackRepoXLFileStruct struct {
	// stores XLFileDB according to their gorm ID
	Map_XLFileDBID_XLFileDB *map[uint]*XLFileDB

	// stores XLFileDB ID according to XLFile address
	Map_XLFilePtr_XLFileDBID *map[*models.XLFile]uint

	// stores XLFile according to their gorm ID
	Map_XLFileDBID_XLFilePtr *map[uint]*models.XLFile

	db *gorm.DB
}

func (backRepoXLFile *BackRepoXLFileStruct) GetDB() *gorm.DB {
	return backRepoXLFile.db
}

// GetXLFileDBFromXLFilePtr is a handy function to access the back repo instance from the stage instance
func (backRepoXLFile *BackRepoXLFileStruct) GetXLFileDBFromXLFilePtr(xlfile *models.XLFile) (xlfileDB *XLFileDB) {
	id := (*backRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]
	xlfileDB = (*backRepoXLFile.Map_XLFileDBID_XLFileDB)[id]
	return
}

// BackRepoXLFile.Init set up the BackRepo of the XLFile
func (backRepoXLFile *BackRepoXLFileStruct) Init(db *gorm.DB) (Error error) {

	if backRepoXLFile.Map_XLFileDBID_XLFilePtr != nil {
		err := errors.New("In Init, backRepoXLFile.Map_XLFileDBID_XLFilePtr should be nil")
		return err
	}

	if backRepoXLFile.Map_XLFileDBID_XLFileDB != nil {
		err := errors.New("In Init, backRepoXLFile.Map_XLFileDBID_XLFileDB should be nil")
		return err
	}

	if backRepoXLFile.Map_XLFilePtr_XLFileDBID != nil {
		err := errors.New("In Init, backRepoXLFile.Map_XLFilePtr_XLFileDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.XLFile, 0)
	backRepoXLFile.Map_XLFileDBID_XLFilePtr = &tmp

	tmpDB := make(map[uint]*XLFileDB, 0)
	backRepoXLFile.Map_XLFileDBID_XLFileDB = &tmpDB

	tmpID := make(map[*models.XLFile]uint, 0)
	backRepoXLFile.Map_XLFilePtr_XLFileDBID = &tmpID

	backRepoXLFile.db = db
	return
}

// BackRepoXLFile.CommitPhaseOne commits all staged instances of XLFile to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXLFile *BackRepoXLFileStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for xlfile := range stage.XLFiles {
		backRepoXLFile.CommitPhaseOneInstance(xlfile)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, xlfile := range *backRepoXLFile.Map_XLFileDBID_XLFilePtr {
		if _, ok := stage.XLFiles[xlfile]; !ok {
			backRepoXLFile.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoXLFile.CommitDeleteInstance commits deletion of XLFile to the BackRepo
func (backRepoXLFile *BackRepoXLFileStruct) CommitDeleteInstance(id uint) (Error error) {

	xlfile := (*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[id]

	// xlfile is not staged anymore, remove xlfileDB
	xlfileDB := (*backRepoXLFile.Map_XLFileDBID_XLFileDB)[id]
	query := backRepoXLFile.db.Unscoped().Delete(&xlfileDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoXLFile.Map_XLFilePtr_XLFileDBID), xlfile)
	delete((*backRepoXLFile.Map_XLFileDBID_XLFilePtr), id)
	delete((*backRepoXLFile.Map_XLFileDBID_XLFileDB), id)

	return
}

// BackRepoXLFile.CommitPhaseOneInstance commits xlfile staged instances of XLFile to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoXLFile *BackRepoXLFileStruct) CommitPhaseOneInstance(xlfile *models.XLFile) (Error error) {

	// check if the xlfile is not commited yet
	if _, ok := (*backRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]; ok {
		return
	}

	// initiate xlfile
	var xlfileDB XLFileDB
	xlfileDB.CopyBasicFieldsFromXLFile(xlfile)

	query := backRepoXLFile.db.Create(&xlfileDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile] = xlfileDB.ID
	(*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[xlfileDB.ID] = xlfile
	(*backRepoXLFile.Map_XLFileDBID_XLFileDB)[xlfileDB.ID] = &xlfileDB

	return
}

// BackRepoXLFile.CommitPhaseTwo commits all staged instances of XLFile to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLFile *BackRepoXLFileStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, xlfile := range *backRepoXLFile.Map_XLFileDBID_XLFilePtr {
		backRepoXLFile.CommitPhaseTwoInstance(backRepo, idx, xlfile)
	}

	return
}

// BackRepoXLFile.CommitPhaseTwoInstance commits {{structname }} of models.XLFile to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLFile *BackRepoXLFileStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, xlfile *models.XLFile) (Error error) {

	// fetch matching xlfileDB
	if xlfileDB, ok := (*backRepoXLFile.Map_XLFileDBID_XLFileDB)[idx]; ok {

		xlfileDB.CopyBasicFieldsFromXLFile(xlfile)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers xlfile.Sheets into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, xlsheetAssocEnd := range xlfile.Sheets {

			// get the back repo instance at the association end
			xlsheetAssocEnd_DB :=
				backRepo.BackRepoXLSheet.GetXLSheetDBFromXLSheetPtr(xlsheetAssocEnd)

			// encode reverse pointer in the association end back repo instance
			xlsheetAssocEnd_DB.XLFile_SheetsDBID.Int64 = int64(xlfileDB.ID)
			xlsheetAssocEnd_DB.XLFile_SheetsDBID.Valid = true
			xlsheetAssocEnd_DB.XLFile_SheetsDBID_Index.Int64 = int64(idx)
			xlsheetAssocEnd_DB.XLFile_SheetsDBID_Index.Valid = true
			if q := backRepoXLFile.db.Save(xlsheetAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoXLFile.db.Save(&xlfileDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown XLFile intance %s", xlfile.Name))
		return err
	}

	return
}

// BackRepoXLFile.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoXLFile *BackRepoXLFileStruct) CheckoutPhaseOne() (Error error) {

	xlfileDBArray := make([]XLFileDB, 0)
	query := backRepoXLFile.db.Find(&xlfileDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	xlfileInstancesToBeRemovedFromTheStage := make(map[*models.XLFile]any)
	for key, value := range models.Stage.XLFiles {
		xlfileInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, xlfileDB := range xlfileDBArray {
		backRepoXLFile.CheckoutPhaseOneInstance(&xlfileDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		xlfile, ok := (*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[xlfileDB.ID]
		if ok {
			delete(xlfileInstancesToBeRemovedFromTheStage, xlfile)
		}
	}

	// remove from stage and back repo's 3 maps all xlfiles that are not in the checkout
	for xlfile := range xlfileInstancesToBeRemovedFromTheStage {
		xlfile.Unstage()

		// remove instance from the back repo 3 maps
		xlfileID := (*backRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]
		delete((*backRepoXLFile.Map_XLFilePtr_XLFileDBID), xlfile)
		delete((*backRepoXLFile.Map_XLFileDBID_XLFileDB), xlfileID)
		delete((*backRepoXLFile.Map_XLFileDBID_XLFilePtr), xlfileID)
	}

	return
}

// CheckoutPhaseOneInstance takes a xlfileDB that has been found in the DB, updates the backRepo and stages the
// models version of the xlfileDB
func (backRepoXLFile *BackRepoXLFileStruct) CheckoutPhaseOneInstance(xlfileDB *XLFileDB) (Error error) {

	xlfile, ok := (*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[xlfileDB.ID]
	if !ok {
		xlfile = new(models.XLFile)

		(*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[xlfileDB.ID] = xlfile
		(*backRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile] = xlfileDB.ID

		// append model store with the new element
		xlfile.Name = xlfileDB.Name_Data.String
		xlfile.Stage()
	}
	xlfileDB.CopyBasicFieldsToXLFile(xlfile)

	// preserve pointer to xlfileDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_XLFileDBID_XLFileDB)[xlfileDB hold variable pointers
	xlfileDB_Data := *xlfileDB
	preservedPtrToXLFile := &xlfileDB_Data
	(*backRepoXLFile.Map_XLFileDBID_XLFileDB)[xlfileDB.ID] = preservedPtrToXLFile

	return
}

// BackRepoXLFile.CheckoutPhaseTwo Checkouts all staged instances of XLFile to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLFile *BackRepoXLFileStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, xlfileDB := range *backRepoXLFile.Map_XLFileDBID_XLFileDB {
		backRepoXLFile.CheckoutPhaseTwoInstance(backRepo, xlfileDB)
	}
	return
}

// BackRepoXLFile.CheckoutPhaseTwoInstance Checkouts staged instances of XLFile to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoXLFile *BackRepoXLFileStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, xlfileDB *XLFileDB) (Error error) {

	xlfile := (*backRepoXLFile.Map_XLFileDBID_XLFilePtr)[xlfileDB.ID]
	_ = xlfile // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem xlfile.Sheets in the stage from the encode in the back repo
	// It parses all XLSheetDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	xlfile.Sheets = xlfile.Sheets[:0]
	// 2. loop all instances in the type in the association end
	for _, xlsheetDB_AssocEnd := range *backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if xlsheetDB_AssocEnd.XLFile_SheetsDBID.Int64 == int64(xlfileDB.ID) {
			// 4. fetch the associated instance in the stage
			xlsheet_AssocEnd := (*backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr)[xlsheetDB_AssocEnd.ID]
			// 5. append it the association slice
			xlfile.Sheets = append(xlfile.Sheets, xlsheet_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(xlfile.Sheets, func(i, j int) bool {
		xlsheetDB_i_ID := (*backRepo.BackRepoXLSheet.Map_XLSheetPtr_XLSheetDBID)[xlfile.Sheets[i]]
		xlsheetDB_j_ID := (*backRepo.BackRepoXLSheet.Map_XLSheetPtr_XLSheetDBID)[xlfile.Sheets[j]]

		xlsheetDB_i := (*backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetDB)[xlsheetDB_i_ID]
		xlsheetDB_j := (*backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetDB)[xlsheetDB_j_ID]

		return xlsheetDB_i.XLFile_SheetsDBID_Index.Int64 < xlsheetDB_j.XLFile_SheetsDBID_Index.Int64
	})

	return
}

// CommitXLFile allows commit of a single xlfile (if already staged)
func (backRepo *BackRepoStruct) CommitXLFile(xlfile *models.XLFile) {
	backRepo.BackRepoXLFile.CommitPhaseOneInstance(xlfile)
	if id, ok := (*backRepo.BackRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]; ok {
		backRepo.BackRepoXLFile.CommitPhaseTwoInstance(backRepo, id, xlfile)
	}
}

// CommitXLFile allows checkout of a single xlfile (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutXLFile(xlfile *models.XLFile) {
	// check if the xlfile is staged
	if _, ok := (*backRepo.BackRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]; ok {

		if id, ok := (*backRepo.BackRepoXLFile.Map_XLFilePtr_XLFileDBID)[xlfile]; ok {
			var xlfileDB XLFileDB
			xlfileDB.ID = id

			if err := backRepo.BackRepoXLFile.db.First(&xlfileDB, id).Error; err != nil {
				log.Panicln("CheckoutXLFile : Problem with getting object with id:", id)
			}
			backRepo.BackRepoXLFile.CheckoutPhaseOneInstance(&xlfileDB)
			backRepo.BackRepoXLFile.CheckoutPhaseTwoInstance(backRepo, &xlfileDB)
		}
	}
}

// CopyBasicFieldsFromXLFile
func (xlfileDB *XLFileDB) CopyBasicFieldsFromXLFile(xlfile *models.XLFile) {
	// insertion point for fields commit

	xlfileDB.Name_Data.String = xlfile.Name
	xlfileDB.Name_Data.Valid = true

	xlfileDB.NbSheets_Data.Int64 = int64(xlfile.NbSheets)
	xlfileDB.NbSheets_Data.Valid = true
}

// CopyBasicFieldsFromXLFileWOP
func (xlfileDB *XLFileDB) CopyBasicFieldsFromXLFileWOP(xlfile *XLFileWOP) {
	// insertion point for fields commit

	xlfileDB.Name_Data.String = xlfile.Name
	xlfileDB.Name_Data.Valid = true

	xlfileDB.NbSheets_Data.Int64 = int64(xlfile.NbSheets)
	xlfileDB.NbSheets_Data.Valid = true
}

// CopyBasicFieldsToXLFile
func (xlfileDB *XLFileDB) CopyBasicFieldsToXLFile(xlfile *models.XLFile) {
	// insertion point for checkout of basic fields (back repo to stage)
	xlfile.Name = xlfileDB.Name_Data.String
	xlfile.NbSheets = int(xlfileDB.NbSheets_Data.Int64)
}

// CopyBasicFieldsToXLFileWOP
func (xlfileDB *XLFileDB) CopyBasicFieldsToXLFileWOP(xlfile *XLFileWOP) {
	xlfile.ID = int(xlfileDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	xlfile.Name = xlfileDB.Name_Data.String
	xlfile.NbSheets = int(xlfileDB.NbSheets_Data.Int64)
}

// Backup generates a json file from a slice of all XLFileDB instances in the backrepo
func (backRepoXLFile *BackRepoXLFileStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "XLFileDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*XLFileDB, 0)
	for _, xlfileDB := range *backRepoXLFile.Map_XLFileDBID_XLFileDB {
		forBackup = append(forBackup, xlfileDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json XLFile ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json XLFile file", err.Error())
	}
}

// Backup generates a json file from a slice of all XLFileDB instances in the backrepo
func (backRepoXLFile *BackRepoXLFileStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*XLFileDB, 0)
	for _, xlfileDB := range *backRepoXLFile.Map_XLFileDBID_XLFileDB {
		forBackup = append(forBackup, xlfileDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("XLFile")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&XLFile_Fields, -1)
	for _, xlfileDB := range forBackup {

		var xlfileWOP XLFileWOP
		xlfileDB.CopyBasicFieldsToXLFileWOP(&xlfileWOP)

		row := sh.AddRow()
		row.WriteStruct(&xlfileWOP, -1)
	}
}

// RestoreXL from the "XLFile" sheet all XLFileDB instances
func (backRepoXLFile *BackRepoXLFileStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoXLFileid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["XLFile"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoXLFile.rowVisitorXLFile)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoXLFile *BackRepoXLFileStruct) rowVisitorXLFile(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var xlfileWOP XLFileWOP
		row.ReadStruct(&xlfileWOP)

		// add the unmarshalled struct to the stage
		xlfileDB := new(XLFileDB)
		xlfileDB.CopyBasicFieldsFromXLFileWOP(&xlfileWOP)

		xlfileDB_ID_atBackupTime := xlfileDB.ID
		xlfileDB.ID = 0
		query := backRepoXLFile.db.Create(xlfileDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoXLFile.Map_XLFileDBID_XLFileDB)[xlfileDB.ID] = xlfileDB
		BackRepoXLFileid_atBckpTime_newID[xlfileDB_ID_atBackupTime] = xlfileDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "XLFileDB.json" in dirPath that stores an array
// of XLFileDB and stores it in the database
// the map BackRepoXLFileid_atBckpTime_newID is updated accordingly
func (backRepoXLFile *BackRepoXLFileStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoXLFileid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "XLFileDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json XLFile file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*XLFileDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_XLFileDBID_XLFileDB
	for _, xlfileDB := range forRestore {

		xlfileDB_ID_atBackupTime := xlfileDB.ID
		xlfileDB.ID = 0
		query := backRepoXLFile.db.Create(xlfileDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoXLFile.Map_XLFileDBID_XLFileDB)[xlfileDB.ID] = xlfileDB
		BackRepoXLFileid_atBckpTime_newID[xlfileDB_ID_atBackupTime] = xlfileDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json XLFile file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<XLFile>id_atBckpTime_newID
// to compute new index
func (backRepoXLFile *BackRepoXLFileStruct) RestorePhaseTwo() {

	for _, xlfileDB := range *backRepoXLFile.Map_XLFileDBID_XLFileDB {

		// next line of code is to avert unused variable compilation error
		_ = xlfileDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoXLFile.db.Model(xlfileDB).Updates(*xlfileDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoXLFileid_atBckpTime_newID map[uint]uint
