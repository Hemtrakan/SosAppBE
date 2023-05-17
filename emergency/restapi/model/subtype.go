package model

import "time"

type SubTypeReq struct {
	TypeId       string `json:"typeId"`
	NameSubType  string `json:"nameSubType"`
	ImageSubType string `json:"imageSubType"`
}

type SubTypeRes struct {
	ID           string    `json:"id,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	NameSubType  string    `json:"nameSubType,omitempty"`
	ImageSubType string    `json:"imageSubType,omitempty"`
	DeletedBy    string    `json:"deletedBy,omitempty"`
	TypeId       string    `json:"typeId,omitempty"`
}
