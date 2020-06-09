package util

import (
	_ "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string, cost int) (string, error) {
	if cost < 1 {
		cost = 1
	}
	bytes, e := bcrypt.GenerateFromPassword([]byte(password), cost)
	if e != nil {
		return "", e
	}
	return string(bytes), nil
}
func ValidatePassword(hashPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
