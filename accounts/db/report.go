package db

import (
	"accounts/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetLogLogin() (res []structure.LogLogin, Error error) {
	var data []structure.LogLogin
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
	res = data
	return
}
