package responsedb

type InformInfoList struct {
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
	NotiID          *string
	NotiCreatedAt   *string
	// UserNotiID  ไอดีของผู้รับแจ้งเหตุ
	UserNotiID *string
	NotiDes    *string
	Status     *string
}

type InformInfoById struct {
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
	NotiID          *string
	NotiCreatedAt   *string
	// UserNotiID  ไอดีของผู้รับแจ้งเหตุ
	UserNotiID *string
	NotiDes    *string
	Status     *string
	ImageInfo  []*ImageInfo
}

type ImageInfo struct {
	ImageId *string
	Image   *string
}