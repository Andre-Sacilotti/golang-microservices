package main

import (
	"log"
	"net/http"

	_ "github.com/Andre-Sacilotti/golang-credit-backend/auth_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"

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

	router := mux.NewRouter()

	router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))

}
