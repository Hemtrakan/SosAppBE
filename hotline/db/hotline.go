package db

import (
	"errors"
	"gorm.io/gorm"
	"hotline/db/structure"
)

func (factory GORMFactory) GetHotLineById(id uint) (response structure.HotlineNumber, Error error) {
	var data structure.HotlineNumber
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

func (factory GORMFactory) PostHotLine(req structure.HotlineNumber) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) PutHotLine(req structure.HotlineNumber) (Error error) {
	var data structure.HotlineNumber
	err := factory.client.Model(&data).Where("id = ?", req.ID).Updates(
		structure.HotlineNumber{
			Number:           req.Number,
			Description:      req.Description,
			DeletedBy:        req.DeletedBy,
			UserIDLogUpdated: req.UserIDLogUpdated,
		}).Error

	if err != nil {
		Error = err
	}

	return
}

func (factory GORMFactory) DeleteHotLine(req structure.HotlineNumber) (Error error) {
	var data structure.HotlineNumber
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

func (factory GORMFactory) GetHistory() (response []structure.History, Error error) {
	var data []structure.History
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

func (factory GORMFactory) PostHistory(req structure.History) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}
