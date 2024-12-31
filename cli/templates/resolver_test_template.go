package resolvers

import (
	"testing"
)

func TestResolve{{.ModuleName}}Query(t *testing.T) {
	resolver := New{{.ModuleName}}Resolver()

	result := resolver.Resolve{{.ModuleName}}Query()
	expected := "Resolved {{.ModuleName}} query"

	if result != expected {
		t.Fatalf("Expected %s, but got %s", expected, result)
	}
}
