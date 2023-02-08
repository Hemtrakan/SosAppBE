package db

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	"gorm.io/gorm"
)

func (factory GORMFactory) AddRoleDB(req structure.Role) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) GetRoleListDB() (response []structure.Role, Error error) {
	var data []structure.Role
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

func (factory GORMFactory) GetRoleDBByName(req structure.Role) (response structure.Role, Error error) {
	var data structure.Role
	err := factory.client.Where("name = ?", req.Name).First(&data).Error
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

func (factory GORMFactory) GetRoleDBById(req structure.Role) (response structure.Role, Error error) {
	var data structure.Role
	err := factory.client.Where("id = ?", req.ID).First(&data).Error
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
