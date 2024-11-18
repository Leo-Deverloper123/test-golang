package models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Username   string   `json:"username" gorm:"unique"`
	Password   string   `json:"-"` // Password will be hashed
	HospitalID uint     `json:"hospital_id"`
	Hospital   Hospital `json:"hospital"`
}