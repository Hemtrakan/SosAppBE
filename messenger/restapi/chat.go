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
	request := new(request.RoomChatReq)
	var res response.RespMag
	APIName := "RoomChat"
	loggers.LogStart(APIName)

	err := c.Bind(&request)
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

	Token := token.GetAuthToken(c)
	values := token.GetValuesToken(c)
	userId := values.ID
	err = ctrl.Ctx.RoomChat(userId, *request, Token)
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

	var req request.GroupChat

	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	Token := token.GetAuthToken(c)
	resp, err := ctrl.Ctx.JoinChat(req, Token)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	if resp.Mag != "" {
		res.Code = constant.SuccessCode
		res.Msg = constant.SuccessMsg
		res.Data = resp
		return response.EchoSucceed(c, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
