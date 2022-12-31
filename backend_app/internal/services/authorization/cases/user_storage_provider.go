package cases

import "errors"

type UserStorage interface {
	CreateUser(user User) error

	// authID is any data for auth, like login, email or phone number(if they will added)
	GetUser(authID string) (User, error)

	// oldUserAuthID is any data for auth, like login, email or phone number(if they will added)
	UpdateUser(oldUserAuthId string, newUser User) error

	DeleteUser(user User) error
}
type RefreshInfo struct {
	Body       string
	AccessPart string
}
type User struct {
	Login        string
	Password     string
	RefreshInfos []RefreshInfo
}

var (
	ErrUserNotFound     error = errors.New("user not found")
	ErrUserAlreadyExist error = errors.New("user already exists")
	ErrInvalidAuthData  error = errors.New("invalid login data")
)
