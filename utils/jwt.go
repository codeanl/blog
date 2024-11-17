package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

// Claims 通常指数据
type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenToken(secret, issuer string, expireHour, userId int) (string, error) {
	claims := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	// 生成jwt的header 和 playload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 进行签名 生成最终的jwt
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, token string) (*MyClaims, error) {
	// 解析token
	t, err := jwt.ParseWithClaims(token, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		switch jerr, ok := err.(*jwt.ValidationError); ok {
		case jerr.Errors&jwt.ValidationErrorMalformed != 0:
			return nil, ErrTokenMalformed
		case jerr.Errors&jwt.ValidationErrorExpired != 0:
			return nil, ErrTokenExpired
		case jerr.Errors&jwt.ValidationErrorNotValidYet != 0:
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}

	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
