package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/models"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/repository"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/domain"

	// _ "github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/docs"

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

	ar := repository.CitizenRepositoryInterface(db)

	fmt.Println("TODOS CITIZEN")
	fmt.Println(ar.GetAllCitizen())
	fmt.Println("CRIOU CITIZEN")
	fmt.Println(ar.CreateCitizen(domain.Citizen{Name: "sdsdsds", CPF: "12", Birthdate: time.Now(), Debts: []domain.Debt{{DebtorID: 1}}}))
	// // fmt.Println(ar.CreateCitizen(domain.Citizen{Name: "sdsdsds", CPF: "333", Birthdate: time.Now(), Debts: []domain.Debt{{DebtorID: 1}}}))
	// fmt.Println("TODOS CITIZEN")
	// fmt.Println(ar.GetAllCitizen())

	// fmt.Println("NOVO DEBTO")
	// fmt.Println(ar.InsertNewDebt(domain.Debt{DebtorID: 1, Value: 2000}, 1))
	// fmt.Println("TODOS DEBTOS DO CIDADAO 1")
	// fmt.Println(ar.GetDebtsByCitizenId(1))

	router := mux.NewRouter()
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}
