package models

import "gorm.io/gorm"

type ContactInfo struct {
	Name         string `gorm:"not null" json:"name"`
	ContactEmail string `gorm:"unique" json:"contact_email"`
	ContactPhone string `gorm:"unique" json:"contact_phone"`
	Address      string `json:"address"`
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
	Password    string `json:"-"`
}
