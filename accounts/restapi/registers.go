package restapi

import (
	"accounts/constant"
	singup "accounts/restapi/model/singup/request"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"accounts/utility/token"
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
	var values = token.ValueKey{}
	if c.Request().Header.Get("Authorization") != "" {
		values = token.GetValuesToken(c)
	}

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

	_, err = ctrl.Ctx.PostUser(request, values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) ImageVerifyAgain(c echo.Context) error {
	var request = new(singup.UpdateImageVerifyAgain)
	var res response.RespMag
	APIName := "ImageVerify"
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

	errArr := ctrl.Ctx.ImageVerifyAgain(request)
	if err != nil {
		errRes := ""
		for _, m1 := range errArr {
			errRes = errRes + m1.Error() + " | "
		}
		res.Code = constant.ErrorCode
		res.Msg = errRes
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
