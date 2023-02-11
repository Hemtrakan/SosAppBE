package token

import (
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
