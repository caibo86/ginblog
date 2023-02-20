package utils

import (
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtKeyBytes = []byte(JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// SetToken 生成token
func SetToken(username string) (string, error) {
	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour)),
			Issuer:    "ginblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JwtKeyBytes)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// CheckToken 验证token
func CheckToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKeyBytes, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errmsg.ErrJwtTokenInvalid
}
