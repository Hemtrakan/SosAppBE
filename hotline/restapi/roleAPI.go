package restapi

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/role/request"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) AddRole(c echo.Context) error {
	var request = new(request.AddRole)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, "Succeed")
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}
	err = ctrl.Ctx.AddRoleCon(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) GetRoleList(c echo.Context) error {
	responses, err := ctrl.Ctx.GetRoleListCon()
	if err != nil {
		return response.EchoError(c, 400, err.Error())
	}
	return response.EchoSucceed(c, responses)
}
