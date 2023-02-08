package structure

import (
	"gorm.io/gorm"
)

type Inform struct {
	gorm.Model
	Description         string
	PhoneNumberCallBack string
	UserID              uint
	DeletedBy           uint
	SubTypeID           uint
}

type InformImage struct {
	gorm.Model
	Image    string
	Inform   Inform
	InformID uint
}

type InformNotification struct {
	gorm.Model
	Inform      Inform
	InformID    uint
	UserID      uint
	Description string
	Status      string
	DeletedBy   uint
}
