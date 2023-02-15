package request

import "time"

type UserProfile struct {
	PhoneNumber  string    `json:"phoneNumber,omitempty"`
	FirstName    string    `json:"firstName,omitempty"`
	LastName     string    `json:"lastName,omitempty"`
	Email        string    `json:"email,omitempty"`
	Birthday     time.Time `json:"birthday,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	ImageProfile string    `json:"imageProfile,omitempty"`
	IdCard       IdCard    `json:"idCard,omitempty"`
	Address      Address   `json:"address,omitempty"`
}

type IdCard struct {
	TextIDCard string `json:"textIDCard,omitempty"`
	PathImage  string `json:"pathImage,omitempty"`
}

type Address struct {
	Address     string `json:"address,omitempty"`
	SubDistrict string `json:"subDistrict,omitempty"`
	District    string `json:"district,omitempty"`
	Province    string `json:"province,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Country     string `json:"country,omitempty"`
}
