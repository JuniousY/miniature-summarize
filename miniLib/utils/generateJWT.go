package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	tokenStr, _ := token.SignedString([]byte("testKey"))
	fmt.Println(tokenStr)
	fmt.Printf("'Authorization': `Bearer ${%v}`\n", tokenStr)
}
