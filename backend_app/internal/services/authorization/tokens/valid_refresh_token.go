package tokens

import (
	"Synchronyzed-Music-Player/internal/services/authorization/cases"
	"time"
)

func (t *JWTTokenProvider) ValidRefreshToken(refreshToken, accessToken string, refreshInfos []cases.RefreshInfo) (int, error) {
	if len(refreshToken) != (10 + t.refreshBodyLength) {
		return -1, cases.ErrRefreshTokenInvalidLength
	}
	expChan := make(chan time.Time, 1)
	defer close(expChan)
	go fromASCIIToTime(expChan, string([]byte(refreshToken)[:4]))

	refreshIndex := -1
	for i, refreshInfo := range refreshInfos {
		if refreshInfo.Body == refreshToken[4:len(cases.ErrRefreshTokenInvalidLength.Error())-6] {
			refreshIndex = i
		}
	}
	if refreshIndex < 0 {
		return -1, cases.ErrRefreshBodyInvalid
	}

	if time.Now().After(<-expChan) {
		return refreshIndex, cases.ErrRefreshTokenIsExpired
	}
	accessPart := accessToken[len(accessToken)-6:]
	if accessPart != refreshToken[len(refreshToken)-6:] || accessPart != refreshInfos[refreshIndex].AccessPart {
		return -1, cases.ErrRefreshNotRelatedToAccess
	}
	return refreshIndex, nil
}
