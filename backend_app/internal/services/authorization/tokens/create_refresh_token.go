package tokens

import (
	"Synchronyzed-Music-Player/internal/services/authorization/cases"
	"time"
)

func (t *JWTTokenProvider) CreateRefreshToken(accessToken string) (token string, tokenInfo cases.RefreshInfo) {
	accessTokenByte := []byte(accessToken)
	if len(accessTokenByte) < 7 {
		return "", cases.RefreshInfo{}
	}

	bodyChan := make(chan string, 1)
	lifeTimeChan := make(chan string, 1)

	defer close(bodyChan)
	defer close(lifeTimeChan)

	go generateRefreshBody(bodyChan, t.refreshBodyLength)
	go timeToASCII(lifeTimeChan, time.Now().Add(t.refreshLifeTime))

	refreshBody := <-(bodyChan)
	var resultToken string = <-lifeTimeChan + refreshBody + accessToken[len(accessTokenByte)-6:]

	return resultToken, cases.RefreshInfo{
		Body:       refreshBody,
		AccessPart: accessToken[len(accessTokenByte)-6:],
	}
}
