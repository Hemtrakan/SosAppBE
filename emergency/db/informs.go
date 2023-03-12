package db

import (
	"emergency/db/structure"
	"errors"
	"gorm.io/gorm"
)

func (factory GORMFactory) PostInform(req structure.InformImage) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&req).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		} else {
			Error = err
			return
		}
	}
	return
}
