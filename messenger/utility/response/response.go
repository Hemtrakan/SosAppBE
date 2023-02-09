package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RespMag struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func EchoSucceed(c echo.Context, msg interface{}) error {
	return c.JSON(http.StatusOK, msg)
}

func EchoError(c echo.Context, statusCode int, msg interface{}) error {
	return c.JSON(statusCode, msg)
}
