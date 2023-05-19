package restapi

import (
	"accounts/constant"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl Controller) GetLogLogin(c echo.Context) error {
	var res response.RespMag
	APIName := "GetLogLogin"
	loggers.LogStart(APIName)

	responses, err := ctrl.Ctx.GetLogLogin()
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
