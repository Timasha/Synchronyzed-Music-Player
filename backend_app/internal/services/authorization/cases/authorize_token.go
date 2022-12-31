package cases

func (a *AuthCasesProvider) AuthorizeToken(login string, accessToken string) error {
	return a.tokensProvider.ValidAccessToken(login, accessToken)
}
