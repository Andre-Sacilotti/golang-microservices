package mocks

import (
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	mock "github.com/stretchr/testify/mock"
)

type AuthRepository struct {
	mock.Mock
}

func (AuthRepo *AuthRepository) Search(user string) (res domain.Auth, err error) {
	ret := AuthRepo.Mock.Called(user)

	var r0 domain.Auth
	if rf, ok := ret.Get(0).(func(string) domain.Auth); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.Auth)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
