package mocks

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	mock "github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	mock.Mock
}

func (AuthUC *AuthUsecase) Authenticate(tokenstr string) (is_valid bool) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
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

func (AuthUC *AuthUsecase) Login(user string, password string) (is_authenticated bool, token string) {
	if user == "admin" {
		TruePass := "$2a$15$UZKMN5zBXk.DOuzi6RtmvOBUElwtD.1ztXyKTr6mufUuaTHCnDro6"
		err := bcrypt.CompareHashAndPassword([]byte(TruePass), []byte(password))
		if err != nil {
			return false, ""
		}
	} else if user == "admin2" {
		TruePass := "$2a$15$UZKMN5zBXk.DOuzi6RtmvOBUElwtD.1ztXyKTr6mufUuaTHCnDro6"
		err := bcrypt.CompareHashAndPassword([]byte(TruePass), []byte(password))
		if err != nil {
			return false, ""
		}
	} else {
		return false, ""
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user
	atClaims["exp"] = time.Now().Add(time.Hour * 196).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ = at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return true, token
}
