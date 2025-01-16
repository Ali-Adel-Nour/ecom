package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Ali-Adel-Nour/ecom/cmd/service/auth"
	"github.com/Ali-Adel-Nour/ecom/cmd/types"
	"github.com/Ali-Adel-Nour/ecom/cmd/utils"

)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	//router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

	}
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already %s exists", user.Email))
		return
	}
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
