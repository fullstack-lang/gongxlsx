// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongxlsx/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	displayselectionDBs map[uint]*DisplaySelectionDB

	nextIDDisplaySelectionDB uint

	xlcellDBs map[uint]*XLCellDB

	nextIDXLCellDB uint

	xlfileDBs map[uint]*XLFileDB

	nextIDXLFileDB uint

	xlrowDBs map[uint]*XLRowDB

	nextIDXLRowDB uint

	xlsheetDBs map[uint]*XLSheetDB

	nextIDXLSheetDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		displayselectionDBs: make(map[uint]*DisplaySelectionDB),

		xlcellDBs: make(map[uint]*XLCellDB),

		xlfileDBs: make(map[uint]*XLFileDB),

		xlrowDBs: make(map[uint]*XLRowDB),

		xlsheetDBs: make(map[uint]*XLSheetDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *DisplaySelectionDB:
		db.nextIDDisplaySelectionDB++
		v.ID = db.nextIDDisplaySelectionDB
		db.displayselectionDBs[v.ID] = v
	case *XLCellDB:
		db.nextIDXLCellDB++
		v.ID = db.nextIDXLCellDB
		db.xlcellDBs[v.ID] = v
	case *XLFileDB:
		db.nextIDXLFileDB++
		v.ID = db.nextIDXLFileDB
		db.xlfileDBs[v.ID] = v
	case *XLRowDB:
		db.nextIDXLRowDB++
		v.ID = db.nextIDXLRowDB
		db.xlrowDBs[v.ID] = v
	case *XLSheetDB:
		db.nextIDXLSheetDB++
		v.ID = db.nextIDXLSheetDB
		db.xlsheetDBs[v.ID] = v
	default:
		return nil, errors.New("unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		delete(db.displayselectionDBs, v.ID)
	case *XLCellDB:
		delete(db.xlcellDBs, v.ID)
	case *XLFileDB:
		delete(db.xlfileDBs, v.ID)
	case *XLRowDB:
		delete(db.xlrowDBs, v.ID)
	case *XLSheetDB:
		delete(db.xlsheetDBs, v.ID)
	default:
		return nil, errors.New("unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		db.displayselectionDBs[v.ID] = v
		return db, nil
	case *XLCellDB:
		db.xlcellDBs[v.ID] = v
		return db, nil
	case *XLFileDB:
		db.xlfileDBs[v.ID] = v
		return db, nil
	case *XLRowDB:
		db.xlrowDBs[v.ID] = v
		return db, nil
	case *XLSheetDB:
		db.xlsheetDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *DisplaySelectionDB:
		if existing, ok := db.displayselectionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *XLCellDB:
		if existing, ok := db.xlcellDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *XLFileDB:
		if existing, ok := db.xlfileDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *XLRowDB:
		if existing, ok := db.xlrowDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *XLSheetDB:
		if existing, ok := db.xlsheetDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	default:
		return nil, errors.New("unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]DisplaySelectionDB:
        *ptr = make([]DisplaySelectionDB, 0, len(db.displayselectionDBs))
        for _, v := range db.displayselectionDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]XLCellDB:
        *ptr = make([]XLCellDB, 0, len(db.xlcellDBs))
        for _, v := range db.xlcellDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]XLFileDB:
        *ptr = make([]XLFileDB, 0, len(db.xlfileDBs))
        for _, v := range db.xlfileDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]XLRowDB:
        *ptr = make([]XLRowDB, 0, len(db.xlrowDBs))
        for _, v := range db.xlrowDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]XLSheetDB:
        *ptr = make([]XLSheetDB, 0, len(db.xlsheetDBs))
        for _, v := range db.xlsheetDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *DisplaySelectionDB:
		tmp, ok := db.displayselectionDBs[uint(i)]

		displayselectionDB, _ := instanceDB.(*DisplaySelectionDB)
		*displayselectionDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *XLCellDB:
		tmp, ok := db.xlcellDBs[uint(i)]

		xlcellDB, _ := instanceDB.(*XLCellDB)
		*xlcellDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *XLFileDB:
		tmp, ok := db.xlfileDBs[uint(i)]

		xlfileDB, _ := instanceDB.(*XLFileDB)
		*xlfileDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *XLRowDB:
		tmp, ok := db.xlrowDBs[uint(i)]

		xlrowDB, _ := instanceDB.(*XLRowDB)
		*xlrowDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *XLSheetDB:
		tmp, ok := db.xlsheetDBs[uint(i)]

		xlsheetDB, _ := instanceDB.(*XLSheetDB)
		*xlsheetDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("Unkown type")
	}
	
	return db, nil
}

