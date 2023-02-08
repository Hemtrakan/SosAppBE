package control

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	rdbmsstructure "github.com/chiraponkub/DPU-SosApp-v.1.git/db/structure"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/role/request"
	response "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/role/response"
	"strings"
)

func (ctrl ConController) AddRoleCon(req *request.AddRole) (Error error) {
	var newReq rdbmsstructure.Role
	newReq.Name = strings.ToLower(req.Name)

	res, err := ctrl.Access.RDBMS.GetRoleDBByName(newReq)
	if res.Name == req.Name {
		Error = errors.New("มี Role นี้ในระบบแล้ว")
		return
	}

	role := rdbmsstructure.Role{
		Name: req.Name,
	}
	err = ctrl.Access.RDBMS.AddRoleDB(role)
	if err != nil {
		Error = err
		return
	}
	return
}

func (ctrl ConController) GetRoleListCon() (res response.ResponseMain, Error error) {
	data, err := ctrl.Access.RDBMS.GetRoleListDB()
	if err != nil {
		Error = err
		return
	}
	var resp []response.GetRoleList
	for _, m1 := range data {
		arr := response.GetRoleList{
			Name: m1.Name,
		}
		resp = append(resp, arr)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.GetRoleList = resp
	return
}
