package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// {{.ControllerName}}Controller handles the requests for {{.ControllerName}} module
type {{.ControllerName}}Controller struct {}

// New{{.ControllerName}}Controller creates a new controller
func New{{.ControllerName}}Controller() *{{.ControllerName}}Controller {
	return &{{.ControllerName}}Controller{}
}

// Get{{.ControllerName}} handles GET requests for {{.ControllerName}} module
func (ctrl *{{.ControllerName}}Controller) Get{{.ControllerName}}(c *gin.Context) {
	// Add your business logic here
	c.JSON(http.StatusOK, gin.H{"message": "{{.ControllerName}} fetched successfully"})
}
