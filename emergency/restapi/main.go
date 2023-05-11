package restapi

import (
	"emergency/constant"
	"emergency/utility/response"
	"emergency/utility/token"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
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

	r := e.Group(config.GetString("service.endpoint"))

	// user
	u := r.Group(config.GetString("role.user"))
	u.Use(echojwt.WithConfig(configs), AuthRoleUser)
	u.GET("/", ctrl.GetInformList)
	u.GET("/:id", ctrl.GetInformById)
	u.POST("/", ctrl.PostInform)
	//u.POST("/:id", ctrl.PutInform)
	//u.POST("/:id", ctrl.DeleteInform)

	o := r.Group(config.GetString("role.ops"))
	o.Use(echojwt.WithConfig(configs), AuthRoleOps)
	o.GET("/", ctrl.GetInformOpsList)
	o.GET("/all", ctrl.GetAllInformList)
	o.GET("/:id", ctrl.GetInformOpsById)
	o.PUT("/:id", ctrl.UpdateInform)
	o.DELETE("/:id", ctrl.DeleteInform)

	// admin
	a := r.Group(config.GetString("role.admin"))
	a.Use(echojwt.WithConfig(configs), AuthRoleAdmin)
	//a.GET("/")
	//a.GET("/:id")
	//a.POST("/")
	//a.PUT("/:id")
	//a.DELETE("/:id")

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
