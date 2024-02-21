package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerify struct {
	publicKey *rsa.PublicKey
}

func (v *JWTTokenVerify) Verify(token string) (string, error) {
	tk, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return v.publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token:%v", err)
	}
	//验证签名
	if !tk.Valid {
		return "", fmt.Errorf("token not valid: %v", err)
	}
	//验证内容结构类型
	clm, ok := tk.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("claim not standardClaim")
	}
	//验证claim内容valid
	if err = clm.Valid(); err != nil {
		return "", fmt.Errorf("claim is not valid:%v", err)
	}
	return clm.Subject, nil
}
