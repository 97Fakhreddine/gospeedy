package resolvers

import (
	"fmt"
)

// {{.ModuleName}}Resolver is responsible for resolving the GraphQL queries related to {{.ModuleName}} module
type {{.ModuleName}}Resolver struct {}

// New{{.ModuleName}}Resolver creates a new resolver for {{.ModuleName}} module
func New{{.ModuleName}}Resolver() *{{.ModuleName}}Resolver {
	return &{{.ModuleName}}Resolver{}
}

// Resolve{{.ModuleName}}Query resolves the GraphQL query for {{.ModuleName}}
func (r *{{.ModuleName}}Resolver) Resolve{{.ModuleName}}Query() string {
	return fmt.Sprintf("Resolved {{.ModuleName}} query")
}
