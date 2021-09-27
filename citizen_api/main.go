package main

import (
	"log"
	"net/http"

	_ "github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}
