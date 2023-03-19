package restapi

import (
	"accounts/constant"
	singin "accounts/restapi/model/singin/request"
	singinResp "accounts/restapi/model/singin/response"
	singup "accounts/restapi/model/singup/request"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) SendOTP(c echo.Context) error {
	var request = new(singup.PhoneNumber)
	var res response.RespMag
	APIName := "sendOTP"
	loggers.LogStart(APIName)
	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	resp, err := ctrl.Ctx.SentOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) VerifyOTP(c echo.Context) error {
	var request = new(singup.OTP)
	var res response.RespMag
	APIName := "verifyOTP"
	loggers.LogStart(APIName)

	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.VerifyOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "VerifySuccess"
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) CreateUser(c echo.Context) error {
	var request = new(singup.SingUp)
	var res response.RespMag
	APIName := "createUser"
	loggers.LogStart(APIName)

	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	data, err := ctrl.Ctx.PostUser(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	var requestToken = new(singin.Login)
	ip := c.RealIP()
	system := c.Request().Header.Get("User-Agent")
	requestToken.Username = data.PhoneNumber
	requestToken.Password = data.Password

	token, err := ctrl.Ctx.LoginLogic(requestToken, ip, system)
	if err != nil {
		res.Msg = err.Error()
		res.Code = constant.ErrorCode
		return response.EchoError(c, http.StatusBadRequest, res, APIName)

	}

	resp := singinResp.TokenRes{
		Token: token,
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, res, APIName)
}
