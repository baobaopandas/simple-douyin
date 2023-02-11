package util

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/config"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	"github.com/golang-jwt/jwt/v4"
)

// 自定义claim
type Claims struct {
	UserId               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        //v4,老版本这里是jwt.StandardClaims
}

// 生成token
func GenerateToken(user *mydb.User) (tokenstring string, err error) {

	sercetKey := []byte(config.CONFIG.JWTConfig.JWTSecret)
	claim := Claims{
		UserId:   user.UserID,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.CONFIG.JWTConfig.ExpireTime) * time.Hour * time.Duration(1))),
			//令牌签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	//这里选择HS256是单密钥算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err = token.SignedString(sercetKey)
	return tokenstring, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(config.CONFIG.JWTConfig.JWTSecret), nil // 这是我的secret
	}
}

func ParseToken(tokenstring string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			//token格式错误||过期||未激活||无法处理改密钥
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
