package db

import (
	"emergency/db/structure"
	"emergency/db/structure/responsedb"
	"errors"
	"gorm.io/gorm"
)

const getInformInfo = `SELECT i.id AS ID
     , i.created_at             AS InformCreatedAt
     , i.user_id                AS UserInformID
     , i.description            AS Description
     , i.phone_number_call_back AS CALLBack
     , i.latitude               as Latitude
     , i.longitude              as longitude
     , st.id                    AS SubTypeId
     , st.name                  AS SubTypeName
     , t.id                     AS TypeID
     , t.name                   AS Type
     , i.ops_id                 AS UserNotiID
     , i.status                 as Status
FROM informs AS i
         INNER JOIN sub_types st ON st.id = i.sub_type_id
         INNER JOIN types t ON t.id = st.type_id
`

const getInformImage = `SELECT ii.id
     , ii.image
FROM inform_images ii
    INNER JOIN informs i ON i.id = ii.inform_id
WHERE  ii.inform_id =  ?`

func (factory GORMFactory) GetInformList(UserId uint) (response []*responsedb.InformInfoList, Error error) {
	sql := getInformInfo + "WHERE i.user_id = ? "
	rows, err := factory.client.Raw(sql, UserId).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var dataArr []*responsedb.InformInfoList
	for rows.Next() {
		var data = new(responsedb.InformInfoList)
		rows.Scan(
			&data.ID,
			&data.InformCreatedAt,
			&data.UserInformID,
			&data.Description,
			&data.CALLBack,
			&data.Latitude,
			&data.Longitude,
			&data.SubTypeId,
			&data.SubTypeName,
			&data.TypeID,
			&data.Type,
			&data.UserNotiID,
			&data.Status,
		)
		dataArr = append(dataArr, data)
	}

	response = dataArr
	return
}

func (factory GORMFactory) GetImageByInformId(informId uint) (response *responsedb.InformInfoById, Error error) {
	sql := getInformInfo + "WHERE i.id = ? "
	rows, err := factory.client.Raw(sql, informId).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var data = new(responsedb.InformInfoById)
	for rows.Next() {
		rows.Scan(
			&data.ID,
			&data.InformCreatedAt,
			&data.UserInformID,
			&data.Description,
			&data.CALLBack,
			&data.Latitude,
			&data.Longitude,
			&data.SubTypeId,
			&data.SubTypeName,
			&data.TypeID,
			&data.Type,
			&data.UserNotiID,
			&data.Status,
		)

		var imageInfoArr []*responsedb.ImageInfo
		getInformImageRow, _ := factory.client.Raw(getInformImage, data.ID).Rows()

		for getInformImageRow.Next() {
			var imageInfo = new(responsedb.ImageInfo)
			getInformImageRow.Scan(
				&imageInfo.ImageId,
				&imageInfo.Image,
			)
			imageInfoArr = append(imageInfoArr, imageInfo)
		}
		data.ImageInfo = imageInfoArr
	}

	response = data
	return
}

func (factory GORMFactory) GetAllInformListForAdmin() (response []*responsedb.InformInfoList, Error error) {
	sql := getInformInfo
	rows, err := factory.client.Raw(sql).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var dataArr []*responsedb.InformInfoList
	for rows.Next() {
		var data = new(responsedb.InformInfoList)
		rows.Scan(
			&data.ID,
			&data.InformCreatedAt,
			&data.UserInformID,
			&data.Description,
			&data.CALLBack,
			&data.Latitude,
			&data.Longitude,
			&data.SubTypeId,
			&data.SubTypeName,
			&data.TypeID,
			&data.Type,
			&data.UserNotiID,
			&data.Status,
		)
		dataArr = append(dataArr, data)
	}

	response = dataArr
	return
}

func (factory GORMFactory) GetAllInformList() (response []*responsedb.InformInfoList, Error error) {
	sql := getInformInfo + "WHERE i.ops_id = 0"

	rows, err := factory.client.Raw(sql).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var dataArr []*responsedb.InformInfoList
	for rows.Next() {
		var data = new(responsedb.InformInfoList)
		rows.Scan(
			&data.ID,
			&data.InformCreatedAt,
			&data.UserInformID,
			&data.Description,
			&data.CALLBack,
			&data.Latitude,
			&data.Longitude,
			&data.SubTypeId,
			&data.SubTypeName,
			&data.TypeID,
			&data.Type,
			&data.UserNotiID,
			&data.Status,
		)
		dataArr = append(dataArr, data)
	}

	response = dataArr
	return
}

func (factory GORMFactory) GetInformListByOpsId(OpsId uint) (response []*responsedb.InformInfoList, Error error) {
	sql := getInformInfo + "WHERE i.ops_id = ? "
	rows, err := factory.client.Raw(sql, OpsId).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var dataArr []*responsedb.InformInfoList
	for rows.Next() {
		var data = new(responsedb.InformInfoList)
		rows.Scan(
			&data.ID,
			&data.InformCreatedAt,
			&data.UserInformID,
			&data.Description,
			&data.CALLBack,
			&data.Latitude,
			&data.Longitude,
			&data.SubTypeId,
			&data.SubTypeName,
			&data.TypeID,
			&data.Type,
			&data.UserNotiID,
			&data.Status,
		)
		dataArr = append(dataArr, data)
	}

	response = dataArr
	return
}

func (factory GORMFactory) PostInform(imageArr []structure.InformImage, inform structure.Inform) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&inform).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		} else {
			Error = err
			return
		}
	}

	for _, m1 := range imageArr {
		image := structure.InformImage{
			Model: gorm.Model{
				CreatedAt: m1.CreatedAt,
				UpdatedAt: m1.UpdatedAt,
			},
			InformID: inform.ID,
			Image:    m1.Image,
		}
		err = factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&image).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRegistered) {
				Error = err
				return
			} else {
				Error = err
				return
			}
		}
	}

	return
}

func (factory GORMFactory) PutInform(informID structure.Inform) (Error error) {
	err := factory.client.Model(&informID).Where("id = ?", informID.ID).Updates(
		structure.Inform{
			Description:         informID.Description,
			PhoneNumberCallBack: informID.PhoneNumberCallBack,
			Latitude:            informID.Latitude,
			Longitude:           informID.Longitude,
			DeletedBy:           informID.DeletedBy,
			SubTypeID:           informID.SubTypeID,
			OpsID:               informID.OpsID,
			Status:              informID.Status,
		}).Error

	if err != nil {
		Error = err
	}
	return
}

func (factory GORMFactory) DeleteInform(informID uint) (Error error) {
	err := factory.client.Where("id = ?", informID).Delete(&informID).Error
	if err != nil {
		Error = err
	}
	return
}
