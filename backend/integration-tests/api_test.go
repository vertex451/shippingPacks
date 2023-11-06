package integration_tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"shippingPacks/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoints(t *testing.T) {
	// Define test cases for the HTTP requests
	testCases := []struct {
		url            string
		expectedCode   int
		expectedResult map[int]int // Replace with the expected response body
		expectedErr    string      // Replace with the expected error value
	}{
		{
			url:            "/api/v1/get-packs-number/aaa",
			expectedCode:   http.StatusBadRequest,
			expectedResult: nil,
			expectedErr:    "failed to convert itemsOrdered to number\n",
		},
		{
			url:            "/api/v1/get-packs-number/-1",
			expectedCode:   http.StatusBadRequest,
			expectedResult: nil,
			expectedErr:    "itemsOrdered should not be negative\n",
		},
		{
			url:            "/api/v1/get-packs-number/12001",
			expectedCode:   http.StatusOK,
			expectedResult: map[int]int{5000: 2, 2000: 1, 250: 1},
			expectedErr:    "",
		},
	}

	cfg, err := config.LoadConfig("../.env")
	assert.Nil(t, err)

	// Iterate through test cases and run the tests
	for _, tc := range testCases {
		res, err := http.Get(fmt.Sprintf("http://localhost:%s%s", cfg.Port, tc.url))
		if err != nil {
			t.Fatalf("Failed to create HTTP request: %v", err)
		}

		defer res.Body.Close()

		// Check the HTTP status code
		if status := res.StatusCode; status != tc.expectedCode {
			t.Errorf("Expected status code %d, but got %d for URL: %s", tc.expectedCode, status, tc.url)
		}

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if tc.expectedResult != nil {
			var actual map[int]int
			err = json.Unmarshal(body, &actual)
			assert.Nil(t, err)
			if !reflect.DeepEqual(tc.expectedResult, actual) {
				t.Errorf("Maps not equal. Expected: %v, but got: %v", tc.expectedResult, actual)
			}
		}

		// Check the error value
		if tc.expectedErr != "" {
			assert.Equal(t, tc.expectedErr, string(body))
		}
	}
}
