package model

import "time"

type TypeReq struct {
	NameType  string `json:"nameType"`
	ImageType string `json:"imageType"`
}

type TypeRes struct {
	ID         string       `json:"id,omitempty"`
	CreatedAt  time.Time    `json:"createdAt,omitempty"`
	UpdatedAt  time.Time    `json:"updatedAt,omitempty"`
	NameType   string       `json:"nameType,omitempty"`
	ImageType  string       `json:"imageType,omitempty"`
	DeletedBy  string       `json:"deletedBy,omitempty"`
	SubTypeRes []SubTypeRes `json:"subTypeRes,omitempty"`
}
