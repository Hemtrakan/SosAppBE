package structure

import "gorm.io/gorm"

type GroupChat struct {
	gorm.Model
	UserID     uint
	RoomChat   RoomChat
	RoomChatID uint
}

type RoomChat struct {
	gorm.Model
	Name        string
	UserOwnerId uint
	DeletedBy   uint
	InformId    uint
}

type Message struct {
	gorm.Model
	RoomChat     RoomChat
	RoomChatID   uint
	Message      string
	Image        string
	SenderUserId uint
	ReadingDate  int // todo นับจำนวนการอ่านข้อความ
	DeletedBy    uint
}
