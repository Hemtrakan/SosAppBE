package token

import (
	"github.com/labstack/echo/v4"
	config "github.com/spf13/viper"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(id uint, role string) (token string, Error error) {
	claims := &JwtCustomClaims{}
	claims.ID = id
	claims.Role = role
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	tokenResp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tokenResp.SignedString([]byte(config.GetString("jwt.secret")))
	if err != nil {
		Error = err
		return
	}
	token = t
	return
}

type ValueKey struct {
	ID   uint
	Role string
}

func GetTokenKay(c echo.Context) (res ValueKey) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	res = ValueKey{
		ID:   claims.ID,
		Role: claims.Role,
	}
	return
}
