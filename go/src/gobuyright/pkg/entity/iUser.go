package entity

type IUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserService interface {
	CreateUser(u *IUser) error
	GetByUsername(username string) (*IUser, error)
}
