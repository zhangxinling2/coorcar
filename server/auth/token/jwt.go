package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtTokenGen struct {
	issuer     string
	nowFunc    func() time.Time
	privateKey *rsa.PrivateKey
}

func NewJwtTokenGen(issuer string, privateKey *rsa.PrivateKey) *JwtTokenGen {
	return &JwtTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}
func (t *JwtTokenGen) GenerateToken(accountId string, expire time.Duration) (string, error) {
	now := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Subject:   accountId,
		Issuer:    t.issuer,
		IssuedAt:  now,
		ExpiresAt: now + int64(expire.Seconds()),
	})
	return tkn.SignedString(t.privateKey)
}
