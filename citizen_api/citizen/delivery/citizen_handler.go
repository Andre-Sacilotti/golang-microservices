package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/domain"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string           `json:"message"`
	Data    []domain.Citizen `json:"data"`
}

type ResponseAddress struct {
	Message string           `json:"message"`
	Data    []domain.Address `json:"data"`
}

type ResponseDebt struct {
	Message string        `json:"message"`
	Data    []domain.Debt `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CitizenHandler struct {
	CitizenUsecase domain.CitizenUsecase
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CitizenHandlerInterface(router *mux.Router, au domain.CitizenUsecase) {
	handler := &CitizenHandler{au}

	router.HandleFunc("/citizen", handler.InsertCitizen).Methods("POST")
	router.HandleFunc("/citizen", handler.GetAllCitizen).Queries("offset", "{[0-9]*?}").Queries("limit", "{[0-9]*?}").Methods("GET")
	router.Path("/citizen").HandlerFunc(handler.GetAllCitizen)
	// router.HandleFunc("/citizen/{CPF}", handler.GetCitizen).Methods("GET")
	// router.HandleFunc("/citizen/{CPF}", handler.GetCitizen).Methods("PUT")

	// router.HandleFunc("/citizen/{CPF}/debts", handler.GetCitizenDebts).Methods("GET")
	// router.HandleFunc("/citizen/{CPF}/debts/{DebtID}", handler.UpdateCitizenDebt).Methods("PUT")
	// router.HandleFunc("/citizen/{CPF}/debts", handler.InsertCitizenDebt).Methods("POST")
	// router.HandleFunc("/citizen/{CPF}/debts/{DebtID}", handler.DeleteCitizenDebt).Methods("DELETE")

	// router.HandleFunc("/citizen/{CPF}/address", handler.GetCitizenAddresses).Methods("GET")
	// router.HandleFunc("/citizen/{CPF}/address/{AddressID}", handler.DeleteCitizenAddress).Methods("DELETE")
	// router.HandleFunc("/citizen/{CPF}/address/{AddressID}", handler.UpdateCitizenAddress).Methods("PUT")
	// router.HandleFunc("/citizen/{CPF}/address/{AddressID}", handler.InsertCitizenAddress).Methods("POST")
}

// Insert new citizen godoc
// @Summary Insert a new citizen data
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param citizen_data body domain.Citizen true "Citizen can have many addresses and debts"
// @Success 200 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /citizen [post]
func (CitizenHandler *CitizenHandler) InsertCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Citizen domain.Citizen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Citizen); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	res, err := CitizenHandler.CitizenUsecase.CreateCitizen(Citizen)

	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusConflict, "Citizen with same CPF already exists")
		return
	}

	respondWithJSON(w, http.StatusOK, Response{"Citizen inserted", res})

}

// Get a list of all citizens godoc
// @Summary Get a list of citizens
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param offset query int false "Offset for pagination in request"
// @Param limit query int false "Limit how many citizens will be returned"
// @Success 200 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /citizen [get]
func (CitizenHandler *CitizenHandler) GetAllCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Offset := 0
	Limit := 0
	if r.FormValue("offset") != "" {
		Offset, err := strconv.Atoi(r.FormValue("offset"))
		log.Println(err)
		log.Println(Offset)
	}

	if r.FormValue("limit") == "" {
		Limit, err := strconv.Atoi(r.FormValue("limit"))
		log.Println(err)
		log.Println(Limit)
	}

	if Limit < 1 {
		respondWithError(w, http.StatusBadRequest, "Limit should be greater than equals one")
		return
	}
	if Offset < 0 {
		respondWithError(w, http.StatusBadRequest, "Offset should be greater than equals zero")
		return
	}

	res := CitizenHandler.CitizenUsecase.GetAllCitizen(Offset, Limit)

	respondWithJSON(w, http.StatusOK, Response{"Get all citizens", res})

}
