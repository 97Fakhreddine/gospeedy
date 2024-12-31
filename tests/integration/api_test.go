// api_test.go
package integration_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "gospeedy"
    "net/http"
    "net/http/httptest"
)

func TestUserAPI(t *testing.T) {
    // Create a new instance of the GoSpeedy app
    app := gospeedy.NewApp()

    // Create a mock HTTP request
    req, err := http.NewRequest("GET", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Record the HTTP response
    recorder := httptest.NewRecorder()
    app.ServeHTTP(recorder, req)

    // Assert status code and response
    assert.Equal(t, http.StatusOK, recorder.Code)
    assert.Contains(t, recorder.Body.String(), "users")
}
