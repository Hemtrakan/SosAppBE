package db

import (
	"accounts/db/structure"
)

import (
	"errors"

	"gorm.io/gorm"
)

func (factory GORMFactory) SearchUser(value string, id uint) (response []*structure.Users, Error error) {
	var data []*structure.Users
	value = "%" + value + "%"
	err := factory.client.Preload("Role").Preload("IDCard").Preload("Address").
		Where("firstname LIKE ? OR lastname LIKE ? OR phone_number LIKE ? OR workplace LIKE ? ", value, value, value, value).
		Where("role_id <> ?", 1).
		Where("id <> ?", id).
		Find(&data).Error
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

func (factory GORMFactory) GetUserByPhone(req structure.Users) (response *structure.Users, Error error) {
	var data = new(structure.Users)
	err := factory.client.Preload("Role").Preload("IDCard").Preload("Address").Where("phone_number = ?", req.PhoneNumber).Find(&data).Error
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

func (factory GORMFactory) GetUserByID(req structure.Users) (response *structure.Users, Error error) {
	var data = new(structure.Users)
	err := factory.client.Preload("Role").Preload("IDCard").Preload("Address").Where("id = ?", req.ID).First(&data).Error
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

func (factory GORMFactory) GetUserList() (response []*structure.Users, Error error) {
	var data []*structure.Users
	err := factory.client.Preload("Role").Preload("IDCard").Preload("Address").Find(&data).Error
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

func (factory GORMFactory) PostUser(req structure.Users) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) PutUser(user *structure.Users, address *structure.Address, idCard *structure.IDCard) (Error []error) {

	if user.ID != 0 {
		err := factory.client.Model(&user).Where("id = ?", user.ID).Updates(
			structure.Users{
				Firstname:    user.Firstname,
				Lastname:     user.Lastname,
				Email:        user.Email,
				Birthday:     user.Birthday,
				Gender:       user.Gender,
				ImageProfile: user.ImageProfile,
				Workplace:    user.Workplace,
				UpdateBy:     &user.ID,
			}).Error

		if err != nil {
			Error = append(Error, err)
		}
	}

	if address.ID != 0 {
		err := factory.client.Model(&address).Where("id = ?", address.ID).Updates(
			structure.Address{
				Address:     address.Address,
				SubDistrict: address.SubDistrict,
				District:    address.District,
				Province:    address.Province,
				PostalCode:  address.PostalCode,
				Country:     address.Country,
				UpdateBy:    &user.ID,
			}).Error

		if err != nil {
			Error = append(Error, err)
		}
	}

	if idCard.ID != 0 {
		err := factory.client.Model(&idCard).Where("id = ?", idCard.ID).Updates(
			structure.IDCard{
				TextIDCard:  idCard.TextIDCard,
				PathImage:   idCard.PathImage,
				Verify:      idCard.Verify,
				Description: idCard.Description,
				UpdateBy:    &user.ID,
			}).Error

		if err != nil {
			Error = append(Error, err)
		}
	}

	return
}

func (factory GORMFactory) ChangePassword(req *structure.Users) (Error error) {
	var user structure.Users
	err := factory.client.Model(&user).Where("id = ?", req.ID).Updates(
		structure.Users{
			Password: req.Password,
		}).Error
	if err != nil {
		Error = err
	}
	return
}

func (factory GORMFactory) DeleteUser(req structure.Users) (Error error) {
	var data structure.Users
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
