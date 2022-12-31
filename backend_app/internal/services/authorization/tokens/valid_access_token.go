package tokens

import (
	"Synchronyzed-Music-Player/internal/services/authorization/cases"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (t *JWTTokenProvider) ValidAccessToken(login, accessToken string) error {
	token, parseErr := jwt.ParseWithClaims(accessToken, &JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.jwtKey, nil
	})
	if parseErr != nil {
		return parseErr
	}
	claims, ok := token.Claims.(*JWTTokenClaims)
	if time.Now().After(time.Unix(claims.ExpiresAt, 0)) {
		return cases.ErrAccessTokenIsExpired
	} else if !ok {
		return cases.ErrInvalidClaims
	} else if claims.Subject != login {
		return cases.ErrInvalidClaims
	}
	return nil
}
