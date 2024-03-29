package restapi

import (
	"accounts/constant"
	"accounts/restapi/model/role/request"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"accounts/utility/token"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ctrl Controller) AddRole(c echo.Context) error {
	var request = new(request.Role)
	var res response.RespMag
	APIName := "roleAdd"
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

	err = ctrl.Ctx.AddRole(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetRoleList(c echo.Context) error {
	var res response.RespMag
	APIName := "getRoleList"
	loggers.LogStart(APIName)

	responses, err := ctrl.Ctx.GetRoleList()
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetRoleById(c echo.Context) error {
	var res response.RespMag
	APIName := "getRoleById"
	loggers.LogStart(APIName)

	id := c.Param("id")
	responses, err := ctrl.Ctx.GetRoleById(id)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = responses
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) UpdateRole(c echo.Context) error {
	var request = new(request.Role)
	var res response.RespMag
	APIName := "updateRole"
	loggers.LogStart(APIName)

	strId := c.Param("id")
	userId := token.GetValuesToken(c)
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
	roleId, err := strconv.ParseUint(strId, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PutRole(request, uint(roleId), userId.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteRole(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteRole"
	loggers.LogStart(APIName)

	strId := c.Param("id")
	userId := token.GetValuesToken(c)

	roleId, err := strconv.ParseUint(strId, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteRole(uint(roleId), userId.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
