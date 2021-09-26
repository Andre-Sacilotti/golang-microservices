package domain

type Auth struct {
	ID       int64  `json:"id"`
	Username string `json:"title" validate:"required"`
	Password string `json:"content" validate:"required"`
}

type AuthUsecase interface {
	Authenticate(User string, password string) bool
}

type AuthRepository interface {
	Search(User string) (Auth, error)
}
