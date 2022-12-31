package tokens

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func (t *JWTTokenProvider) CreateAccessToken(login string) (token string, err error) {
	var claims JWTTokenClaims = JWTTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.jwtLifeTime).Unix(),
			Subject:   login,
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return rawToken.SignedString(t.jwtKey)
}
