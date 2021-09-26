package repository

import (
	"errors"
	"fmt"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	Conn *gorm.DB
}

func NewMysqlAuthRepository(Conn *gorm.DB) domain.AuthRepository {
	return &mysqlAuthRepository{Conn}
}

func (AuthRepo *mysqlAuthRepository) Search(user string) (res domain.Auth, err error) {
	var auth domain.Auth

	if result := AuthRepo.Conn.First(&auth, "Username = ?", user); result.Error != nil {
		fmt.Println(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return auth, domain.ErrNotFound
		}
		return
	}

	return auth, err
}
