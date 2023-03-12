package restapi

import (
	"emergency/constant"
	"emergency/restapi/model/inform"
	"emergency/utility/response"
	"emergency/utility/token"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) GetInformList(c echo.Context) error {
	var res response.RespMag

	APIName := "getInformList"
	values := token.GetValuesToken(c)
	token := token.GetAuthToken(c)
	resp, err := ctrl.Ctx.GetInform(values.ID, token)
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

//
//func (ctrl Controller) GetInformById(c echo.Context) error {
//	var res response.RespMag
//	APIName := "getInformById"
//	id := c.Param("id")
//	responses, err := ctrl.Ctx.GetRoleById(id)
//	if err != nil {
//
//	}
//
//
//	res.Msg = constant.SuccessMsg
//	res.Code = constant.SuccessCode
//	res.Data = responses
//	return response.EchoSucceed(c, res, APIName)
//}

func (ctrl Controller) PostInform(c echo.Context) error {
	var request = new(inform.InformRequest)
	var res response.RespMag
	APIName := "postInform"

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

	err = ctrl.Ctx.PostInform(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "InformSuccess"
	return response.EchoSucceed(c, res, APIName)
}
