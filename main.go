package main

import (
	"example.com/shopping/handlers"
	"example.com/shopping/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	userHandler := handlers.NewUserHandler(services.NewUserService())
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", userHandler.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
