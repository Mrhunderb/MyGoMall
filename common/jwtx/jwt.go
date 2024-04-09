package jwtx

import (
	"github.com/golang-jwt/jwt"
)

func GetToken(key string, iat, second int64, uid uint64) (string, error) {
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": iat + second,
		"iat": iat,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func IsExpired(tokenString, key string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true)
}
