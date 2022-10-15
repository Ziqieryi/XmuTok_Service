package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var myKey = []byte("v1.0")

// GenerateToken
// 生成 token
func GenerateToken(userId int64, username string, password string) (string, error) {
	UserClaim := &UserClaims{
		UserId:         userId,
		Username:       username,
		Password:       password,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
