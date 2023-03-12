package inform

type InformRequest struct {
	// Description คำอธิบายเพิ่มเต็ม
	Description string `json:"description,omitempty"`

	// Image รูปภาพที่เก็บสำหรับการแจ้งเหตุ
	Images string `json:"images,omitempty"`

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
