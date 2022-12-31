package cases

func (a *AuthCasesProvider) AuthorizeUser(user User) (string, string, error) {

	dbUser, getUserErr := a.userStorage.GetUser(user.Login)

	if getUserErr == ErrUserNotFound {
		return "", "", getUserErr
	}

	if dbUser.Password != user.Password {
		return "", "", ErrInvalidAuthData
	}

	if getUserErr != nil {
		//todo: обработать неожиданные ошибки
	}

	accessToken, createAccessErr := a.tokensProvider.CreateAccessToken(dbUser.Login)
	if createAccessErr != nil {
		//todo: обработать неожиданные ошибки
	}

	refreshToken, refreshInfo := a.tokensProvider.CreateRefreshToken(accessToken)
	updateErr := a.userStorage.UpdateUser(dbUser.Login, User{
		Login:        dbUser.Login,
		Password:     dbUser.Password,
		RefreshInfos: append(dbUser.RefreshInfos, refreshInfo),
	})

	if updateErr != nil {
		//todo: обработать неожиданные ошибки
	}

	return accessToken, refreshToken, nil
}
