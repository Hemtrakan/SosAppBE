package request

type RoomChatReq struct {
	RoomName  string    `json:"roomName" validate:"required"`
	GroupChat GroupChat `json:"groupChat"`
}

type GroupChat struct {
	RoomChatID uint   `json:"roomChatID,omitempty"`
	UserID     []uint `json:"userID,omitempty"`
}
