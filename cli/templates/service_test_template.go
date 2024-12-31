package services

import (
	"testing"
)

func TestGet{{.ServiceName}}Data(t *testing.T) {
	service := New{{.ServiceName}}Service()

	result := service.Get{{.ServiceName}}Data()
	expected := "Data for {{.ServiceName}}"

	if result != expected {
		t.Fatalf("Expected %s, but got %s", expected, result)
	}
}
