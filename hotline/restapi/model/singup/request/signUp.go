package request

type PhoneNumber struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type OTP struct {
	OTP         int    `json:"otp" validate:"required"`
	VerifyCode  string `json:"verifyCode" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type Account struct {
	PhoneNumber     string `json:"phoneNumber,omitempty" validate:"required"`
	Password        string `json:"password"  validate:"required"`
	ConfirmPassword string `json:"confirmPassword"  validate:"required"`
	FirstName       string `json:"firstName,omitempty" validate:"required"`
	LastName        string `json:"lastName,omitempty" validate:"required"`
	//Birthday        time.Time `json:"birthday,omitempty" validate:"required"`
	Gender     string `json:"gender,omitempty" validate:"required"`
	IDCard     string `json:"idCard,omitempty" validate:"required"`
	RoleID     uint   `json:"roleID,omitempty" validate:"required"`
	Email      string `json:"email,omitempty"`
	Key        int    `json:"key,omitempty" validate:"required"`
	VerifyCode string `json:"verifyCode,omitempty" validate:"required"`
}
