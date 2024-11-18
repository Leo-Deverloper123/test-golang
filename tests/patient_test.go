package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchPatient(t *testing.T) {
	router := setupTestRouter()

	// Login first to get token
	loginPayload := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
		"hospital": 1,
	}
	jsonLogin, _ := json.Marshal(loginPayload)
	loginReq, _ := http.NewRequest("POST", "/staff/login", bytes.NewBuffer(jsonLogin))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, loginReq)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	token := response["token"]

	tests := []struct {
		name       string
		query      map[string]interface{}
		token      string
		statusCode int
	}{
		{
			name: "Valid Search",
			query: map[string]interface{}{
				"national_id": "1234567890123",
			},
			token:      token,
			statusCode: http.StatusOK,
		},
		{
			name: "Search Without Token",
			query: map[string]interface{}{
				"national_id": "1234567890123",
			},
			token:      "",
			statusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.query)
			req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(jsonValue))
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}