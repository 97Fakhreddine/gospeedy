package mutations

import (
	"fmt"
)

// {{.ModuleName}}Mutation is responsible for handling mutations for {{.ModuleName}} module
type {{.ModuleName}}Mutation struct {}

// New{{.ModuleName}}Mutation creates a new mutation handler for {{.ModuleName}} module
func New{{.ModuleName}}Mutation() *{{.ModuleName}}Mutation {
	return &{{.ModuleName}}Mutation{}
}

// Execute{{.ModuleName}}Mutation executes the mutation for {{.ModuleName}} module
func (m *{{.ModuleName}}Mutation) Execute{{.ModuleName}}Mutation() string {
	return fmt.Sprintf("Mutation for {{.ModuleName}} executed")
}
