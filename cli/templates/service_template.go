package services

import (
	"fmt"
)

// {{.ServiceName}}Service handles the business logic for {{.ServiceName}} module
type {{.ServiceName}}Service struct {}

// New{{.ServiceName}}Service creates a new service
func New{{.ServiceName}}Service() *{{.ServiceName}}Service {
	return &{{.ServiceName}}Service{}
}

// Get{{.ServiceName}}Data retrieves the data for {{.ServiceName}}
func (srv *{{.ServiceName}}Service) Get{{.ServiceName}}Data() string {
	// Add your business logic here
	return fmt.Sprintf("Data for {{.ServiceName}}")
}
