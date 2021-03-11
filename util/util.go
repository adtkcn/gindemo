package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("asdasjkasioasndlkahsdhalksd")

// Claims 生成token结构体
type Claims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 创建token
func GenerateToken(ID, Name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		ID,
		Name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	// SigningMethodHS256、SigningMethodHS384、SigningMethodHS512
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// CheckExp 检查token是否有效
func CheckExp(token string) bool {
	claims, _ := ParseToken(token)
	nowTime := time.Now().Unix()
	if claims.ExpiresAt < nowTime {
		fmt.Println("登录过期")
		return false
	}
	fmt.Println("登录没过期")
	return true
}
