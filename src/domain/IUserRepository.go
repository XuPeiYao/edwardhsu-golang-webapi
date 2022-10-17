package domain

type IUserRepository interface {
	GetUser(id string) *User
}
