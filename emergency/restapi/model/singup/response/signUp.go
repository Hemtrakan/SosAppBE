package model

type OTP struct {
	OTP        int    `json:"otp"`
	VerifyCode string `json:"verifyCode"`
}
