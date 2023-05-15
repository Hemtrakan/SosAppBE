package restapi

import (
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"hotline/constant"
	"hotline/utility/response"
	"hotline/utility/token"
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
	r.GET("/", ctrl.GetHotLine)
	r.POST("/", ctrl.PostHistory)

	a := r.Group(config.GetString("role.admin"))
	a.Use(echojwt.WithConfig(configs), AuthRoleAdmin)
	a.GET("/:id", ctrl.GetHotLineById)
	a.POST("/", ctrl.PostHotLine)
	a.PUT("/:id", ctrl.PutHotLine)
	a.DELETE("/:id", ctrl.DeleteHotLine)

	a.GET("/history", ctrl.GetHistory)

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
