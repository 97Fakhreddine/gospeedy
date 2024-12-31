package core

import (
	"fmt"
	"net/http"
)

// Controller defines a common interface for all controllers
type Controller interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

// BaseController provides common controller functionality
type BaseController struct {
}

// HandleRequest provides default request handling logic
func (c *BaseController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Request handled by %T", c)))
}
