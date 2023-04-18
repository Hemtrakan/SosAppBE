package db

import (
	"errors"
	"gorm.io/gorm"
	"messenger/db/structure"
)

func (factory GORMFactory) RoomChat(groupChat structure.GroupChat) (res structure.GroupChat, Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&groupChat).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}

	err = factory.client.Preload("RoomChat").Where("id = ? ", groupChat.ID).Find(&res).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}

	return
}

func (factory GORMFactory) JoinChat(groupChat structure.GroupChat) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&groupChat).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}
func (factory GORMFactory) CheckRoomChatForUser(RoomChatID, UserID uint) (res structure.GroupChat, Error error) {
	err := factory.client.Where("room_chat_id = ? AND user_id = ?", RoomChatID, UserID).Find(&res).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	return
}
