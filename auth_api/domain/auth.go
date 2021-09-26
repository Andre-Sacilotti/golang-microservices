package domain

type Auth struct {
	Username string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthUsecase interface {
	Login(User string, password string) (bool, string)
	Authenticate(token string) bool
}

type AuthRepository interface {
	Search(User string) (Auth, error)
}
