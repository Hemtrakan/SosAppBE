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
	ID             uint      `json:"id"`
	RoomChatID     uint      `json:"roomChatID"`
	RoomName       string    `json:"roomName"`
	Message        string    `json:"message"`
	Image          string    `json:"image"`
	SenderUserId   uint      `json:"senderUserId"`
	SenderUsername string    `json:"senderUsername"`
	ReadingDate    int       `json:"readingDate"` // todo นับจำนวนการอ่านข้อความ
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAT      time.Time `json:"deletedAT"`
	DeletedBy      uint      `json:"deletedBy"`
}

type GetMemberRoomChat struct {
	RoomChatID     string           `json:"roomChatID,omitempty"`
	RoomName       string           `json:"roomName,omitempty"`
	OwnerId        string           `json:"ownerId,omitempty"`
	MemberRoomChat []MemberRoomChat `json:"memberRoomChat,omitempty"`
}

type MemberRoomChat struct {
	UserId    uint   `json:"userId,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	RoleID    string `json:"roleID,omitempty"`
	RoleName  string `json:"roleName,omitempty"`
}
