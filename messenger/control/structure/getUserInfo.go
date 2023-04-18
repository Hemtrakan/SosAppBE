package structure

import "time"

type UserRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID           string    `json:"id"`
		PhoneNumber  string    `json:"phoneNumber"`
		FirstName    string    `json:"firstName"`
		LastName     string    `json:"lastName"`
		Email        string    `json:"email"`
		Birthday     time.Time `json:"birthday"`
		Gender       string    `json:"gender"`
		ImageProfile string    `json:"imageProfile"`
		Workplace    string    `json:"workplace,omitempty"`
		IdCard       IdCard    `json:"idCard,omitempty"`
		Address      Address   `json:"address,omitempty"`
	} `json:"data"`
}

type IdCard struct {
	TextIDCard string `json:"textIDCard,omitempty"`
	PathImage  string `json:"pathImage,omitempty"`
	Verify     bool   `json:"verify"`
}

type Address struct {
	Address     string `json:"address,omitempty"`
	SubDistrict string `json:"subDistrict,omitempty"`
	District    string `json:"district,omitempty"`
	Province    string `json:"province,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Country     string `json:"country,omitempty"`
}
