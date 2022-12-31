package cases

func (a *AuthCasesProvider) CreateUser(user User) error {
	err := a.userStorage.CreateUser(user)
	if err == ErrUserAlreadyExist {
		return err
	}
	if err != nil {
		// todo: добавить обработку неожиданных ошибок
	}
	return nil
}
