package main

import (
	"log"
	"net/http"

	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/delivery"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/models"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/repository"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/usecase"

	_ "github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=postgresql user=test password=test dbname=citzens_financial_data port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Citizen{})
	db.AutoMigrate(&models.Debt{})
	db.AutoMigrate(&models.Address{})

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
	ar := repository.CitizenRepositoryInterface(db)
	au := usecase.UsecaseInterface(ar)
	delivery.CitizenHandlerInterface(router, au)

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":82", router))
}
