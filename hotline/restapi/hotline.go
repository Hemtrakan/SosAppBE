package restapi

import (
	"github.com/labstack/echo/v4"
	"hotline/constant"
	"hotline/utility/response"
	"net/http"
)

func (ctrl Controller) GetHotLine(c echo.Context) error {
	var res response.RespMag
	APIName := "getHotLine"
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
