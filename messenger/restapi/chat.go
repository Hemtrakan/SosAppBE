package restapi

import (
	"github.com/labstack/echo/v4"
	"messenger/constant"
	"messenger/restapi/model/chat/request"
	"messenger/utility/loggers"
	"messenger/utility/response"
	"messenger/utility/token"
	"net/http"
	"strconv"
)

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

func (ctrl Controller) UpdateRoomChat(c echo.Context) error {
	request := new(request.RoomChatReq)
	var res response.RespMag
	APIName := "UpdateRoomChat"
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

	values := token.GetValuesToken(c)
	userId := values.ID

	roomChatIdStr := c.Param("roomId")
	roomChatID, err := strconv.ParseUint(roomChatIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.UpdateRoomChat(userId, uint(roomChatID), *request, values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteRoomChat(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteRoomChat"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	userId := values.ID

	roomChatIdStr := c.Param("roomId")
	roomChatID, err := strconv.ParseUint(roomChatIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteRoomChat(userId, uint(roomChatID), values.Role)
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

func (ctrl Controller) GetChatList(c echo.Context) error {
	var res response.RespMag
	APIName := "GetChatList"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	userId := values.ID

	resp, err := ctrl.Ctx.GetChatList(userId, values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	if len(resp) > 0 {
		res.Data = resp
		res.Total = len(resp)
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetMembersRoomChat(c echo.Context) error {
	var res response.RespMag
	APIName := "GetMembersRoomChat"
	loggers.LogStart(APIName)

	roomChatIdStr := c.Param("roomChatId")
	roomChatId, err := strconv.ParseUint(roomChatIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	values := token.GetValuesToken(c)
	token := token.GetAuthToken(c)

	resp, err := ctrl.Ctx.GetMembersRoomChat(uint(roomChatId), values.Role, token)
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

func (ctrl Controller) GetMessageByRoomChatId(c echo.Context) error {
	var res response.RespMag
	APIName := "GetMessageByRoomChatId"
	loggers.LogStart(APIName)

	roomChatIdStr := c.Param("roomChatId")
	roomChatId, err := strconv.ParseUint(roomChatIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	resp, err := ctrl.Ctx.GetMessageByRoomChatId(uint(roomChatId))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	if len(resp) > 0 {
		res.Data = resp
		res.Total = len(resp)
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) SendMessage(c echo.Context) error {
	var res response.RespMag
	APIName := "SendMessage"
	loggers.LogStart(APIName)

	var req request.SendMessage

	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	values := token.GetValuesToken(c)
	userId := values.ID

	err = ctrl.Ctx.SendMessage(req, userId)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) UpdateMessage(c echo.Context) error {
	var res response.RespMag
	APIName := "UpdateMessage"
	loggers.LogStart(APIName)

	var req request.SendMessage

	err := c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	values := token.GetValuesToken(c)
	userId := values.ID

	messageIdStr := c.Param("messageId")
	messageId, err := strconv.ParseUint(messageIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.UpdateMessage(req, uint(messageId), userId, values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteMessage(c echo.Context) error {
	var res response.RespMag
	APIName := "DeleteMessage"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	userId := values.ID

	messageIdStr := c.Param("messageId")
	messageId, err := strconv.ParseUint(messageIdStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	roomChatIDStr := c.Param("roomChatId")
	roomChatID, err := strconv.ParseUint(roomChatIDStr, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	err = ctrl.Ctx.DeleteMessage(uint(messageId), uint(roomChatID), userId, values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Msg = constant.SuccessMsg
	res.Code = constant.SuccessCode
	return response.EchoSucceed(c, res, APIName)
}
