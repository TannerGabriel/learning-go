package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/gabriel", nil)
	res := httptest.NewRecorder()
	NewServer().ServeHTTP(res, req)

	expectedResponse := "Hello, gabriel"
	errorMessage := fmt.Sprintf("Expected %s but got ", expectedResponse)

	if res.Body.String() != expectedResponse {
		t.Error(errorMessage, res.Body.String())
	}
}
