package db

import (
	"accounts/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) AddRole(req structure.Role) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) GetRoleList() (response []structure.Role, Error error) {
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

func (factory GORMFactory) GetRoleByName(req structure.Role) (response structure.Role, Error error) {
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

func (factory GORMFactory) GetRoleById(req structure.Role) (response structure.Role, Error error) {
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

func (factory GORMFactory) PutRole(req structure.Role) (Error error) {
	var data structure.Role
	err := factory.client.Model(&data).Where("id = ?", req.ID).Updates(
		structure.Role{
			Name:     req.Name,
			UpdateBy: req.UpdateBy,
		}).Error

	if err != nil {
		Error = err
	}

	return
}

func (factory GORMFactory) DeleteRole(req structure.Role) (Error error) {
	var data structure.Role
	err := factory.client.Where("id = ?", req.ID).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
		} else {
			Error = errors.New("record not found")
			return
		}
		return
	}
	return
}
