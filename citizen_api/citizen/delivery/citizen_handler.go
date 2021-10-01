package delivery

import (
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/citizen/{CPF}", handler.GetCitizen).Methods("GET")
	router.HandleFunc("/citizen/{CPF}", handler.UpdateCitizen).Methods("PUT")

	router.HandleFunc("/citizen/{CPF}/debts", handler.GetCitizenDebts).Methods("GET")
	router.HandleFunc("/citizen/{CPF}/debts/{DebtID}", handler.UpdateCitizenDebt).Methods("PUT")
	router.HandleFunc("/citizen/{CPF}/debts", handler.InsertCitizenDebt).Methods("POST")
	router.HandleFunc("/citizen/{CPF}/debts/{DebtID}", handler.DeleteCitizenDebt).Methods("DELETE")

	router.HandleFunc("/citizen/{CPF}/address", handler.GetCitizenAddresses).Methods("GET")
	router.HandleFunc("/citizen/{CPF}/address/{AddressID}", handler.DeleteCitizenAddress).Methods("DELETE")
	router.HandleFunc("/citizen/{CPF}/address/{AddressID}", handler.UpdateCitizenAddress).Methods("PUT")
	router.HandleFunc("/citizen/{CPF}/address", handler.InsertCitizenAddress).Methods("POST")
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
	fmt.Println("INSERT CITI")
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
	Limit := 10
	if r.FormValue("offset") != "" {
		Offset, err := strconv.Atoi(r.FormValue("offset"))
		if err != nil {
			Offset = 1
		}
		log.Println(Offset)
	}

	if r.FormValue("limit") != "" {
		Limit, err := strconv.Atoi(r.FormValue("limit"))
		if err != nil {
			Limit = 1
		}
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

// Get a specific citizen by CPF godoc
// @Summary Get a citizen data
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "Citizen CPF to return "
// @Success 200 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF} [get]
func (CitizenHandler *CitizenHandler) GetCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	res := CitizenHandler.CitizenUsecase.GetCitizenByCPF(params["CPF"])

	respondWithJSON(w, http.StatusOK, Response{"Get specific citizens", res})

}

// Update a specific citizen data godoc
// @Summary Update citizen data
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "Citizen CPF to update"
// @Success 200 {object} Response
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF} [put]
func (CitizenHandler *CitizenHandler) UpdateCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var Citizen domain.Citizen
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Citizen); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	res := CitizenHandler.CitizenUsecase.UpdateCitizenByCPF(Citizen, params["CPF"])

	respondWithJSON(w, http.StatusOK, Response{"Citizen updated", res})

}

// Get a list of debts from a citizen godoc
// @Summary Get a list of debts
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Success 200 {object} ResponseDebt
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/debts [get]
func (CitizenHandler *CitizenHandler) GetCitizenDebts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	res := CitizenHandler.CitizenUsecase.GetDebtsByCitizenCPF(params["CPF"])

	respondWithJSON(w, http.StatusOK, ResponseDebt{"Debts selected", res})

}

// Update a specific debt from citizen
// @Summary Update a specific debt
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param ID path string false "Id of the debt"
// @Param debt_data body domain.Debt true "Debt data"
// @Success 200 {object} ResponseDebt
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/debts/{ID} [put]
func (CitizenHandler *CitizenHandler) UpdateCitizenDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var Debt domain.Debt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Debt); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	ID, err := strconv.Atoi(params["DebtID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID must be a integer")
		return
	}

	res, err := CitizenHandler.CitizenUsecase.UpdateDebt(Debt, ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cant updated the Debt")
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseDebt{"Debt Updated", res})

}

// Delete a specific debt from citizen
// @Summary Delete a specific debt
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param ID path string false "Id of the debt"
// @Success 200 {object} ResponseDebt
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/debts/{ID} [delete]
func (CitizenHandler *CitizenHandler) DeleteCitizenDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["DebtID"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "ID must be a integer")
		return
	}

	res := CitizenHandler.CitizenUsecase.DeleteDebt(ID)

	respondWithJSON(w, http.StatusOK, ResponseDebt{"Debt Deleted", res})

}

// Insert a new debt from citizen
// @Summary Insert a new debt
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param debt_data body domain.Debt true "Debt data"
// @Success 200 {object} ResponseDebt
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/debts [post]
func (CitizenHandler *CitizenHandler) InsertCitizenDebt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var Debt domain.Debt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Debt); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	res, err := CitizenHandler.CitizenUsecase.InsertNewDebt(Debt, params["CPF"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cant insert the Debt")
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseDebt{"Debt Inserted", res})

}

// Get a list of address from a citizen godoc
// @Summary Get a list of address
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the address"
// @Success 200 {object} ResponseAddress
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/address [get]
func (CitizenHandler *CitizenHandler) GetCitizenAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	res := CitizenHandler.CitizenUsecase.GetAddressByCitizenCPF(params["CPF"])

	respondWithJSON(w, http.StatusOK, ResponseAddress{"Address selected", res})

}

// Update a specific Address from citizen
// @Summary Update a specific Address
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param ID path string false "Id of the Address"
// @Param address_data body domain.Address true "Address data"
// @Success 200 {object} ResponseAddress
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/address/{ID} [put]
func (CitizenHandler *CitizenHandler) UpdateCitizenAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var Address domain.Address
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Address); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	ID, err := strconv.Atoi(params["AddressID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID must be a integer")
		return
	}

	res, err := CitizenHandler.CitizenUsecase.UpdateAddress(Address, ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cant updated the Address")
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseAddress{"Address Updated", res})

}

// Delete a specific Address from citizen
// @Summary Delete a specific Address
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param ID path string false "Id of the Address"
// @Success 200 {object} ResponseAddress
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/address/{ID} [delete]
func (CitizenHandler *CitizenHandler) DeleteCitizenAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["AddressID"])
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "ID must be a integer")
		return
	}

	res := CitizenHandler.CitizenUsecase.DeleteAddress(ID)

	respondWithJSON(w, http.StatusOK, ResponseAddress{"Address Deleted", res})

}

// Insert a new Address to citizen
// @Summary Insert a new Address
// @Tags Citizen
// @Accept  json
// @Produce  json
// @Param CPF path string false "CPF of the citizen"
// @Param address_data body domain.Address true "Address data"
// @Success 200 {object} ResponseAddress
// @Failure 401 {object} ErrorResponse
// @Router /citizen/{CPF}/address [post]
func (CitizenHandler *CitizenHandler) InsertCitizenAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var Address domain.Address
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Address); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	res, err := CitizenHandler.CitizenUsecase.InsertNewAddress(Address, params["CPF"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cant insert the Address")
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseAddress{"Address Inserted", res})

}
