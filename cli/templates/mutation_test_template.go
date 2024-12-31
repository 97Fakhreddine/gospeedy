package mutations

import (
	"testing"
)

func TestExecute{{.ModuleName}}Mutation(t *testing.T) {
	mutation := New{{.ModuleName}}Mutation()

	result := mutation.Execute{{.ModuleName}}Mutation()
	expected := "Mutation for {{.ModuleName}} executed"

	if result != expected {
		t.Fatalf("Expected %s, but got %s", expected, result)
	}
}
