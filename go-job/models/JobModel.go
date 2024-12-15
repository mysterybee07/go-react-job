package models

import "gorm.io/gorm"

type Job struct {
	*gorm.Model
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Salary      string  `json:"salary"`
	CompanyID   uint    `json:"company_id"`
	Company     Company `json:"company" gorm:"foreignKey:CompanyID"`
}
