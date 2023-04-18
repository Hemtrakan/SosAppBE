package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type ValueKey struct {
	ID   uint
	Role string
}

func GetValuesToken(c echo.Context) (res ValueKey) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	res = ValueKey{
		ID:   claims.ID,
		Role: claims.Role,
	}
	return
}

func GetAuthToken(c echo.Context) string {
	return c.Request().Header.Get("Authorization")
}
