package restapi

import (
	"emergency/constant"
	"emergency/restapi/model/inform"
	"emergency/utility/loggers"
	"emergency/utility/response"
	"emergency/utility/token"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ctrl Controller) GetInformOpsList(c echo.Context) error {
	var res response.RespMag
	APIName := "GetInformOpsList"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	authToken := token.GetAuthToken(c)

	resp, err := ctrl.Ctx.GetInformOps(values.ID, authToken, values.Role)
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

func (ctrl Controller) GetAllInformList(c echo.Context) error {
	var res response.RespMag
	APIName := "GetAllInformList"
	loggers.LogStart(APIName)

	authToken := token.GetAuthToken(c)

	resp, err := ctrl.Ctx.GetAllInformOps(authToken)
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

func (ctrl Controller) GetInformOpsById(c echo.Context) error {
	var res response.RespMag
	APIName := "GetInformOpsById"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	authToken := token.GetAuthToken(c)
	id := c.Param("id")

	responses, err := ctrl.Ctx.GetInformOpsById(id, authToken, values.Role)
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

func (ctrl Controller) UpdateInform(c echo.Context) error {
	var res response.RespMag
	APIName := "UpdateInform"
	loggers.LogStart(APIName)
	var request = new(inform.UpdateInformRequest)

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

	if request.Status > 4 {
		res.Code = constant.ErrorCode
		res.Msg = "invalid status"
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	strInform := c.Param("id")
	informId, err := strconv.ParseUint(strInform, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	values := token.GetValuesToken(c)
	authToken := token.GetAuthToken(c)

	err = ctrl.Ctx.UpdateInform(request, authToken, uint(informId), values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "UpdateSuccess"
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteInform(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteInform"
	loggers.LogStart(APIName)

	strInform := c.Param("id")
	informId, err := strconv.ParseUint(strInform, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	values := token.GetValuesToken(c)

	err = ctrl.Ctx.DeleteInform(values.ID, uint(informId))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	res.Data = "DeleteSuccess"
	return response.EchoSucceed(c, res, APIName)
}
