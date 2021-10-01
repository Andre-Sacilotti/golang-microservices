package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepo domain.AuthRepository
}

func UsecaseInterface(a domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{a}
}

func generateToken(user string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user
	atClaims["exp"] = time.Now().Add(time.Hour * 196).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return token, err
}

func decodeToken(tokenSTR string) bool {
	token, err := jwt.Parse(tokenSTR, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

func (AuthUC *AuthUsecase) Authenticate(tokenstr string) (is_valid bool) {
	return decodeToken(tokenstr)
}

func (AuthUC *AuthUsecase) Login(user string, password string) (is_authenticated bool, token string) {

	res, err := AuthUC.AuthRepo.Search(user)
	fmt.Println(res)
	if err != nil {
		return false, ""
	}
	fmt.Println(user)
	fmt.Println(res)
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))

	if err != nil {
		return false, ""
	}

	fmt.Println(err)
	fmt.Println("user")
	token, _ = generateToken(user)

	return true, token

}
