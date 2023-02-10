package restapi

import (
	"accounts/constant"
	"accounts/utility/response"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	jwt.RegisteredClaims
}

func NewControllerMain(ctrl Controller) {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(200)))

	r := e.Group(config.GetString("service.endpoint"))
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Ok")
	})
	r.POST("/sendOTP", ctrl.SendOTP)
	r.POST("/verifyOTP", ctrl.VerifyOTP)
	r.POST("/createUser", ctrl.CreateUser)
	r.POST("/signIn", ctrl.SignInUser)
	// "/user"
	u := r.Group(config.GetString("role.user"))
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtCustomClaims)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				role := c.Get("role").(*jwt.Token)
				claims := role.Claims.(jwt.MapClaims)
				if claims["role"].(string) != config.GetString("role.user") {
					var res response.RespMag
					res.Code = constant.ErrorStatusUnauthorized
					res.Msg = "Unauthorized"
					return c.JSON(http.StatusUnauthorized, res)
				} else {
					var res response.RespMag
					res.Code = constant.ErrorStatusUnauthorized
					res.Msg = "Invalid role"
					return c.JSON(http.StatusUnauthorized, res)
				}
			},
			SigningKey: []byte(config.GetString("jwt.secret")),
		}
		u.Use(echojwt.WithConfig(config))

		u.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "Ok Token")
		})
	}

	a := r.Group("/admin")
	{
		a.GET("/role", ctrl.GetRoleList)
		a.POST("/role", ctrl.AddRole)
		//a.PUT("/role", AddRole)
		//a.DELETE("/role", AddRole)
	}

	e.Start(":" + config.GetString("service.port"))
	err := graceful.ListenAndServe(e.Server, 5*time.Second)

	if err != nil {
		panic(err)
	}
}
