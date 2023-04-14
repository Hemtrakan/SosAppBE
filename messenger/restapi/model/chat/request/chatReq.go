package request

type RoomChatReq struct {
	RoomName  string      `json:"roomName,omitempty"`
	GroupChat []GroupChat `json:"groupChat"`
}

type GroupChat struct {
	UserID uint `json:"userID,omitempty"`
}
