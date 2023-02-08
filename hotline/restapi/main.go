package restapi

import (
	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"hotline/utility/token"
	"net/http"
	"time"
)

func NewControllerMain(ctrl Controller) {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(200)))

	//r := e.Group("/SosApp")
	r := e.Group(config.GetString("service.endpoint"))
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Ok")
	})
	// "/user"
	u := r.Group(config.GetString("role.user"))
	{
		config := middleware.JWTConfig{
			Claims:     &token.JwtCustomClaims{},
			SigningKey: []byte(config.GetString("jwt.secret")),
		}
		r.Use(middleware.JWTWithConfig(config))
		u.GET("/name", func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			name := claims["id"].(string)
			return c.String(http.StatusOK, "Welcome "+name+"!")
		})
	}

	//e.Logger.Fatal(e.Start(":80"))
	e.Start(":" + config.GetString("service.port"))
	//e.Server.Addr = ":" + config.GetString("service.port")
	err := graceful.ListenAndServe(e.Server, 5*time.Second)

	if err != nil {
		panic(err)
	}
}
