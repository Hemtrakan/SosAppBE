package restapi

import (
	"accounts/constant"
	userReq "accounts/restapi/model/user/request"
	"accounts/utility/loggers"
	"accounts/utility/response"
	"accounts/utility/token"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ctrl Controller) SearchUser(c echo.Context) error {
	var res response.RespMag
	APIName := "searchUser"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	value := c.Param("value")

	if len(value) < 3 {
		res.Code = constant.ErrorCode
		res.Msg = "ตัวอักษรต้องไม่ต่ำกว่า 3 ตัวขึ้นไป"
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	id := values.ID

	data, err := ctrl.Ctx.SearchUser(value, id)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Total = len(data)
	res.Data = data
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetUserByToken(c echo.Context) error {
	var res response.RespMag
	APIName := "getUserByToken"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)

	data, err := ctrl.Ctx.GetUser(values.ID)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = data
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetUserList(c echo.Context) error {
	var res response.RespMag
	APIName := "getUserList"
	loggers.LogStart(APIName)

	data, err := ctrl.Ctx.GetUserList()
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Total = len(data)
	res.Data = data
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetUserById(c echo.Context) error {
	var res response.RespMag
	APIName := "getUserById"
	loggers.LogStart(APIName)

	//values := token.GetValuesToken(c)
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	data, err := ctrl.Ctx.GetUser(uint(id))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = data

	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) GetImageById(c echo.Context) error {
	var res response.RespMag
	APIName := "getImageById"
	loggers.LogStart(APIName)

	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	data, err := ctrl.Ctx.GetImage(uint(id))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = data
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) UpdateUser(c echo.Context) error {
	var res response.RespMag
	APIName := "updateUser"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	if values.Role != "admin" {
		if uint(id) != values.ID {
			res.Code = constant.ErrorCode
			res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
			return response.EchoError(c, http.StatusBadRequest, res, APIName)
		}
	}

	var req = new(userReq.UserReq)
	err = c.Bind(&req)
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

	errArr := ctrl.Ctx.PutUser(req, uint(id))
	if errArr != nil {
		errRes := ""
		for _, m1 := range errArr {
			errRes = errRes + m1.Error() + " | "
		}
		res.Code = constant.ErrorCode
		res.Msg = errRes
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "UpdateSuccess"
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) VerifyIDCard(c echo.Context) error {
	var res response.RespMag
	APIName := "VerifyIDCard"
	loggers.LogStart(APIName)

	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	var req = new(userReq.VerifyIDCard)
	err = c.Bind(&req)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	errArr := ctrl.Ctx.VerifyIDCard(uint(id), req)
	if errArr != nil {
		errRes := ""
		for _, m1 := range errArr {
			errRes = errRes + m1.Error() + " | "
		}
		res.Code = constant.ErrorCode
		res.Msg = errRes
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "VerifySuccess"
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) DeleteUser(c echo.Context) error {
	var res response.RespMag
	APIName := "deleteUser"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	if values.Role != "admin" {
		if uint(id) != values.ID {
			res.Code = constant.ErrorCode
			res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
			return response.EchoError(c, http.StatusBadRequest, res, APIName)
		}
	}
	err = ctrl.Ctx.DeleteUser(uint(id))
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "DeleteSuccess"
	return response.EchoSucceed(c, res, APIName)
}

func (ctrl Controller) ChangePassword(c echo.Context) error {
	var res response.RespMag
	APIName := "changePassword"
	loggers.LogStart(APIName)

	values := token.GetValuesToken(c)
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 0, 0)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}
	if values.Role != constant.Admin {
		if uint(id) != values.ID {
			res.Code = constant.ErrorCode
			res.Msg = "ไม่สามารถทำรายการได้ กรุณาติดต่อผู้ดูแลระบบ"
			return response.EchoError(c, http.StatusBadRequest, res, APIName)
		}
	}

	var req = new(userReq.ChangePassword)
	err = c.Bind(&req)
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

	err = ctrl.Ctx.ChangePassword(req, uint(id), values.Role)
	if err != nil {
		res.Code = constant.ErrorCode
		res.Msg = err.Error()
		return response.EchoError(c, http.StatusBadRequest, res, APIName)
	}

	res.Code = constant.SuccessCode
	res.Msg = constant.SuccessMsg
	res.Data = "ChangePassword"
	return response.EchoSucceed(c, res, APIName)
}
