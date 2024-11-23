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

type Company struct {
	*gorm.Model
	Name         string `json:"name"`
	Description  string `json:"description"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
}
