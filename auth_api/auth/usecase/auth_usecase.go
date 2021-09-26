package usecase

import (
	"os"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepo domain.AuthRepository
}

func UsecaseInterface(a domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{a}
}

func GenerateToken(user string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user
	atClaims["exp"] = time.Now().Add(time.Hour * 196).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return token, err
}

func (AuthUC *AuthUsecase) Login(user string, password string) (is_authenticated bool, token string) {

	res, err := AuthUC.AuthRepo.Search(user)

	if err != nil {
		return false, ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))

	if err != nil {
		return false, ""
	}

	token, _ = GenerateToken(user)

	return true, token

}
