package modules

import (
	"testing"
)

func TestGet{{.ModuleName}}Info(t *testing.T) {
	module := New{{.ModuleName}}()

	result := module.Get{{.ModuleName}}Info()
	expected := "Module: {{.ModuleName}}"

	if result != expected {
		t.Fatalf("Expected %s, but got %s", expected, result)
	}
}
