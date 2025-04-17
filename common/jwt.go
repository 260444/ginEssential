package common

import (
	"time"

	"github.com/260444/ginEssential/model"
	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

// 生成token函数
func ReleaseToken(user model.User) (string, error) {
	exoirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exoirationTime),
			// 颁发时间 也就是生成时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签发人
			Issuer: "oceanlearn.tech",
			//主题
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

// 解析token函数
func ParseToken(tokenstring string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	return token, claims, err
}
