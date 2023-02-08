package db

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	"gorm.io/gorm"
)

func (factory GORMFactory) SendOTPDB(req structure.OTP) (Error error) {
	var data []structure.OTP
	db := factory.client.Model(&data).Where("phone_number = ?", req.PhoneNumber).Update("active", false).Error
	if db != nil {
		Error = db
		return
	}

	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) GetOTPDb(req structure.OTP) (response *structure.OTP, Error error) {
	var data = new(structure.OTP)
	err := factory.client.Where("phone_number = ? and active = ?", req.PhoneNumber, true).Find(&data).Error
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

func (factory GORMFactory) UpdateOTPDB(req structure.OTP) (Error error) {
	var data structure.OTP
	db := factory.client.Where("phone_number = ? and key = ? and verify_code = ? and active = ?", req.PhoneNumber, req.Key, req.VerifyCode, true).Take(&data).Error
	if db != nil {
		if !errors.Is(db, gorm.ErrRecordNotFound) {
			Error = db
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}

	data.Active = false
	db = factory.client.Save(&data).Error
	if db != nil {
		return db
	}
	return
}

func (factory GORMFactory) CreateUserDB(req structure.Users) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}
