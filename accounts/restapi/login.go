package restapi

import (
	"accounts/constant"
	singin "accounts/restapi/model/singin/request"
	singinResp "accounts/restapi/model/singin/response"
	"accounts/utility/response"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) SignIn(c echo.Context) error {
	var request = new(singin.Login)
	var res response.RespMag

	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.Ctx.LoginLogic(request)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		return response.EchoSucceed(c, res)
	}

	if token == "" {
		res.Msg = errors.New("ชื่อผู้ใช้งานหรือรหัส่ผ่านผิด").Error()
		res.Code = constant.ErrorCode
		return response.EchoSucceed(c, res)
	}

	resp := singinResp.TokenRes{
		Token: token,
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, resp)
}
