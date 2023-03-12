package structure

import (
	"gorm.io/gorm"
)

type Inform struct {
	gorm.Model
	Description         string
	PhoneNumberCallBack string
	Latitude            string
	Longitude           string
	UserID              uint
	DeletedBy           uint
	SubTypeID           uint
}

type InformImage struct {
	gorm.Model
	InformID uint
	Inform   Inform
	Image    string
}

type InformNotification struct {
	gorm.Model
	InformID uint
	Inform   Inform
	// UserID คนที่รับแจ้งเหตุ
	UserID      uint
	Description string
	Status      string
	DeletedBy   uint
}
