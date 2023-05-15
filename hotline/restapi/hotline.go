package restapi

import (
	"github.com/labstack/echo/v4"
	"hotline/constant"
	req "hotline/restapi/model/hotline"
	"hotline/utility/loggers"
	"hotline/utility/response"
	"hotline/utility/token"
	"net/http"
	"strconv"
)

func (ctrl Controller) GetHistory(c echo.Context) error {
	var res response.RespMag
	APIName := "GetHistory"
	loggers.LogStart(APIName)

	responses, err := ctrl.Ctx.GetHistory()
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

func (ctrl Controller) PostHistory(c echo.Context) error {
	var res response.RespMag
	APIName := "PostHistory"
	loggers.LogStart(APIName)

	var req = new(req.HistoryReq)
	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PostHistory(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetHotLine(c echo.Context) error {
	var res response.RespMag
	APIName := "getHotLine"
	loggers.LogStart(APIName)

	responses, err := ctrl.Ctx.GetHotLine()
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

func (ctrl Controller) GetHotLineById(c echo.Context) error {
	var res response.RespMag
	APIName := "GetHotLineById"
	loggers.LogStart(APIName)

	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	responses, err := ctrl.Ctx.GetHotLineById(uint(id))
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

func (ctrl Controller) PostHotLine(c echo.Context) error {
	var res response.RespMag
	APIName := "PostHotLine"
	loggers.LogStart(APIName)

	var req = new(req.HotlineReq)
	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	values := token.GetValuesToken(c)
	err = ctrl.Ctx.PostHotLine(req, values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) PutHotLine(c echo.Context) error {
	var res response.RespMag
	APIName := "PutHotLine"
	loggers.LogStart(APIName)

	var req = new(req.HotlineReq)
	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ValidateStruct(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	values := token.GetValuesToken(c)
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.PutHotLine(req, uint(id), values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteHotLine(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteHotLine"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteHotLine(uint(id), values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
