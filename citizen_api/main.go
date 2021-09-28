package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/repository"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/domain"

	// _ "github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	time.Sleep(4 * time.Second)

	dsn := "host=postgresql user=test password=test dbname=citzens_financial_data port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
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

	ar := repository.CitizenRepositoryInterface(db)

	fmt.Println("aa")
	// fmt.Println(ar.GetCitizenByCPF("11122233340"))
	// fmt.Println(ar.GetDebtsByCitizenId(1))
	// fmt.Println(ar.GetAddressByCitizenId(1))
	fmt.Println(ar.CreateCitizen(domain.Citizen{Name: "Andreaa", CPF: "99988877766", Birthdate: time.Now(), Debts: []domain.Debt{{DebtorID: 1}}}))
	fmt.Println(ar.GetAllCitizen())

	router := mux.NewRouter()
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}
