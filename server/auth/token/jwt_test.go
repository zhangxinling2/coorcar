package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8JA0AXCXqGqtx
CIg0aRLTx2c6/QSZYsJ20lu1YLTnuv7qO66GQ3fsWdQzDu/85U5ADuBSYNXXpxxX
jHa9ankyBgAbxBUYWkROt0hucmfPeOTNWnYO+6roOWBXOpXDUdueyIBZvHpg4EMd
c2yWI15f3m9RM9k+zqScsEKX5MfgvRyjk6wx0A6hZGvdLYpdEB48bli9+2jTnf0j
X98GW9vJEdfP1u2di9RjaVfJg3+5jiDVuuuzM0gcdLM7CrsPYJVQG9UYLj3Lc8ug
LYCMAyUIFk4P+buKTcUruO77m1NC6ozO3nBWsqOkrkvEwYN7y1vaGIy5Ei70UFCW
AhPgiYYfAgMBAAECggEAVwbqCK+XTinWAewrbOrprZLAWpXtSbJYN7RofYblR4fF
IL5IpYyR7X9D2/Zd5oJ/Zrw64YAvipifUJQF3Opic1Q5Op7CD0i57htz4rsC8AKA
+8iLcbTSQqc/vDm6SuYTi+g9DsVTanpyMS5SHUTxA7GVll9Ab4AQvvX9Z9uI5Zhp
SWijm1Y2Z7Dw3OHcnNXZF0XK8fZOIdGaOSYCrdqxmvyzlKyenDpqTb70/lZ3WoVw
mKPeTRMBzJIx7E5Ev1nbnC+HRgxa/MVFjROj2G+3BADAgZHKZdLdCh8NUVVvkgN1
SAh4oXoyAOV+prIr/8xXMSyxEAxxshewZody96qbAQKBgQDoSL/ped4COY/S5rjo
HnRWnCM5weJ8kgy7VTzTWGgRSvWyDg7OyRA8dOPIYQSAoAvUwpYPwqoG/4RJIfNy
vNKHZ2QB8sb4bmu5ocOvo9wCQy4ZgbEupkQ8Gk78O48fQhGWBr4JpMvLqo/tnJiK
WQPlsrraZiy6x3vRnVLwj3QlnwKBgQDPWYTWLwZ2HeCrMBNsTXecpnWC68WObSk3
6yrtCxxWoEZDwfc/aSrzbaU/LEX2vVIASqwTWp2MRiypt0QureMkoh6XrIVrCLNv
ZbT1Ex9/fnSGpvB23s7LpVjCeAIhJIAqdX06TUBRuz5vnmuuH0WQ189PIGSXI9Mm
LIb2ZgnPgQKBgQClvVMCGqyoVb4js/D33tMQAXS4/3jFfciL2nP+MTSAMkZBd/pk
3d4x5Q2qql72EDyg4P8F29ZKDAOt0wXMNqLZbpJTOZFzLga7e4sPRhHzo0+/b2Sn
bod9gyAcpbqniRLXQ5RReT7Y4ai02GmXPZCkUvU0ZHdCBph4ppLJhvJO7QKBgGiz
ll0lD8kzGxC2TQiCMaT7dRkv/ukFfCvNZzadRqkDc5ehaT0u2M5lDm1Q6N9ICBE3
o55+QlHj2Ye1kAtWKP2Fk5cZOaGrmhflLrQV/PC/JQ4dXA1aQugYTTKdsM2VhdZq
z1pSSek7eLKaN1uMVO9n2jnlbIXZRDThhSEEs22BAoGBALfuQrPPCXyH0WS6bFfj
EwrX+PP4utTncIprkrvy314I/0oB8RSKJO/MbPEQ3BxOG+r7MdiSr5CyqiTDwALO
ET8pC/EiS4mYOrYWjPl8U9FSBuGktXj4VSO2ZSo5oQDawoONXGvcCuXR+i+iYnLN
PfgeBQDKSl64L8h1IgCjMpGY
-----END PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse privateKey: %v", err)
	}
	g := NewJwtTokenGen("server/auth", pk)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tk, err := g.GenerateToken("65a713221b886d3a9bba21e3", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token:%v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoic2VydmVyL2F1dGgiLCJzdWIiOiI2NWE3MTMyMjFiODg2ZDNhOWJiYTIxZTMifQ.PwszhS6SUgpYUgZ0xVQ-ov1B2D6Nwj0jZVUfQNqtCUtsnxq3OQy6MMzr1tp0tWPpCol9jxZ0MDDuqYpxcRFOBGvaHoPI1t6jx2RyW2tkgg5vYP5a3J5_eYipEaUls-1n1U5WBFgB8KKvhkwtngoLwQG1SYzV9x0qM1hQfbRyr4Ka37m9hKrP5_U16cn-GU_Kparu5iCWRiRO5BIg4Bq4Rj96iDL_L1ELlIzzfP5Po1X7RqFFGMfIrtK2Z-wD-G8xS13sZIpxEDtFPE-Er1Bmi1HeOzt1XMdg1Z8fJ-mZ3rBAL4J0psfcjTjh2u8e12ZCY3aHl0DSCG0cempzOyPzGw"
	if tk != want {
		t.Fatalf("want :%s;got :%s", want, tk)
	}
}
