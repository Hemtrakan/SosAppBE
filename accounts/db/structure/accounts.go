package structure

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	gorm.Model
	PhoneNumber  string `gorm:"size:10;unique"`
	Password     string
	Firstname    string
	Lastname     string
	Email        string
	Birthday     time.Time
	Gender       string
	ImageProfile *string
	Workplace    *string
	IDCard       IDCard
	IDCardID     uint
	Address      Address
	AddressID    uint
	Role         Role
	RoleID       uint
	UpdateBy     *uint
	DeletedBy    *uint
}

type OTP struct {
	gorm.Model
	PhoneNumber string
	Key         string
	VerifyCode  string
	Expired     time.Time
	Active      bool
}

// todo ต้องทำ Log ในการ Login
type LogLogin struct {
	gorm.Model
	UserID      uint
	System      string
	IP          string
	Status      string
	Description string
}

type Role struct {
	gorm.Model
	Name      string
	UpdateBy  *uint
	DeletedBy *uint
}

type Address struct {
	gorm.Model
	Address     string
	SubDistrict string
	District    string
	Province    string
	PostalCode  string
	Country     string
	UpdateBy    *uint
	DeletedBy   *uint
}

type IDCard struct {
	gorm.Model
	TextIDCard  string `gorm:"size:13;unique"`
	PathImage   string
	Verify      bool
	Description string
	UpdateBy    *uint
	DeletedBy   *uint
}
