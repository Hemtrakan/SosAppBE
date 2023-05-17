package request

type UserReq struct {
	FirstName    string   `json:"firstName,omitempty"`
	LastName     string   `json:"lastName,omitempty"`
	Email        string   `json:"email,omitempty"`
	Birthday     string   `json:"birthday,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	Workplace    string   `json:"workplace,omitempty"`
	ImageProfile string   `json:"imageProfile,omitempty"`
	IdCard       *IdCard  `json:"idCard,omitempty"`
	Address      *Address `json:"address,omitempty"`
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

type ChangePassword struct {
	OldPassword     string `json:"oldPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
}
