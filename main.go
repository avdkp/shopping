package main

import (
	"example.com/shopping/handlers"
	"example.com/shopping/middlewares"
	"example.com/shopping/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	userService := services.GetUserService()
	authService := services.GetAuthService()
	authMiddleware := middlewares.NewAuthMiddleware(authService)
	userHandler := handlers.NewUserHandler(userService)
	inventoryHandler := handlers.NewInventoryHandler(services.NewInventoryService())

	publicRouter := router.PathPrefix("/public").Subrouter()
	publicRouter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	publicRouter.HandleFunc("/login", userHandler.Login).Methods("POST")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(authMiddleware.AuthorizeAdmin)
	adminRouter.HandleFunc("/add-items", inventoryHandler.AddItems).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
