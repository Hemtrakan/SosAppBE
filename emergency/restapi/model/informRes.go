package model

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

	// UserId ของคนแจ้งเหตุ
	UserId string `json:"userId,omitempty"`
	// UserName ของคนแจ้งเหตุ
	UserName string `json:"username,omitempty"`

	// UserId ของคนรับแจ้งเหตุ
	UserIdOps string `json:"userIdOps,omitempty"`
	// UserName ของคนรับแจ้งเหตุ
	UserNameOps string `json:"UserNameOps,omitempty"`

	//// PhoneNumber
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// Workplace
	Workplace string `json:"workplace,omitempty"`

	// SubTypeID ประเภทของการแจ้งเหตุ
	SubTypeName string `json:"subTypeName,omitempty"`
	// Date วันเวลาที่แจ้ง
	Date string `json:"date,omitempty"`

	// UpdateDate วันเวลาที่อัพเดท
	UpdateDate string `json:"updateDate,omitempty"`

	// DeletedAt วันเวลาที่ลบ
	DeletedAt string `json:"deletedAt,omitempty"`

	// Status สถานะการแจ้งเหตู
	Status string `json:"status,omitempty"`

	// StatusChat สถานะการสร้างห้องแชท
	StatusChat bool `json:"statusChat"`
}

type ImageInfo struct {
	ImageId string
	Image   string
}
