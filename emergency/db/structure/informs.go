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
	OpsID               uint
	Status              string
}

type InformImage struct {
	gorm.Model
	InformID uint
	Inform   Inform
	Image    string
}
