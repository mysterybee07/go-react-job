package payloads

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
	Name         string `json:"name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	Address      string `json:"address"`
	ImageUrl     string `json:"image_url,omitempty"`
	Description  string `json:"description,omitempty"`
	Password     string `json:"password"`
}
