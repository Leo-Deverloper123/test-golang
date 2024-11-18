package models

import (
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name     string    `json:"name"`
	APIKey   string    `json:"api_key"`
	Patients []Patient `json:"patients"`
	Staff    []Staff   `json:"staff"`
}