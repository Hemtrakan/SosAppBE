package db

import (
	"emergency/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) GetSubType() (res []structure.SubType, Error error) {
	err := factory.client.Preload("Type").Find(&res).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = gorm.ErrRecordNotFound
			return
		}
	}
	return
}

func (factory GORMFactory) GetSubTypeByTypeId(id uint) (res []structure.SubType, Error error) {
	var data []structure.SubType
	err := factory.client.Where("type_id = ?", id).Find(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = gorm.ErrRecordNotFound
			return
		}
	}
	res = data
	return
}

func (factory GORMFactory) GetSubTypeById(id uint) (res structure.SubType, Error error) {
	var data structure.SubType
	err := factory.client.Preload("Type").Where("id = ?", id).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = gorm.ErrRecordNotFound
			return
		}
	}
	res = data
	return
}

func (factory GORMFactory) PostSubType(SubTypes structure.SubType) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&SubTypes).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) PutSubType(SubTypes structure.SubType) (Error error) {
	err := factory.client.Where("id = ?", SubTypes.ID).Updates(SubTypes).Error
	if err != nil {
		Error = err
	}
	return
}

func (factory GORMFactory) DeleteSubType(id uint) (Error error) {
	var subType structure.SubType
	err := factory.client.Where("id = ?", id).Delete(&subType).Error
	if err != nil {
		Error = err
	}
	return
}
