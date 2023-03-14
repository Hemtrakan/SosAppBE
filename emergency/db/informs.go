package db

import (
	"emergency/db/structure"
	"emergency/db/structure/responsedb"
	"errors"
	"gorm.io/gorm"
)

const getInfoInfo = `SELECT i.id   AS ID
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
     , ii.id                    AS ImageId
     , ii.image                 AS Image
     , inf.id                   AS NotiID
     , inf.created_at           AS NotiCreatedAt
     , inf.user_id              AS UserNotiID
     , inf.description          AS NotiDes
     , inf.status               as Status
FROM informs AS i
         LEFT JOIN inform_images ii ON i.id = ii.inform_id
         LEFT JOIN inform_notifications inf ON i.id = inf.inform_id
         INNER JOIN sub_types st on st.id = i.sub_type_id
         INNER JOIN types t on t.id = st.type_id
where i.user_id = ?`

func (factory GORMFactory) GetInformList(UserId uint) (response []*responsedb.InformInfo, Error error) {
	rows, err := factory.client.Raw(getInfoInfo, UserId).Rows()
	if err != nil {
		Error = err
		return
	}
	defer rows.Close()

	var dataArr []*responsedb.InformInfo
	for rows.Next() {
		var data = new(responsedb.InformInfo)
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
			&data.ImageId,
			&data.Image,
			&data.NotiID,
			&data.NotiCreatedAt,
			&data.UserNotiID,
			&data.NotiDes,
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
