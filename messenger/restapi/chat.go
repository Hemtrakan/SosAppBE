package restapi

import (
	"github.com/labstack/echo/v4"
	"messenger/constant"
	"messenger/restapi/model/chat/request"
	"messenger/utility/loggers"
	"messenger/utility/response"
	"messenger/utility/token"
	"net/http"
)

//func (ctrl Controller) GetChat(c echo.Context) error {
//	var res response.RespMag
//	APIName := "Chat"
//	loggers.LogStart(APIName)
//
//	responses, err := ctrl.Ctx.Chat()
//	if err != nil {
//		res.Code = constant.ErrorCode
//		res.Msg = err.Error()
//		return response.EchoError(c, http.StatusBadRequest, res, APIName)
//	}
//
//	res.Msg = constant.SuccessMsg
//	res.Code = constant.SuccessCode
//	res.Data = responses
//	return response.EchoSucceed(c, res, APIName)
//}

func (ctrl Controller) RoomChat(c echo.Context) error {
	var res response.RespMag
	APIName := "RoomChat"
	loggers.LogStart(APIName)

	req := request.RoomChatReq{}

	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	values := token.GetValuesToken(c)
	userId := values.ID
	err = ctrl.Ctx.RoomChat(userId, req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) JoinChat(c echo.Context) error {
	var res response.RespMag
	APIName := "JoinChat"
	loggers.LogStart(APIName)

	var req []request.GroupChat

	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	values := token.GetValuesToken(c)
	userId := values.ID

	err = ctrl.Ctx.JoinChat(userId, req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
