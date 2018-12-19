package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
