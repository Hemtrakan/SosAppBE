package db

import (
	"errors"
	"gorm.io/gorm"
	"hotline/db/structure"
)

func (factory GORMFactory) GetHotLine() (response []structure.HotlineNumber, Error error) {
	var data []structure.HotlineNumber
	err := factory.client.Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	response = data
	return
}
