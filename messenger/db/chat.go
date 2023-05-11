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

	err = factory.client.Preload("RoomChat").Where("id = ? ", groupChat.ID).First(&res).Error
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

func (factory GORMFactory) CheckRoomChatUser(RoomChatID, UserID uint) (res structure.GroupChat, Error error) {
	err := factory.client.Where("room_chat_id = ? AND user_id = ?", RoomChatID, UserID).First(&res).Error
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

func (factory GORMFactory) GetRoomChatById(roomChatId uint) (res structure.RoomChat, Error error) {
	var data structure.RoomChat
	err := factory.client.Where("id = ? ", roomChatId).First(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}

	res = data
	return
}

func (factory GORMFactory) PutRoomChat(groupChat structure.RoomChat) (Error error) {
	if groupChat.ID != 0 {
		err := factory.client.Model(&groupChat).Where("id = ?", groupChat.ID).Updates(
			structure.RoomChat{
				Name:      groupChat.Name,
				DeletedBy: groupChat.DeletedBy,
			}).Error

		if err != nil {
			Error = err
		}
	}

	return
}

func (factory GORMFactory) DeleteRoomChatById(roomChatId uint) (Error error) {
	var data structure.RoomChat
	err := factory.client.Where("id = ? ", roomChatId).Delete(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		} else {
			Error = errors.New("record not found")
			return
		}
	}
	var data1 structure.GroupChat
	err = factory.client.Where("room_chat_id = ? ", roomChatId).Delete(&data1).Error
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

func (factory GORMFactory) GetRoomChatListByUserId(UserID uint) (res []structure.GroupChat, Error error) {
	err := factory.client.Preload("RoomChat").Where("user_id = ?", UserID).Order("created_at DESC").Find(&res).Error
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

func (factory GORMFactory) GetMembersRoomChat(RoomChatID uint) (res []structure.GroupChat, Error error) {
	err := factory.client.Preload("RoomChat").Where("room_chat_id = ?", RoomChatID).Order("created_at DESC").Find(&res).Error
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

func (factory GORMFactory) GetMessengerByRoomChatId(roomChatId uint) (res []structure.Message, Error error) {
	err := factory.client.Where("room_chat_id = ?", roomChatId).Order("created_at desc").Find(&res).Error
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

func (factory GORMFactory) GetMessage(roomChatId uint) (res []structure.Message, Error error) {
	err := factory.client.Where("room_chat_id = ?", roomChatId).Order("created_at asc").Find(&res).Error
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

func (factory GORMFactory) PostChat(message structure.Message) (Error error) {
	err := factory.client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&message).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (factory GORMFactory) PutChat(message structure.Message) (Error error) {
	if message.ID != 0 {
		err := factory.client.Model(&message).Where("id = ?", message.ID).Updates(
			structure.Message{
				RoomChatID:   message.RoomChatID,
				Message:      message.Message,
				Image:        message.Image,
				SenderUserId: message.SenderUserId,
				ReadingDate:  message.ReadingDate,
				DeletedBy:    message.DeletedBy,
			}).Error

		if err != nil {
			Error = err
		}
	}

	return
}

func (factory GORMFactory) DeleteChat(messageId uint) (Error error) {
	var data structure.Message
	err := factory.client.Where("id = ?", messageId).Delete(&data).Error
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

// GetAllForAdminChatList todo admin
func (factory GORMFactory) GetAllForAdminChatList() (res []structure.GroupChat, Error error) {
	err := factory.client.Preload("RoomChat").Order("created_at DESC").Find(&res).Error
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
