package cases

func (a *AuthCasesProvider) RefreshTokens(login, accessToken, refreshToken string) (string, string, error) {
	accessValidErr := a.tokensProvider.ValidAccessToken(login, accessToken)
	if accessValidErr != ErrAccessTokenIsExpired {
		return "", "", accessValidErr
	}
	dbUser, getUserErr := a.userStorage.GetUser(login)

	if getUserErr == ErrUserNotFound {
		return "", "", getUserErr
	}

	if getUserErr != nil {
		// обработать неожиданные ошибки
	}

	refreshIndex, validRefreshErr := a.tokensProvider.ValidRefreshToken(refreshToken, accessToken, dbUser.RefreshInfos)

	if validRefreshErr != nil {
		return "", "", validRefreshErr
	}

	newAccessToken, createAccessTokenErr := a.tokensProvider.CreateAccessToken(login)

	if createAccessTokenErr != nil {
		// обработать неожиданные ошибки
	}

	newRefreshToken, newRefreshTokenInfo := a.tokensProvider.CreateRefreshToken(newAccessToken)

	dbUser.RefreshInfos[refreshIndex] = newRefreshTokenInfo

	updateErr := a.userStorage.UpdateUser(dbUser.Login, dbUser)
	if updateErr != nil {
		// обработать неожиданные ошибки
	}
	return newAccessToken, newRefreshToken, nil
}
