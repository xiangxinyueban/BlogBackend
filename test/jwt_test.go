package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var mykey = []byte("blogbackend")

func TestGenerateToken(t *testing.T) {
	UserClaim := &UserClaims{
		Identity: "User1",
		Name:     "get",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	signedString, err := token.SignedString(mykey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(signedString)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IlVzZXIxIiwibmFtZSI6ImdldCJ9.VbYwCBvfG3h4DiphLuVtXdix9ERdVnxfkflsFnr6eFw
}

func TestDecodeToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IlVzZXIxIiwibmFtZSI6ImdldCJ9.VbYwCBvfG3h4DiphLuVtXdix9ERdVnxfkflsFnr6eFw"
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return mykey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}
