package restapi

import (
	"accounts/constant"
	singup "accounts/restapi/model/singup/request"
	"accounts/utility/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) SendOTP(c echo.Context) error {
	var request = new(singup.PhoneNumber)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		return response.EchoError(c, 400, err)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	resp, err := ctrl.Ctx.SentOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) VerifyOTP(c echo.Context) error {
	var request = new(singup.OTP)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	err = ctrl.Ctx.VerifyOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "VerifySuccess"
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) CreateUser(c echo.Context) error {
	var request = new(singup.SingUp)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, err)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	err = ctrl.Ctx.PostUser(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "Succeed"
	return response.EchoSucceed(c, res)

}
