package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstNameTH   string    `json:"first_name_th"`
	MiddleNameTH  string    `json:"middle_name_th"`
	LastNameTH    string    `json:"last_name_th"`
	FirstNameEN   string    `json:"first_name_en"`
	MiddleNameEN  string    `json:"middle_name_en"`
	LastNameEN    string    `json:"last_name_en"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	PatientHN     string    `json:"patient_hn" gorm:"unique"`
	NationalID    string    `json:"national_id" gorm:"unique"`
	PassportID    string    `json:"passport_id" gorm:"unique"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	Gender        string    `json:"gender" gorm:"type:char(1)"`
	HospitalID    uint      `json:"hospital_id"`
	Hospital      Hospital  `json:"hospital"`
}