package entity

type GfUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type GfUserService interface {
	CreateUser(u *GfUser) error
	GetByUsername(username string) (*GfUser, error)
}
