package control

import (
	"accounts/constant"
	rdbmsstructure "accounts/db/structure"
	"accounts/restapi/model/role/request"
	response "accounts/restapi/model/role/response"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func (ctrl Controller) AddRole(req *request.AddRole) (Error error) {
	var newReq rdbmsstructure.Role
	newReq.Name = strings.ToLower(req.Name)

	res, err := ctrl.Access.RDBMS.GetRoleByName(newReq)
	if res.Name == req.Name {
		Error = errors.New("มี Role นี้ในระบบแล้ว")
		return
	}

	role := rdbmsstructure.Role{
		Name: req.Name,
	}
	err = ctrl.Access.RDBMS.AddRole(role)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl Controller) GetRoleList() (res response.ResponseMain, Error error) {
	data, err := ctrl.Access.RDBMS.GetRoleList()
	if err != nil {
		Error = err
		return
	}
	var resp []response.GetRoleList
	for _, m1 := range data {
		arr := response.GetRoleList{
			Id:   strconv.Itoa(int(m1.ID)),
			Name: m1.Name,
		}
		resp = append(resp, arr)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.GetRoleList = resp
	return
}

func (ctrl Controller) GetRoleById(roleID string) (res response.ResponseMain, Error error) {
	ID, err := strconv.ParseUint(roleID, 10, 64)
	req := rdbmsstructure.Role{
		Model: gorm.Model{
			ID: uint(ID),
		},
	}
	data, err := ctrl.Access.RDBMS.GetRoleById(req)
	if err != nil {
		Error = err
		return
	}
	var resp []response.GetRoleList
	arr := response.GetRoleList{
		Id:   strconv.Itoa(int(data.ID)),
		Name: data.Name,
	}
	resp = append(resp, arr)

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.GetRoleList = resp
	return
}
