package inform

type InformResponse struct {
	ID string `json:"id,omitempty"`
	// Description คำอธิบายเพิ่มเต็ม
	Description string `json:"description,omitempty"`
	// Image รูปภาพที่เก็บสำหรับการแจ้งเหตุ เป็น Base64
	Image []ImageInfo `json:"image,omitempty"`
	// PhoneNumberCallBack เบอร์โทรติดต่อกลับ
	PhoneNumberCallBack string `json:"phoneNumberCallBack,omitempty"`
	// Latitude
	Latitude string `json:"latitude,omitempty"`
	// Longitude
	Longitude string `json:"longitude,omitempty"`

	// UserName
	UserName string `json:"username,omitempty"`
	// Workplace
	Workplace string `json:"workplace,omitempty"`

	// SubTypeID ประเภทของการแจ้งเหตุ
	SubTypeName string `json:"subTypeName,omitempty"`
	// Date วันเวลาที่แจ้ง
	Date string `json:"date,omitempty"`

	// Status สถานะการแจ้งเหตู
	Status string `json:"status,omitempty"`
}

type ImageInfo struct {
	ImageId string
	Image   string
}
