package tokens

import (
	"time"
)

type JWTTokenProvider struct {
	jwtLifeTime time.Duration
	jwtKey      []byte

	refreshLifeTime   time.Duration
	refreshBodyLength int
}

func (t *JWTTokenProvider) New(jwtLifeTime, refreshLifeTime time.Duration, jwtKey []byte, refreshBodyLength int) {
	t.jwtLifeTime = jwtLifeTime
	t.refreshLifeTime = refreshLifeTime
	t.jwtKey = jwtKey
	t.refreshBodyLength = refreshBodyLength
}
