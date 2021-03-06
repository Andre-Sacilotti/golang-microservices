package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/delivery"
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/repository"
	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/usecase"
	_ "github.com/Andre-Sacilotti/golang-credit-backend/auth_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

// @title Auth API
// @version 1.0
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @description REST API to login and authenticate a token
// @contact.name André Sacilotti
// @contact.email andre.sacilotti@gmail.com
// @host 0.0.0.0:81
// @BasePath /auth
func main() {

	time.Sleep(4 * time.Second)

	dsn := "host=postgresql_auth user=test password=test dbname=auth_users port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := mux.NewRouter()

	ar := repository.AuthRepositoryInterface(db)
	au := usecase.UsecaseInterface(ar)
	delivery.AuthHandlerInterface(router, au)
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":81", router))

}
