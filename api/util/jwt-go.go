package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("mySecret")

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func GenerateToken(nickname string, expiration time.Duration) (string, error) {
	now := time.Now()
	expireTime := now.Add(expiration)
	claims := Claims{
		nickname,
		jwt.StandardClaims{
			Audience:  "all",             //接受对象
			ExpiresAt: expireTime.Unix(), //在什么时候过期
			//Id:        "",//唯一身份标识，主要用来作为一次性token，用来回避重放攻击使用的
			IssuedAt:  now.Unix(),   //在什么时候签发
			Issuer:    "com.weiguo", //签发人
			NotBefore: now.Unix(),   //生效时间
			Subject:   nickname,     //使用人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token,err
}

func keyFunc(token *jwt.Token) (interface{},error) {
	return jwtSecret,nil
}
func ParseToken(token string) (*Claims, error) {
	claims, e := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if claims != nil{
		if resultClaims,ok := claims.Claims.(*Claims); ok && claims.Valid{
			return resultClaims,nil
		}
	}
	return nil ,e
}