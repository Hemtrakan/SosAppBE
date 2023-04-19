package model

import (
	"gorm.io/gorm"
	"time"
)

type JoinChatRes struct {
	Mag      string   `json:"mag,omitempty"`
	Username []string `json:"username,omitempty"`
}

type GetChatList struct {
	RoomChatID string         `json:"roomChatID,omitempty"`
	RoomName   string         `json:"roomName,omitempty"`
	OwnerId    string         `json:"ownerId,omitempty"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAT  gorm.DeletedAt `json:"deletedAT"`
	DeleteBy   string         `json:"deleteBy,omitempty"`
}
