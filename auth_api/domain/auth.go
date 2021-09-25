package domain

// Article ...
type Article struct {
	ID       int64  `json:"id"`
	User     string `json:"title" validate:"required"`
	Password string `json:"content" validate:"required"`
}
