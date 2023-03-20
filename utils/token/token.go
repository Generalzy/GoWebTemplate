package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	SecretKey = []byte("hello world")
	Issuer    = "Generalzy"
	ExpireAt  = time.Hour * 24 * 7
)

type JsonWebTokenClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := JsonWebTokenClaim{
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireAt)),
			Issuer:    Issuer, // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

func ParseToken(token string) (*JsonWebTokenClaim, error) {
	t, err := jwt.ParseWithClaims(token, new(JsonWebTokenClaim), func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := t.Claims.(*JsonWebTokenClaim); ok && t.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
