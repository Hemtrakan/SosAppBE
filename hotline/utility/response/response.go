package response

import (
	"github.com/labstack/echo/v4"
	"hotline/utility/loggers"
	"net/http"
	"strconv"
)

type RespMag struct {
	Code  string      `json:"code"`
	Msg   string      `json:"message"`
	Total int         `json:"total,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func EchoSucceed(c echo.Context, msg interface{}, APIName string) error {
	loggers.LogProvider(strconv.Itoa(http.StatusOK), APIName, msg)
	return c.JSON(http.StatusOK, msg)
}

func EchoError(c echo.Context, statusCode int, msg interface{}, APIName string) error {
	loggers.LogProvider(strconv.Itoa(statusCode), APIName, msg)
	return c.JSON(statusCode, msg)
}
