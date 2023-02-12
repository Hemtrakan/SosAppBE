package model

import "time"

type UserProfile struct {
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	Birthday    time.Time `json:"birthday,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	IDCard      string    `json:"idCard,omitempty"`
	RoleID      string    `json:"roleID,omitempty"`
	Email       string    `json:"email,omitempty"`
}
