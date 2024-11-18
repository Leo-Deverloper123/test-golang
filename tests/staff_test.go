package tests

import (
	"bytes"
	"encoding/json"
	"hospital-middleware/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateStaff(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		payload    models.Staff
		statusCode int
	}{
		{
			name: "Valid Staff Creation",
			payload: models.Staff{
				Username:   "testuser",
				Password:   "password123",
				HospitalID: 1,
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "Invalid Staff Creation - Missing Username",
			payload: models.Staff{
				Password:   "password123",
				HospitalID: 1,
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonValue, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/staff/create", bytes.NewBuffer(jsonValue))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}