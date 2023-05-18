package model

import (
	"time"
)

type JoinChatRes struct {
	Mag      string   `json:"mag,omitempty"`
	Username []string `json:"username,omitempty"`
}

type GetChatList struct {
	RoomChatID string    `json:"roomChatID,omitempty"`
	RoomName   string    `json:"roomName,omitempty"`
	OwnerId    string    `json:"ownerId,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	DeletedAT  time.Time `json:"deletedAT,omitempty"`
	DeleteBy   string    `json:"deleteBy,omitempty"`
}

type GetChat struct {
	ID           uint      `json:"id,omitempty"`
	RoomChatID   uint      `json:"roomChatID,omitempty"`
	RoomName     string    `json:"roomName,omitempty"`
	Message      string    `json:"message,omitempty"`
	Image        string    `json:"image,omitempty"`
	SenderUserId uint      `json:"senderUserId,omitempty"`
	ReadingDate  int       `json:"readingDate,omitempty"` // todo นับจำนวนการอ่านข้อความ
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	DeletedAT    time.Time `json:"deletedAT,omitempty"`
	DeletedBy    uint      `json:"deletedBy,omitempty"`
}

type GetMemberRoomChat struct {
	RoomChatID     string           `json:"roomChatID,omitempty"`
	RoomName       string           `json:"roomName,omitempty"`
	OwnerId        string           `json:"ownerId,omitempty"`
	MemberRoomChat []MemberRoomChat `json:"memberRoomChat,omitempty"`
}

type MemberRoomChat struct {
	UserId uint `json:"userId,omitempty"`
}
