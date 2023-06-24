package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/shopping/domain"
	"example.com/shopping/services"
)

type userHandler struct {
	UsrService services.UserService
}

func (uh *userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	// Decode the request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = uh.UsrService.Create(user)
	// Add the user to the users slice
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
