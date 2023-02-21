package model

type ResponseMain struct {
	Code        string        `json:"code"`
	Msg         string        `json:"msg"`
	GetRoleList []GetRoleList `json:"getRoleList"`
}

type GetRoleList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
