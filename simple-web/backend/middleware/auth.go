package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

type jwtConfig struct {
	Secret []byte
}

var config = jwtConfig{Secret: []byte("hamdppyy")}

var db map[string]string

func findUser(username, password string) int {
	i := 0
	for u, p := range db {
		if u == username && p == password {
			return i
		}
		i += 1
	}
	return -1
}
func createJWT(id uint64, username string) (string, error) {
	expiresTime := time.Now().Unix() + int64(3600*24)
	claims := jwt.StandardClaims{
		Audience:  username,     // 受众
		ExpiresAt: expiresTime,       // 失效时间
		Id:        string(id),   // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "dasin",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "auth",           // 主题
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(config.Secret)
}
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return config.Secret, nil
	})
	if err == nil  {
		if jwtToken == nil {
			return nil, errors.New("empty token was got after parsing")
		}
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err

}

func Auth(c * gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	auth = strings.Fields(auth)[1]
//	log.Printf("found request head auth: %v\n", auth)
	if _, err := parseToken(auth); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("fail: %v", err),
		})
		c.Abort()
	} else {
		c.Next()
	}
}

type LoginInfo struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	log.Println("enter login handler")
	var info LoginInfo
	if c.BindJSON(&info) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong request params",
		})
		return
	}
	id := findUser(info.Username, info.Password)
	if id < 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "not found such user",
		})
		return
	}
	token, err := createJWT(uint64(id), info.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "create jwt error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"username": info.Username,
		"jwt": token,
	})
}

func init() {
	db = make(map[string]string)
	db["dasin"] = "123"
}