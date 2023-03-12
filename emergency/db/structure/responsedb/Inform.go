package responsedb

type InformInfo struct {
	ID              *string
	InformCreatedAt *string
	UserInformID    *string
	Description     *string
	CALLBack        *string
	Latitude        *string
	Longitude       *string
	SubTypeId       *string
	SubTypeName     *string
	TypeID          *string
	Type            *string
	ImageId         *string
	Image           *string
	NotiID          *string
	NotiCreatedAt   *string
	// UserNotiID  ไอดีของผู้รับแจ้งเหตุ
	UserNotiID *string
	NotiDes    *string
	Status     *string
}
