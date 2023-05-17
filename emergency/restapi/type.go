package restapi

import (
	"emergency/constant"
	"emergency/restapi/model"
	"emergency/utility/loggers"
	"emergency/utility/response"
	"emergency/utility/token"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ctrl Controller) GetType(c echo.Context) error {
	var res response.RespMag
	APIName := "GetType"
	loggers.LogStart(APIName)
	resp, err := ctrl.Ctx.GetType()
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

func (ctrl Controller) GetTypeById(c echo.Context) error {
	var res response.RespMag
	APIName := "GetTypeById"
	loggers.LogStart(APIName)
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	resp, err := ctrl.Ctx.GetTypeById(uint(id))
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

func (ctrl Controller) PostType(c echo.Context) error {
	var res response.RespMag
	APIName := "PostType"
	loggers.LogStart(APIName)

	var request = model.TypeReq{}
	err := c.Bind(&request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PostType(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) PutType(c echo.Context) error {
	var res response.RespMag
	APIName := "PutType"
	loggers.LogStart(APIName)

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	var request = model.TypeReq{}
	err = c.Bind(&request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PutType(uint(id), request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteType(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteType"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteType(uint(id), values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
