package models

import "gorm.io/gorm"

type ContactInfo struct {
	Name         string `gorm:"not null" json:"name"`
	ContactEmail string `gorm:"unique" json:"contact_email"`
	ContactPhone string `gorm:"unique" json:"contact_phone"`
	ImageUrl     string `gorm:"type:varchar(255)" json:"image_url,omitempty"`
}

type User struct {
	gorm.Model
	ContactInfo
	Resume   string `gorm:"type:varchar(255)" json:"resume,omitempty"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
}

type Company struct {
	gorm.Model
	ContactInfo
	Description string `json:"description"`
}

type RegisterUser struct {
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	ImageUrl     string `json:"image_url,omitempty"`
	Resume       string `json:"resume,omitempty"`
	Password     string `json:"password"`
}
