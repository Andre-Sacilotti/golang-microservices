package repository

type Auth struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"index"`
	Password string
}
