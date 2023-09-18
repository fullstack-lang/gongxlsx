// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxlsx/go/models"
	"github.com/fullstack-lang/gongxlsx/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __XLFile__dummysDeclaration__ models.XLFile
var __XLFile_time__dummyDeclaration time.Duration

var mutexXLFile sync.Mutex

// An XLFileID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXLFile updateXLFile deleteXLFile
type XLFileID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XLFileInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXLFile updateXLFile
type XLFileInput struct {
	// The XLFile to submit or modify
	// in: body
	XLFile *orm.XLFileAPI
}

// GetXLFiles
//
// swagger:route GET /xlfiles xlfiles getXLFiles
//
// # Get all xlfiles
//
// Responses:
// default: genericError
//
//	200: xlfileDBResponse
func (controller *Controller) GetXLFiles(c *gin.Context) {

	// source slice
	var xlfileDBs []orm.XLFileDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLFiles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoXLFile.GetDB()

	query := db.Find(&xlfileDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlfileAPIs := make([]orm.XLFileAPI, 0)

	// for each xlfile, update fields from the database nullable fields
	for idx := range xlfileDBs {
		xlfileDB := &xlfileDBs[idx]
		_ = xlfileDB
		var xlfileAPI orm.XLFileAPI

		// insertion point for updating fields
		xlfileAPI.ID = xlfileDB.ID
		xlfileDB.CopyBasicFieldsToXLFile(&xlfileAPI.XLFile)
		xlfileAPI.XLFilePointersEnconding = xlfileDB.XLFilePointersEnconding
		xlfileAPIs = append(xlfileAPIs, xlfileAPI)
	}

	c.JSON(http.StatusOK, xlfileAPIs)
}

// PostXLFile
//
// swagger:route POST /xlfiles xlfiles postXLFile
//
// Creates a xlfile
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXLFile(c *gin.Context) {

	mutexXLFile.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXLFiles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoXLFile.GetDB()

	// Validate input
	var input orm.XLFileAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlfile
	xlfileDB := orm.XLFileDB{}
	xlfileDB.XLFilePointersEnconding = input.XLFilePointersEnconding
	xlfileDB.CopyBasicFieldsFromXLFile(&input.XLFile)

	query := db.Create(&xlfileDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXLFile.CheckoutPhaseOneInstance(&xlfileDB)
	xlfile := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[xlfileDB.ID]

	if xlfile != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlfile)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlfileDB)

	mutexXLFile.Unlock()
}

// GetXLFile
//
// swagger:route GET /xlfiles/{ID} xlfiles getXLFile
//
// Gets the details for a xlfile.
//
// Responses:
// default: genericError
//
//	200: xlfileDBResponse
func (controller *Controller) GetXLFile(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoXLFile.GetDB()

	// Get xlfileDB in DB
	var xlfileDB orm.XLFileDB
	if err := db.First(&xlfileDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlfileAPI orm.XLFileAPI
	xlfileAPI.ID = xlfileDB.ID
	xlfileAPI.XLFilePointersEnconding = xlfileDB.XLFilePointersEnconding
	xlfileDB.CopyBasicFieldsToXLFile(&xlfileAPI.XLFile)

	c.JSON(http.StatusOK, xlfileAPI)
}

// UpdateXLFile
//
// swagger:route PATCH /xlfiles/{ID} xlfiles updateXLFile
//
// # Update a xlfile
//
// Responses:
// default: genericError
//
//	200: xlfileDBResponse
func (controller *Controller) UpdateXLFile(c *gin.Context) {

	mutexXLFile.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXLFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoXLFile.GetDB()

	// Validate input
	var input orm.XLFileAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlfileDB orm.XLFileDB

	// fetch the xlfile
	query := db.First(&xlfileDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlfileDB.CopyBasicFieldsFromXLFile(&input.XLFile)
	xlfileDB.XLFilePointersEnconding = input.XLFilePointersEnconding

	query = db.Model(&xlfileDB).Updates(xlfileDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlfileNew := new(models.XLFile)
	xlfileDB.CopyBasicFieldsToXLFile(xlfileNew)

	// get stage instance from DB instance, and call callback function
	xlfileOld := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[xlfileDB.ID]
	if xlfileOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlfileOld, xlfileNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlfileDB
	c.JSON(http.StatusOK, xlfileDB)

	mutexXLFile.Unlock()
}

// DeleteXLFile
//
// swagger:route DELETE /xlfiles/{ID} xlfiles deleteXLFile
//
// # Delete a xlfile
//
// default: genericError
//
//	200: xlfileDBResponse
func (controller *Controller) DeleteXLFile(c *gin.Context) {

	mutexXLFile.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXLFile", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoXLFile.GetDB()

	// Get model if exist
	var xlfileDB orm.XLFileDB
	if err := db.First(&xlfileDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xlfileDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlfileDeleted := new(models.XLFile)
	xlfileDB.CopyBasicFieldsToXLFile(xlfileDeleted)

	// get stage instance from DB instance, and call callback function
	xlfileStaged := backRepo.BackRepoXLFile.Map_XLFileDBID_XLFilePtr[xlfileDB.ID]
	if xlfileStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlfileStaged, xlfileDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexXLFile.Unlock()
}
