package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/Andre-Sacilotti/golang-credit-backend/auth_api/domain"
	"github.com/gorilla/mux"
)

type Response struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthResponse struct {
	Message string `json:"message"`
}

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
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

func AuthHandlerInterface(router *mux.Router, au domain.AuthUsecase) {
	handler := &AuthHandler{au}

	router.HandleFunc("/auth/login", handler.Login).Methods("POST")
	router.HandleFunc("/auth/authenticate", handler.Authenticate).Methods("GET")
}

// Login godoc
// @Summary Login and get a authentication token
// @Description Get an JWT authentication token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param credentials body domain.Auth true "Login credentials must have an username and a password"
// @Success 200 Response
// @Success 200 ErrorResponse
// @Router /login [post]
func (AuthHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials domain.Auth

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	isAuthenticated, token := AuthHandler.AuthUsecase.Login(credentials.Username, credentials.Password)

	if isAuthenticated {
		respondWithJSON(w, http.StatusOK, Response{token, "bearer"})
		return
	}

	respondWithError(w, http.StatusUnauthorized, "Not authorized")

}

// Authentcation godoc
// @Summary Authenticate a token
// @Description Authenticate a token
// @Tags Auth
// @Param Authorization header string true "Insert your access token"
// @Accept  json
// @Produce  json
// @Success 200 AuthResponse
// @Success 401 ErrorResponse
// @Router /authenticate [get]
func (AuthHandler *AuthHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var header = r.Header.Get("Authorization")
	isValid := AuthHandler.AuthUsecase.Authenticate(header)

	if isValid {
		respondWithJSON(w, http.StatusOK, AuthResponse{"Authorized"})
		return
	}
	respondWithError(w, http.StatusUnauthorized, "Not authorized")
}
