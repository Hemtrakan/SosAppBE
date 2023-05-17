package model

import "time"

type SubTypeReq struct {
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
	TypeRes      *TypeRes  `json:"type,omitempty"`
}
