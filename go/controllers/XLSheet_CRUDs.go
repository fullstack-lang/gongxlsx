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
var __XLSheet__dummysDeclaration__ models.XLSheet
var __XLSheet_time__dummyDeclaration time.Duration

var mutexXLSheet sync.Mutex

// An XLSheetID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXLSheet updateXLSheet deleteXLSheet
type XLSheetID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XLSheetInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXLSheet updateXLSheet
type XLSheetInput struct {
	// The XLSheet to submit or modify
	// in: body
	XLSheet *orm.XLSheetAPI
}

// GetXLSheets
//
// swagger:route GET /xlsheets xlsheets getXLSheets
//
// # Get all xlsheets
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) GetXLSheets(c *gin.Context) {

	// source slice
	var xlsheetDBs []orm.XLSheetDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLSheets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxlsx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	query := db.Find(&xlsheetDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xlsheetAPIs := make([]orm.XLSheetAPI, 0)

	// for each xlsheet, update fields from the database nullable fields
	for idx := range xlsheetDBs {
		xlsheetDB := &xlsheetDBs[idx]
		_ = xlsheetDB
		var xlsheetAPI orm.XLSheetAPI

		// insertion point for updating fields
		xlsheetAPI.ID = xlsheetDB.ID
		xlsheetDB.CopyBasicFieldsToXLSheet(&xlsheetAPI.XLSheet)
		xlsheetAPI.XLSheetPointersEnconding = xlsheetDB.XLSheetPointersEnconding
		xlsheetAPIs = append(xlsheetAPIs, xlsheetAPI)
	}

	c.JSON(http.StatusOK, xlsheetAPIs)
}

// PostXLSheet
//
// swagger:route POST /xlsheets xlsheets postXLSheet
//
// Creates a xlsheet
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXLSheets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxlsx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Validate input
	var input orm.XLSheetAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xlsheet
	xlsheetDB := orm.XLSheetDB{}
	xlsheetDB.XLSheetPointersEnconding = input.XLSheetPointersEnconding
	xlsheetDB.CopyBasicFieldsFromXLSheet(&input.XLSheet)

	query := db.Create(&xlsheetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXLSheet.CheckoutPhaseOneInstance(&xlsheetDB)
	xlsheet := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]

	if xlsheet != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xlsheet)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xlsheetDB)

	mutexXLSheet.Unlock()
}

// GetXLSheet
//
// swagger:route GET /xlsheets/{ID} xlsheets getXLSheet
//
// Gets the details for a xlsheet.
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) GetXLSheet(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXLSheet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxlsx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Get xlsheetDB in DB
	var xlsheetDB orm.XLSheetDB
	if err := db.First(&xlsheetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xlsheetAPI orm.XLSheetAPI
	xlsheetAPI.ID = xlsheetDB.ID
	xlsheetAPI.XLSheetPointersEnconding = xlsheetDB.XLSheetPointersEnconding
	xlsheetDB.CopyBasicFieldsToXLSheet(&xlsheetAPI.XLSheet)

	c.JSON(http.StatusOK, xlsheetAPI)
}

// UpdateXLSheet
//
// swagger:route PATCH /xlsheets/{ID} xlsheets updateXLSheet
//
// # Update a xlsheet
//
// Responses:
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) UpdateXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXLSheet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxlsx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Validate input
	var input orm.XLSheetAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xlsheetDB orm.XLSheetDB

	// fetch the xlsheet
	query := db.First(&xlsheetDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xlsheetDB.CopyBasicFieldsFromXLSheet(&input.XLSheet)
	xlsheetDB.XLSheetPointersEnconding = input.XLSheetPointersEnconding

	query = db.Model(&xlsheetDB).Updates(xlsheetDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xlsheetNew := new(models.XLSheet)
	xlsheetDB.CopyBasicFieldsToXLSheet(xlsheetNew)

	// get stage instance from DB instance, and call callback function
	xlsheetOld := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]
	if xlsheetOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xlsheetOld, xlsheetNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xlsheetDB
	c.JSON(http.StatusOK, xlsheetDB)

	mutexXLSheet.Unlock()
}

// DeleteXLSheet
//
// swagger:route DELETE /xlsheets/{ID} xlsheets deleteXLSheet
//
// # Delete a xlsheet
//
// default: genericError
//
//	200: xlsheetDBResponse
func (controller *Controller) DeleteXLSheet(c *gin.Context) {

	mutexXLSheet.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXLSheet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxlsx/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXLSheet.GetDB()

	// Get model if exist
	var xlsheetDB orm.XLSheetDB
	if err := db.First(&xlsheetDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xlsheetDB)

	// get an instance (not staged) from DB instance, and call callback function
	xlsheetDeleted := new(models.XLSheet)
	xlsheetDB.CopyBasicFieldsToXLSheet(xlsheetDeleted)

	// get stage instance from DB instance, and call callback function
	xlsheetStaged := backRepo.BackRepoXLSheet.Map_XLSheetDBID_XLSheetPtr[xlsheetDB.ID]
	if xlsheetStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xlsheetStaged, xlsheetDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexXLSheet.Unlock()
}
