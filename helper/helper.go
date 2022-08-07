package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//GetMd5
//生成MD5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

type UserClaims struct {
	Identity uint   `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var mykey = []byte("blogbackend")

//GenerateToken
//生成Token
func GenerateToken(identity uint, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(mykey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeToken

func DecodeToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return mykey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyze token err:%v", err)
	}
	return userClaim, nil
}
