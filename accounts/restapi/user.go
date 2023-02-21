package restapi

import (
	"accounts/constant"
	userReq "accounts/restapi/model/user/request"
	"accounts/utility/response"
	"accounts/utility/token"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (ctrl Controller) GetUserById(c echo.Context) error {
	var res response.RespMag
	values := token.GetValuesToken(c)

	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	if uint(id) != values.ID {
		res.Code = constant.ErrorCode
		res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	data, err := ctrl.Ctx.GetUser(uint(id))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = data
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) UpdateUser(c echo.Context) error {
	var res response.RespMag
	values := token.GetValuesToken(c)
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	if uint(id) != values.ID {
		res.Code = constant.ErrorCode
		res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	var req = new(userReq.UserReq)
	err = c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	err = ValidateStruct(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	errArr := ctrl.Ctx.PutUser(req, uint(id))
	if errArr != nil {
		errRes := ""
		for _, m1 := range errArr {
			errRes = errRes + m1.Error() + " | "
		}
		res.Code = constant.ErrorCode
		res.Msg = errRes
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "UpdateSuccess"
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) DeleteUser(c echo.Context) error {
	var res response.RespMag
	values := token.GetValuesToken(c)
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	if uint(id) != values.ID {
		res.Code = constant.ErrorCode
		res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	err = ctrl.Ctx.DeleteUser(uint(id))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "DeleteSuccess"
	return response.EchoSucceed(c, res)
}

func (ctrl Controller) ChangePassword(c echo.Context) error {
	var res response.RespMag
	values := token.GetValuesToken(c)
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	if uint(id) != values.ID {
		res.Code = constant.ErrorCode
		res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	var req = new(userReq.UserReq)
	err = c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	err = ValidateStruct(req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	errArr := ctrl.Ctx.PutUser(req, uint(id))
	if errArr != nil {
		errRes := ""
		for _, m1 := range errArr {
			errRes = errRes + m1.Error() + " | "
		}
		res.Code = constant.ErrorCode
		res.Msg = errRes
		return response.EchoError(c, http.StatusBadRequest, res)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "UpdateSuccess"
	return response.EchoSucceed(c, res)
}