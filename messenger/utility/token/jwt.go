package token

import (
	"github.com/golang-jwt/jwt"
	config "github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(id uint, role string) (token string, Error error) {
	claims := &JwtCustomClaims{
		id,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	tokenResp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := tokenResp.SignedString([]byte(config.GetString("jwt.secret")))
	if err != nil {
		Error = err
		return
	}
	token = t
	return
}
