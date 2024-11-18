package handlers

import (
	"encoding/json"
	"fmt"
	"hospital-middleware/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatientHandler struct {
	db *gorm.DB
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
	return &PatientHandler{db: db}
}

func (h *PatientHandler) SearchPatient(c *gin.Context) {
	// Get staff hospital ID from JWT token
	hospitalID, exists := c.Get("hospitalID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var searchParams struct {
		NationalID  string `json:"national_id"`
		PassportID  string `json:"passport_id"`
		FirstName   string `json:"first_name"`
		MiddleName  string `json:"middle_name"`
		LastName    string `json:"last_name"`
		DateOfBirth string `json:"date_of_birth"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}

	if err := c.ShouldBindJSON(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build query
	query := h.db.Where("hospital_id = ?", hospitalID)

	if searchParams.NationalID != "" {
		query = query.Where("national_id = ?", searchParams.NationalID)
	}
	if searchParams.PassportID != "" {
		query = query.Where("passport_id = ?", searchParams.PassportID)
	}
	if searchParams.FirstName != "" {
		query = query.Where("first_name_en ILIKE ? OR first_name_th ILIKE ?", 
			"%"+searchParams.FirstName+"%", "%"+searchParams.FirstName+"%")
	}
	// Add other search parameters...

	var patients []models.Patient
	if err := query.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) SearchExternalPatient(c *gin.Context) {
	id := c.Param("id")
	
	// Call Hospital A API
	resp, err := http.Get(fmt.Sprintf("https://hospital-a.api.co.th/patient/search/%s", id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch external patient data"})
		return
	}
	defer resp.Body.Close()

	var patientData models.Patient
	if err := json.NewDecoder(resp.Body).Decode(&patientData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse external patient data"})
		return
	}

	c.JSON(http.StatusOK, patientData)
}