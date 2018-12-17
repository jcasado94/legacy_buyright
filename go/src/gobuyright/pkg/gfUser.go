package root

type GfUser struct {
	ID       string `json:"id"`
	Username string
}

type GfUserService interface {
	CreateUser(u *GfUser) error
	GetById(id string) (*GfUser, error)
}
