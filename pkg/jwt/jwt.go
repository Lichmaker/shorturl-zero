package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/lichmaker/short-url-micro/pkg/helpers"
)

type MyJwt struct {
	Secret    string
	ExpiresAt int64
}

type CustomClaims struct {
	AppId string `json:"app_id"`
	jwt.StandardClaims
}

func (myJwt *MyJwt) Generate(appId string) (string, error) {
	claims := &CustomClaims{
		AppId: appId,
	}
	if myJwt.ExpiresAt > 0 {
		claims.ExpiresAt = myJwt.ExpiresAt
	} else {
		claims.ExpiresAt = helpers.GetTimestamp() + 3600
		myJwt.ExpiresAt = claims.ExpiresAt
	}

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return j.SignedString([]byte(myJwt.Secret))
}
