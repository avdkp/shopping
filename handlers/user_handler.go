package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/shopping/domain"
	"example.com/shopping/services"
)

type userHandler struct {
	userSvc services.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(svc services.UserService) UserHandler {
	return &userHandler{
		userSvc: svc,
	}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uh.userSvc.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
