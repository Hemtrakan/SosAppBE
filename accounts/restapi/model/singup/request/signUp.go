package request

type PhoneNumber struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type OTP struct {
	OTP         string `json:"otp" validate:"required"`
	VerifyCode  string `json:"verifyCode" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type SingUp struct {
	PhoneNumber     string  `json:"phoneNumber,omitempty" validate:"required"`
	Password        string  `json:"password"  validate:"required"`
	ConfirmPassword string  `json:"confirmPassword"  validate:"required"`
	FirstName       string  `json:"firstName,omitempty" validate:"required"`
	LastName        string  `json:"lastName,omitempty" validate:"required"`
	Email           string  `json:"email,omitempty"`
	Birthday        string  `json:"birthday,omitempty" validate:"required"`
	Gender          string  `json:"gender,omitempty" validate:"required"`
	ImageProfile    string  `json:"imageProfile"`
	Workplace       string  `json:"workplace"`
	RoleId          uint    `json:"role"  validate:"required"`
	IDCard          IDCard  `json:"idCard,omitempty" validate:"required"`
	Address         Address `json:"address" validate:"required"`
	Verify          Verify  `json:"verify"`
}

type Address struct {
	Address     string `json:"address,omitempty" validate:"required"`
	SubDistrict string `json:"subDistrict,omitempty" validate:"required"`
	District    string `json:"district,omitempty" validate:"required"`
	Province    string `json:"province,omitempty" validate:"required"`
	PostalCode  string `json:"postalCode,omitempty" validate:"required"`
	Country     string `json:"country,omitempty" validate:"required"`
}

type IDCard struct {
	TextIDCard string `json:"textIDCard" validate:"required"`
	PathImage  string `json:"pathImage" validate:"required"`
}

type Verify struct {
	OTP        string `json:"otp,omitempty" validate:"required"`
	VerifyCode string `json:"verifyCode,omitempty" validate:"required"`
}

type UpdateImageVerifyAgain struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Password    string `json:"password"  validate:"required"`
	TextIDCard  string `json:"textIDCard" validate:"required"`
	PathImage   string `json:"pathImage" validate:"required"`
}
