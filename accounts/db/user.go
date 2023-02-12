package db

import (
	"accounts/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetUserByID(req structure.Users) (response *structure.Users, Error error) {
	var data = new(structure.Users)
	err := factory.client.Preload("Role").Preload("Address").Preload("IDCard").Where("id = ?", req.ID).Find(&data).Error
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
