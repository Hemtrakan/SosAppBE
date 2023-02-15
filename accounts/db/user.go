package db

import (
	"accounts/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) LogLogin(req structure.Users) (Error error) {

	//Error =
	return
}

func (factory GORMFactory) GetUserByID(req structure.Users) (response *structure.Users, Error error) {
	var data = new(structure.Users)
	err := factory.client.Preload("Role").Preload("IDCard").Preload("Address").Where("id = ?", req.ID).Find(&data).Error
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
