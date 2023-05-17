package esendeo

type Address struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Company      string `json:"company"`
	ShopName     string `json:"shopName"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	PhoneNumber  string `json:"phoneNumber"`
	MobileNumber string `json:"mobileNumber"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	City         string `json:"city"`
	PostalCode   string `json:"postalCode"`
	CountryCode  string `json:"countryCode"`
	Instructions string `json:"instructions"`
	Comment      string `json:"comment"`
}
