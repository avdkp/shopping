package main

import (
	"example.com/shopping/domain"
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
	inventoryService := services.NewInventoryService()
	cartService := services.NewCartService(inventoryService)

	authMiddleware := middlewares.NewAuthMiddleware(authService)
	userHandler := handlers.NewUserHandler(userService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	cartHandler := handlers.NewCartHandler(cartService)

	publicRouter := router.PathPrefix("/public").Subrouter()
	publicRouter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	publicRouter.HandleFunc("/login", userHandler.Login).Methods("POST")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(authMiddleware.AuthorizationMiddlewareCreator(domain.AdminRole))
	adminRouter.HandleFunc("/add-items", inventoryHandler.AddItems).Methods("POST")
	adminRouter.HandleFunc("/suspend/user/{user-id}", userHandler.SuspendUser).Methods("PATCH")

	userRouter := router.PathPrefix("").Subrouter()
	userRouter.Use(authMiddleware.AuthorizationMiddlewareCreator(domain.UserRole))
	userRouter.HandleFunc("/all-items", inventoryHandler.GetItems).Methods("GET")
	userRouter.HandleFunc("/add-to-cart", cartHandler.AddToCart).Methods("PATCH")
	userRouter.HandleFunc("/remove-from-cart", cartHandler.RemoveFromCart).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", router))
}
