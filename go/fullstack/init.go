package fullstack

import (
	// gongxlsx stack for model analysis

	gongxlsx_controllers "github.com/fullstack-lang/gongxlsx/go/controllers"
	gongxlsx_orm "github.com/fullstack-lang/gongxlsx/go/orm"
	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gongxlsx/ng/projects"
)

func Init(r *gin.Engine, filenames ...string) {

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	db_inMemory := gongxlsx_orm.SetupModels(false, filenames[0])

	// since gongxlsxsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	gongxlsx_controllers.RegisterControllers(r)
}
