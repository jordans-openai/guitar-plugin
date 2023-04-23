package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChordsHandler(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "Valid chords",
			method:         http.MethodPost,
			requestBody:    `{"chords":[{"name":"Cmaj7"},{"name":"Dm"}]}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid chord",
			method:         http.MethodPost,
			requestBody:    `{"chords":[{"name":"Cmaj7"},{"name":"INVALID_CHORD"}]}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			requestBody:    "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid request body",
			method:         http.MethodPost,
			requestBody:    `INVALID_JSON`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, "/chords", strings.NewReader(tc.requestBody))
			rr := httptest.NewRecorder()
			ChordsHandler(rr, req)
			assert.Equal(t, tc.expectedStatus, rr.Code)
			if rr.Code == http.StatusOK {
				var res TabsResponse
				err := json.NewDecoder(rr.Body).Decode(&res)
				assert.NoError(t, err)
				assert.Equal(t, len(res.Tabs), 2)
			}
		})
	}
}
