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

func (ctrl Controller) GetSubType(c echo.Context) error {
	var res response.RespMag
	APIName := "GetSubType"
	loggers.LogStart(APIName)

	resp, err := ctrl.Ctx.GetSubType()
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

func (ctrl Controller) GetSubTypeById(c echo.Context) error {
	var res response.RespMag
	APIName := "GetSubTypeById"
	loggers.LogStart(APIName)
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	resp, err := ctrl.Ctx.GetSubTypeById(uint(id))
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

func (ctrl Controller) PostSubType(c echo.Context) error {
	var res response.RespMag
	APIName := "PostSubType"
	loggers.LogStart(APIName)

	var request = model.SubTypeReq{}
	err := c.Bind(&request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PostSubType(request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) PutSubType(c echo.Context) error {
	var res response.RespMag
	APIName := "PutSubType"
	loggers.LogStart(APIName)

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	var request = model.SubTypeReq{}
	err = c.Bind(&request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PutSubType(uint(id), request)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteSubType(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteSubType"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteSubType(uint(id), values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
