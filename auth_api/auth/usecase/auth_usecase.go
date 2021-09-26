package usecase

import (
	"fmt"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepo domain.AuthRepository
}

func UsecaseInterface(a domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{a}
}

func (AuthUC *AuthUsecase) Authenticate(user string, password string) (is_authenticated bool) {

	res, err := AuthUC.AuthRepo.Search(user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return false
	}
	fmt.Println(string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(res.Password))

	fmt.Println(err)

	return err != nil

}
