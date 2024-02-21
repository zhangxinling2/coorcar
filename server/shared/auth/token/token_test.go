package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvJJiJIZHRYmw4JxRS3Dq
ESIWYmSifp1t1kI9iIVlZOy0rBK6E1O4PwEzohxkLXL3tOrMEcepxN6KhGb4bSTu
FKOgI01DLRvVeTRlqUeDnSwMzdZHHJI0OhrG4v/8Chc+vOSWihpsF7irREC4tK9z
RR9SWH8eZyizr0kzKl/r2YYLOHBq6MikI5V8gQZI9tLZb5NWmI/e640ctWB+AIfq
0mEJos84io36xRKfPaLw6nV4m6ORI/DUOn7JCm3afpV+4aoRVADY6ob7QvU8ZThT
vqwJV+BwQ5J8g6WB+/DTjPkI6XY6vFAB65T3oPy5znc2nfBw+2bp65lIhHNA6h4x
7wIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	testcases := []struct {
		name    string
		wantErr bool
		tkn     string
		now     time.Time
		want    string
	}{
		{
			name: "valid_token",
			tkn:  "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjVhNzEzMjIxYjg4NmQzYTliYmEyMWUzIn0.aO46Zml1ZmSkmL6m7weEj2R69-CPDG7tYkEhRSdnD_v8gLtzpJ5Kaag0u5QMQ7dKcs13FwTtfYr7AaB1fx1YEBCfTIjjNyDCnaEa4cnULzLMjrZLcj9ByGebj-fE4BFlh4nADQXd96hbFoUUoh6GbisFhVZovpsRVRwy7EFvH3beVo9bp5PUTIlB8GKj4kONOyt8lth0W3uhUzIWdg-Qw-VK6xvIOiKd5iUyqKl_N3HuN6L8wSM1Eq7O38oRtCnkjxYnaWyIyByDsPN1-eJ4d3r5U0JzCESycUq1BMXGYgasGgeqkL-APu1nzJ9t6PeXPvJI_rc7bMMUzB2gbKWCJw",
			now:  time.Unix(1516239122, 0),
			want: "65a713221b886d3a9bba21e3",
		},
		{
			name:    "token_expire",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjVhNzEzMjIxYjg4NmQzYTliYmEyMWUzIn0.aO46Zml1ZmSkmL6m7weEj2R69-CPDG7tYkEhRSdnD_v8gLtzpJ5Kaag0u5QMQ7dKcs13FwTtfYr7AaB1fx1YEBCfTIjjNyDCnaEa4cnULzLMjrZLcj9ByGebj-fE4BFlh4nADQXd96hbFoUUoh6GbisFhVZovpsRVRwy7EFvH3beVo9bp5PUTIlB8GKj4kONOyt8lth0W3uhUzIWdg-Qw-VK6xvIOiKd5iUyqKl_N3HuN6L8wSM1Eq7O38oRtCnkjxYnaWyIyByDsPN1-eJ4d3r5U0JzCESycUq1BMXGYgasGgeqkL-APu1nzJ9t6PeXPvJI_rc7bMMUzB2gbKWCJw",
			now:     time.Unix(1517239322, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1517239322, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signture",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjVhNzEzMjIxYjg4NmQzYTliYmEyMWUzIn0.aO46Zml1ZmSkmL6m7weEj2R69-CPDG7tYkEhRSdnD_v8gLtzpJ5Kaag0u5QMQ7dKcs13FwTtfYr7AaB1fx1YEBCfTIjjNyDCnaEa4cnULzLMjrZLcj9ByGebj-fE4BFlh4nADQXd96hbFoUUoh6GbisFhVZovpsRVRwy7EFvH3beVo9bp5PUTIlB8GKj4kONOyt8lth0W3uhUzIWdg-Qw-VK6xvIOiKd5iUyqKl_N3HuN6L8wSM1Eq7O38oRtCnkjxYnaWyIyByDsPN1-eJ4d3r5U0JzCESycUq1BMXGYgasGgeqkL-APu1nzJ9t6PeXPvJI_rc7bMMUzB2gbKWCJ",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
			if err != nil {
				t.Fatalf("connot parse public key:%v", err)
			}
			jwt.TimeFunc = func() time.Time {
				return tc.now
			}
			v := &JWTTokenVerify{
				publicKey: pk,
			}
			res, err := v.Verify(tc.tkn)
			if !tc.wantErr && err != nil {
				t.Errorf("verification failed:%v", err)
			}
			if tc.wantErr && err == nil {
				t.Errorf("want err,got no err")
			}
			if res != tc.want {
				t.Error("verify failed")
			}
		})
	}
}
