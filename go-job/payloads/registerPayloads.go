package payloads

import "mime/multipart"

type RegisterUser struct {
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	Address      string `jsong:"address"`
	ImageUrl     string `json:"image_url,omitempty"`
	Resume       string `json:"resume,omitempty"`
	Password     string `json:"password"`
}

type RegisterCompany struct {
	Name         string                `form:"name" json:"name"`
	ContactEmail string                `form:"contact_email" json:"contact_email"`
	ContactPhone string                `form:"contact_phone" json:"contact_phone"`
	Address      string                `form:"address" json:"address"`
	Description  string                `form:"description" json:"description"`
	Password     string                `form:"password" json:"password"`
	ImageUrl     *multipart.FileHeader `form:"image_url" json:"image_url"`
}
