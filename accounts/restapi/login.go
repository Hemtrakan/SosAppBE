package restapi

import (
	"accounts/constant"
	singin "accounts/restapi/model/singin/request"
	singinResp "accounts/restapi/model/singin/response"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) SignIn(c echo.Context) error {
	var request = new(singin.Login)
	var res response.RespMag
	APIName := "signIn"
	loggers.LogStart(APIName)

	err := c.Bind(request)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	ip := c.RealIP()
	system := c.Request().Header.Get("User-Agent")

	token, err := ctrl.Ctx.LoginLogic(request, ip, system)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	if token == "" {
		res.Msg = errors.New("ชื่อผู้ใช้งานหรือรหัส่ผ่านไม่ถูกต้อง").Error()
		res.Code = constant.ErrorCode
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	resp := singinResp.TokenRes{
		Token: token,
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, resp, APIName)
}
