package cases

type AuthCasesProvider struct {
	userStorage    UserStorage
	tokensProvider TokensProvider
}

func (a *AuthCasesProvider) New(userStorage UserStorage, tokensProvider TokensProvider) {
	a.userStorage = userStorage
	a.tokensProvider = tokensProvider
}
