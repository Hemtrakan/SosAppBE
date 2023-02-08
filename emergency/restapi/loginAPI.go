package restapi

import (
	"errors"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	singin "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singin/request"
	singinResp "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singin/response"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/logs"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) SignInUser(c echo.Context) error {
	logs.LogStart("SignInUser")
	var request = new(singin.Login)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, err)
	}
	logs.LogRequest(request)
	token, err := ctrl.Ctx.LoginLogic(request)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		logs.LogError(err)
		return response.EchoSucceed(c, res)
	}
	if token == "" {
		res.Msg = errors.New("ชื่อผู้ใช้งานหรือรหัส่ผ่านผิด").Error()
		res.Code = constant.ErrorCode
		logs.LogResponse(res)
		return response.EchoSucceed(c, res)
	}

	resp := singinResp.TokenRes{
		Token: token,
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	logs.LogResponse(res)
	return response.EchoSucceed(c, resp)
}

func (ctrl Controller) SignUpAdmin(c echo.Context) error {

	return response.EchoSucceed(c, "")
}
