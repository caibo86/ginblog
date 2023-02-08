package middleware

import (
	"github.com/caibo86/ginblog/utils"
	"github.com/caibo86/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// SetToken 生成token
func SetToken(username, password string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireTime,
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ErrServer
	}
	return token, errmsg.Success
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.Success
	}
	return nil, errmsg.ErrServer
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		var code errmsg.Code
		if tokenHeader == "" {
			code = errmsg.ErrUnauthorized
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrUnauthorized
			c.Abort()
		}
		key, ok := CheckToken(checkToken[1])
		if ok != errmsg.Success {
			code = ok
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrTokenExpired
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": code.Error(),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
