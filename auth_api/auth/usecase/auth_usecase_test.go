package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/usecase"
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain/mocks"
)

func TestLogin(test *testing.T) {
	mockAuthRepo := new(mocks.AuthRepository)

	MockAuth := domain.Auth{
		Username: "admin", Password: "1234",
	}
	MockAuthWrongpass := domain.Auth{
		Username: "admin", Password: "sdfsdfsdfsdf",
	}
	MockAuthNotExists := domain.Auth{
		Username: "carlos", Password: "1234",
	}

	test.Run("user-exist-right-pass", func(test *testing.T) {
		tmp := MockAuth

		mockAuthRepo.On(
			"Search", mock.Anything,
			mock.AnythingOfType("string"),
		).Return(domain.Auth{
			Username: "admin", Password: "$2a$15$UZKMN5zBXk.DOuzi6RtmvOBUElwtD.1ztXyKTr6mufUuaTHCnDro6",
		}, nil).Once()

		u := usecase.UsecaseInterface(mockAuthRepo)

		login, token := u.Login(tmp.Username, tmp.Password)
		fmt.Println(login)
		assert.Equal(test, login, true)
		assert.NotEqual(test, token, "")

	})

	test.Run("user-exist-wrong-pass", func(test *testing.T) {
		tmp := MockAuthWrongpass

		mockAuthRepo.On(
			"Search", mock.Anything,
			mock.AnythingOfType("string"),
		).Return(domain.Auth{
			Username: "admin", Password: "$2a$15$UZKMN5zBXk.DOuzi6RtmvOBUElwtD.1ztXyKTr6mufUuaTHCnDro6",
		}, nil).Once()

		u := usecase.UsecaseInterface(mockAuthRepo)

		login, token := u.Login(tmp.Username, tmp.Password)

		assert.Equal(test, token, "")
		assert.Equal(test, login, false)
	})

	test.Run("user-not-exist", func(test *testing.T) {
		tmp := MockAuthNotExists

		mockAuthRepo.On(
			"Search", mock.Anything,
			mock.AnythingOfType("string"),
		).Return(domain.Auth{
			Username: "admin", Password: "$2a$15$UZKMN5zBXk.DOuzi6RtmvOBUElwtD.1ztXyKTr6mufUuaTHCnDro6",
		}, gorm.ErrRecordNotFound).Once()

		u := usecase.UsecaseInterface(mockAuthRepo)

		login, token := u.Login(tmp.Username, tmp.Password)

		assert.Equal(test, token, "")
		assert.Equal(test, login, false)
	})

}
