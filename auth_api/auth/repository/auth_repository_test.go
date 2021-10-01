package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/repository"
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
)

func TestSearch(test *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		test.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockAuths := []domain.Auth{
		domain.Auth{Username: "carlos", Password: "123"},
		domain.Auth{Username: "rebeca", Password: "5454"},
		domain.Auth{Username: "alberto.silva", Password: "4545sss"},
		domain.Auth{Username: "luiza", Password: "ggg@@@!#$"},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).
		AddRow(0, mockAuths[0].Username, mockAuths[0].Password).
		AddRow(1, mockAuths[1].Username, mockAuths[1].Password).
		AddRow(2, mockAuths[2].Username, mockAuths[2].Password).
		AddRow(3, mockAuths[3].Username, mockAuths[3].Password)

	query := "SELECT \\* FROM \"auths\" WHERE Username = \\$1 ORDER BY \"auths\".\"username\" LIMIT 1"

	mock.ExpectQuery(query).WillReturnRows(rows)

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	a := repository.AuthRepositoryInterface(gormDB)

	list, err := a.Search("carlos")

	assert.NoError(test, err)
	assert.NotEmpty(test, list)

}
