package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexRoute(t *testing.T) {
	test := struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		description:   "index route",
		route:         "/",
		expectedError: false,
		expectedCode:  200,
		expectedBody:  "Hello, World!",
	}

	app := Setup()

	req, _ := http.NewRequest(
		"GET",
		test.route,
		nil,
	)

	res, err := app.Test(req, -1)

	assert.Equalf(t, test.expectedError, err != nil, test.description)

	assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, test.description)

	assert.Equalf(t, test.expectedBody, string(body), test.description)
}

func TestGetAllItem(t *testing.T) {
	tests := []struct {
		name         string
		route        string
		wantErr      bool
		expectedErr  error
		expectedCode int
	}{
		{
			name:         "create simple item",
			route:        "api/v1/item",
			wantErr:      false,
			expectedErr:  nil,
			expectedCode: 200,
		},
	}

	app := Setup()

	for _, tt := range tests {
		req, _ := http.NewRequest(
			"GET",
			tt.route,
			nil,
		)

		fmt.Println(req.Body)

		req.Close = true
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		defer res.Body.Close()

		fmt.Println("Error: ", err)

		require.Equalf(t, tt.expectedErr, err, fmt.Sprintf("Expected error: %s, but got %s", tt.expectedErr, err))

		require.Equalf(t, tt.expectedCode, res.StatusCode, fmt.Sprintf("Expected status code %d but got %d", tt.expectedCode, res.StatusCode))

		jsonBody, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(jsonBody))

		require.Nilf(t, err, "Didn't expect error")
	}
}
