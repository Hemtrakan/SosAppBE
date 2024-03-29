package restapi

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
	//"github.com/gorilla/websocket"
	"messenger/constant"
	"messenger/utility/response"
	"messenger/utility/token"
	"net/http"
	"strings"
	"time"
)

func NewControllerMain(ctrl Controller) {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(200)))

	configs := echojwt.Config{
		ErrorHandler:  ErrorHandler,
		SigningKey:    []byte(config.GetString("jwt.secret")),
		NewClaimsFunc: NewClaimsFunc,
	}
	//

	r := e.Group(config.GetString("service.endpoint"))
	u := r.Group(config.GetString("role.user"))
	u.Use(echojwt.WithConfig(configs), AuthRoleAllUser)

	u.POST("/createRoomChat", ctrl.RoomChat)
	u.PUT("/updateRoomChat/:roomId", ctrl.UpdateRoomChat)
	u.DELETE("/deleteRoomChat/:roomId", ctrl.DeleteRoomChat)

	u.POST("/joinChat", ctrl.JoinChat)
	u.GET("/getChatList", ctrl.GetChatList)
	u.GET("/getMembersRoomChat/:roomChatId", ctrl.GetMembersRoomChat)
	u.GET("/chat/message/:roomChatId", ctrl.GetMessageByRoomChatId)
	u.GET("/chat/message/image/:messageId", ctrl.GetImageByMessageId)

	u.POST("/chat/message", ctrl.SendMessage)
	u.PUT("/chat/message/:messageId", ctrl.UpdateMessage)
	u.DELETE("/chat/message/:messageId/:roomChatId", ctrl.DeleteMessage)

	a := r.Group(config.GetString("role.admin"))
	a.Use(echojwt.WithConfig(configs), AuthRoleAdmin)

	a.POST("/createRoomChat", ctrl.RoomChat)
	a.PUT("/updateRoomChat/:roomId", ctrl.UpdateRoomChat)
	a.DELETE("/deleteRoomChat/:roomId", ctrl.DeleteRoomChat)

	a.POST("/joinChat", ctrl.JoinChat)
	a.GET("/getChatList", ctrl.GetChatList)
	a.GET("/getMembersRoomChat/:roomChatId", ctrl.GetMembersRoomChat)
	a.GET("/chat/message/:roomChatId", ctrl.GetMessageByRoomChatId)
	a.GET("/chat/message/image/:messageId", ctrl.GetImageByMessageId)
	a.POST("/chat/message", ctrl.SendMessage)

	a.PUT("/chat/message/:messageId", ctrl.UpdateMessage)
	a.DELETE("/chat/message/:messageId/:roomChatId", ctrl.DeleteMessage)

	e.Start(":" + config.GetString("service.port"))
	err := graceful.ListenAndServe(e.Server, 5*time.Second)

	if err != nil {
		panic(err)
	}
}

func NewClaimsFunc(c echo.Context) jwt.Claims {
	return new(token.JwtCustomClaims)
}

func ErrorHandler(c echo.Context, e error) error {
	var res response.RespMag
	res.Code = constant.ErrorStatusUnauthorized
	res.Msg = "Unauthorized"
	return c.JSON(http.StatusUnauthorized, res)
}

func AuthRoleUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*token.JwtCustomClaims)
		var userRole = claims.Role
		role := strings.Replace(config.GetString("role.user"), "/", "", 2)
		if userRole == role {
			return next(c)
		} else {
			c.Error(echo.ErrUnauthorized)
			return nil
		}
	}
}

func AuthRoleOps(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*token.JwtCustomClaims)
		var userRole = claims.Role
		role := strings.Replace(config.GetString("role.ops"), "/", "", 2)
		if userRole == role {
			return next(c)
		} else {
			c.Error(echo.ErrUnauthorized)
			return nil
		}
	}
}

func AuthRoleAllUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*token.JwtCustomClaims)
		var userRole = claims.Role

		fmt.Println(userRole)
		roleUser := strings.Replace(config.GetString("role.user"), "/", "", 2)
		roleOps := strings.Replace(config.GetString("role.ops"), "/", "", 2)
		if (userRole == roleUser) || (userRole == roleOps) {
			return next(c)
		} else {
			c.Error(echo.ErrUnauthorized)
			return nil
		}
	}
}

func AuthRoleAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*token.JwtCustomClaims)
		var userRole = claims.Role
		role := strings.Replace(config.GetString("role.admin"), "/", "", 2)
		if userRole == role {
			return next(c)
		} else {
			c.Error(echo.ErrUnauthorized)
			return nil
		}
	}
}
