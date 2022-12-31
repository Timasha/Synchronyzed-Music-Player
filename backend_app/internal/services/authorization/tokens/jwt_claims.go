package tokens

import "github.com/golang-jwt/jwt"

type JWTTokenClaims struct {
	jwt.StandardClaims
}
