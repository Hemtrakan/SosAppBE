package restapi

import (
	"accounts/constant"
	"accounts/utility/response"
	"accounts/utility/token"
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
	e.Use(middleware.Logger())

	s := e.Group(config.GetString("service.endpoint"))
	s.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Ok")
	})
	s.POST("/sendOTP", ctrl.SendOTP)
	s.POST("/verifyOTP", ctrl.VerifyOTP)
	s.POST("/createUser", ctrl.CreateUser)
	s.POST("/signIn", ctrl.SignIn)

	configs := echojwt.Config{
		ErrorHandler:  ErrorHandler,
		SigningKey:    []byte(config.GetString("jwt.secret")),
		NewClaimsFunc: NewClaimsFunc,
	}

	u := s.Group(config.GetString("role.user"))
	u.Use(echojwt.WithConfig(configs), AuthRoleUser)

	u.GET("/:id", ctrl.GetUserById)
	u.PUT("/:id", ctrl.UpdateUser)
	u.PUT("/:id", ctrl.ChangePassword)
	u.DELETE("/:id", ctrl.DeleteUser)

	// todo Verify ID Card API admin page
	// todo admin
	a := s.Group("/admin")
	a.Use(echojwt.WithConfig(configs), AuthRoleAdmin)

	a.GET("/role", ctrl.GetRoleList)
	a.GET("/role/:id", ctrl.GetRoleById)
	a.POST("/role", ctrl.AddRole)
	a.PUT("/role/:id", ctrl.UpdateRole)    // todo ยังไม่เสร็จ
	a.DELETE("/role/:id", ctrl.DeleteRole) // todo ยังไม่เสร็จ
	//
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
