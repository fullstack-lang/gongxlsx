
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongxslx/go/orm"
)

// genQuery return the name of the column
func genQuery( columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongxslx/go")
	{// insertion point for registrations
		v1.GET("/v1/xlfiles", GetXLFiles)
		v1.GET("/v1/xlfiles/:id", GetXLFile)
		v1.POST("/v1/xlfiles", PostXLFile)
		v1.PATCH("/v1/xlfiles/:id", UpdateXLFile)
		v1.PUT("/v1/xlfiles/:id", UpdateXLFile)
		v1.DELETE("/v1/xlfiles/:id", DeleteXLFile)

		v1.GET("/v1/xlsheets", GetXLSheets)
		v1.GET("/v1/xlsheets/:id", GetXLSheet)
		v1.POST("/v1/xlsheets", PostXLSheet)
		v1.PATCH("/v1/xlsheets/:id", UpdateXLSheet)
		v1.PUT("/v1/xlsheets/:id", UpdateXLSheet)
		v1.DELETE("/v1/xlsheets/:id", DeleteXLSheet)


		v1.GET("/commitnb", GetLastCommitNb)
	}
}

// swagger:route GET /commitnb backrepo GetLastCommitNb
func GetLastCommitNb(c *gin.Context) {
	res := orm.GetLastCommitNb()

	c.JSON(http.StatusOK, res)
}
