package controllers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func TestGet{{.ControllerName}}(t *testing.T) {
	router := gin.Default()
	ctrl := New{{.ControllerName}}Controller()

	router.GET("/{{.ControllerName}}", ctrl.Get{{.ControllerName}})

	req, _ := http.NewRequest("GET", "/{{.ControllerName}}", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
