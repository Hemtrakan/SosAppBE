package structure

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	HotlineNumberID int
	UserID          uint
	Time            time.Time
}

type HotlineNumber struct {
	gorm.Model
	Number           string
	Description      string
	DeletedBy        uint
	UserIDLogUpdated uint
	History          History
	HistoryID        uint
}
