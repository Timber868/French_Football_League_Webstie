package user

import (
	"net/http"

	"github.com/Timber868/french-league/types"
	"github.com/Timber868/french-league/utils"
	"github.com/gorilla/mux"
)

// Each service has a handler
type Handler struct {
	store *types.UserStore //Where i left off
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	//check if the user exists if not create a new user

	//Type we will use to decode our payload
	var payload types.RegisterUserPayload

	//Write it to json so that it can be sent through our api that there was an error
	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//Check in the database if the user exists
}
