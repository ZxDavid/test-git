package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"web.com/ginGormJwt/model"
)

//定义jwt加密的密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//登陆成功后，调用这个方法发放token
func ReleaseToken(user model.User) (string, error) {
	//定义token的有限期(7*24*小时=7day)
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), //发放的时间
			Issuer:    "颁发者",
			Subject:   "user token", //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //使用上述的密钥生成token

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err

}
