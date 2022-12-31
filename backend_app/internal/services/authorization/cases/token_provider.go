package cases

import "errors"

type TokensProvider interface {
	CreateAccessToken(login string) (token string, err error)

	CreateRefreshToken(accessToken string) (token string, tokenInfo RefreshInfo)

	ValidAccessToken(login, accessToken string) error

	ValidRefreshToken(refreshToken, accessToken string, refreshInfos []RefreshInfo) (refreshIndex int, err error)
}

var (
	ErrAccessTokenIsExpired      error = errors.New("access token is expired")
	ErrRefreshTokenIsExpired     error = errors.New("refresh token is expired")
	ErrInvalidClaims             error = errors.New("access token claims is invalid")
	ErrTooShortAccessToken       error = errors.New("access token is too short")
	ErrRefreshTokenInvalidLength error = errors.New("invalid refresh token length")
	ErrRefreshBodyInvalid        error = errors.New("refresh body is not valid")
	ErrRefreshNotRelatedToAccess error = errors.New("refresh token does't relate to access")
)
