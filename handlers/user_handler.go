package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/shopping/domain"
	"example.com/shopping/services"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type userHandler struct {
	userSvc services.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
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

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := uh.userSvc.Login(user)

	if err == domain.UserNotFoundError {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err == domain.UnAuthorizedError {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		Token: string(token),
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Print("Error in writing the response ", err)
	}
}
