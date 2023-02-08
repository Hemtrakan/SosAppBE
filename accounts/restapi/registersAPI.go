package restapi

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/constant"
	singup "github.com/chiraponkub/DPU-SosApp-v.1.git/restapi/model/singup/request"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/logs"
	"github.com/chiraponkub/DPU-SosApp-v.1.git/utility/response"
	"github.com/labstack/echo/v4"
)

func (ctrl Controller) SendOTP(c echo.Context) error {
	logs.LogStart("SendOTP")
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
		logs.LogError(err)
		return response.EchoError(c, 400, res)
	}
	logs.LogRequest(request)

	resp, err := ctrl.Ctx.SentOTPLogic(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		logs.LogError(err)
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = resp
	logs.LogResponse(res)
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
	var request = new(singup.Account)
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

	err = ctrl.Ctx.CreateUserLogin(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, 400, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "Succeed"
	return response.EchoSucceed(c, res)

}
