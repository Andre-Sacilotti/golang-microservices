package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/auth/repository"
	_ "github.com/Andre-Sacilotti/golang-credit-backend/auth_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

// @title Auth API
// @version 1.0
// @description REST API to authenticate bearer token
// @contact.name André Sacilotti
// @contact.email andre.sacilotti@gmail.com
// @host localhost:8080
// @BasePath /auth
func main() {

	time.Sleep(2 * time.Second)

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

	ar := repository.NewMysqlAuthRepository(db)
	fmt.Println(ar.Search("andrae"))
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":80", router))

}
