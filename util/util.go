package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("asdasjkasioasndlkahsdhalksd")

// Claims 生成token结构体
type Claims struct {
	Username string `json:"username"`
	// Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 创建token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		// password,
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
