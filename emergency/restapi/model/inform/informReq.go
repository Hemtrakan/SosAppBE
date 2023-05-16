package inform

type InformRequest struct {
	// Description คำอธิบายเพิ่มเต็ม
	Description string `json:"description,omitempty"`

	// Image รูปภาพที่เก็บสำหรับการแจ้งเหตุ
	Images []string `json:"images,omitempty"`

	// PhoneNumberCallBack เบอร์โทรติดต่อกลับ
	PhoneNumberCallBack string `json:"phoneNumberCallBack,omitempty"`

	// Latitude
	Latitude string `json:"latitude,omitempty"`

	// Longitude
	Longitude string `json:"longitude,omitempty"`

	// UserID ไอดีของผู้แจ้งเหตุ
	UserID string `json:"userID,omitempty"`

	// SubTypeID ประเภทของการแจ้งเหตุ
	SubTypeID uint `json:"subTypeID,omitempty"`
}

type UpdateInformRequest struct {
	Description         *string `json:"description,omitempty"`
	PhoneNumberCallBack *string `json:"phoneNumberCallBack,omitempty"`
	Latitude            *string `json:"latitude,omitempty"`
	Longitude           *string `json:"longitude,omitempty"`
	UserID              *uint   `json:"userID,omitempty"`
	DeletedBy           *uint   `json:"deletedBy,omitempty"`
	SubTypeID           *uint   `json:"subTypeID,omitempty"`
	OpsID               *uint   `json:"opsID,omitempty"`
	Status              *int    `json:"status,omitempty"`
	StatusChat          *bool   `json:"statusChat,omitempty"`
}
