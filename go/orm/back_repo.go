// generated by genORMTranslation.go
package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/fullstack-lang/gongxslx/go/models"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoXLFile BackRepoXLFileStruct

	BackRepoXLSheet BackRepoXLSheetStruct

	CommitNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level
}

func (backRepo *BackRepoStruct) GetLastCommitNb() uint {
	return backRepo.CommitNb
}

func (backRepo *BackRepoStruct) IncrementCommitNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitNb = backRepo.CommitNb + 1
	return backRepo.CommitNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoXLFile.Init(db)
	backRepo.BackRepoXLSheet.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoXLFile.CommitPhaseOne(stage)
	backRepo.BackRepoXLSheet.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoXLFile.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXLSheet.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoXLFile.CheckoutPhaseOne()
	backRepo.BackRepoXLSheet.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoXLFile.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXLSheet.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitNb() uint {
	return BackRepo.GetLastCommitNb()
}