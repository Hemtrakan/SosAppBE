package db

import (
	"emergency/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetType() (response []structure.Type, Error error) {
	var data []structure.Type
	err := factory.client.Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = gorm.ErrRecordNotFound
			return
		}
	}
	response = data
	return
}

func (factory GORMFactory) GetTypeById(id uint) (response structure.Type, Error error) {
	var data structure.Type
	err := factory.client.Where("id = ?", id).First(&data).Error
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

func (factory GORMFactory) PostType(types structure.Type) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&types).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) PutType(types structure.Type) (Error error) {
	err := factory.client.Where("id = ?", types.ID).Updates(types).Error
	if err != nil {
		Error = err
	}
	return
}

func (factory GORMFactory) DeleteType(id uint) (Error error) {
	var types structure.Type
	err := factory.client.Where("id = ?", id).Delete(&types).Error
	if err != nil {
		Error = err
	}
	return
}
