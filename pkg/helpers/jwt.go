package helpers

import (
	"github.com/spf13/cast"
	"go-web/pkg/config"
	"time"
)
import "github.com/golang-jwt/jwt"

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

func SignedJWT(UserId uint64) (string, error) {
	expireTime := time.Now().Add(time.Hour * 7 * 24)

	claims := &Claims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //签名时间
			Issuer:    "127.0.0.1",       //签名颁发者
			Subject:   "UserId",          //签名主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := cast.ToString(config.Viper.Get("jwt.jwt_key"))
	return token.SignedString([]byte(jwtKey))
}
