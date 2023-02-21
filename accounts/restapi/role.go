package restapi

import (
	"accounts/constant"
	"accounts/restapi/model/role/request"
	"accounts/utility/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) AddRole(c echo.Context) error {
	var request = new(request.AddRole)
	var res response.RespMag
	err := c.Bind(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	err = ValidateStruct(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}
	err = ctrl.Ctx.AddRole(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) GetRoleList(c echo.Context) error {
	var res response.RespMag
	responses, err := ctrl.Ctx.GetRoleList()
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, err.Error())
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, responses)
}

func (ctrl Controller) GetRoleById(c echo.Context) error {
	var res response.RespMag
	id := c.Param("id")
	responses, err := ctrl.Ctx.GetRoleById(id)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, err.Error())
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, id)
}

func (ctrl Controller) UpdateRole(c echo.Context) error {
	var res response.RespMag
	id := c.Param("id")
	responses, err := ctrl.Ctx.GetRoleById(id)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, err.Error())
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, id)
}

func (ctrl Controller) DeleteRole(c echo.Context) error {
	var res response.RespMag
	id := c.Param("id")
	responses, err := ctrl.Ctx.GetRoleById(id)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, err.Error())
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, id)
}
