package query

type InformInfoList struct {
	ID              *string
	InformCreatedAt *string
	InformUpdateAt  *string
	InformDeletedAt *string
	UserInformID    *string
	Description     *string
	CALLBack        *string
	Latitude        *string
	Longitude       *string
	SubTypeId       *string
	SubTypeName     *string
	TypeID          *string
	Type            *string
	Status          *string
	StatusChat      *bool
	// UserNotiID  ไอดีของผู้รับแจ้งเหตุ
	UserNotiID *string
}

type InformInfoById struct {
	ID              *string
	InformCreatedAt *string
	InformUpdateAt  *string
	InformDeletedAt *string
	UserInformID    *string
	Description     *string
	CALLBack        *string
	Latitude        *string
	Longitude       *string
	SubTypeId       *string
	SubTypeName     *string
	TypeID          *string
	Type            *string
	StatusChat      *bool
	Status          *string
	ImageInfo       []*ImageInfo
	// UserNotiID  ไอดีของผู้รับแจ้งเหตุ
	UserNotiID *string
}

type ImageInfo struct {
	ImageId *string
	Image   *string
}
