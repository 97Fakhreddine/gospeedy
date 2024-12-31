package modules

import (
	"fmt"
)

// {{.ModuleName}} represents the {{.ModuleName}} module
type {{.ModuleName}} struct {
	Name string
}

// New{{.ModuleName}} creates a new instance of the {{.ModuleName}} module
func New{{.ModuleName}}() *{{.ModuleName}} {
	return &{{.ModuleName}}{Name: "{{.ModuleName}}"}
}

// Get{{.ModuleName}}Info returns information about the module
func (m *{{.ModuleName}}) Get{{.ModuleName}}Info() string {
	return fmt.Sprintf("Module: %s", m.Name)
}
