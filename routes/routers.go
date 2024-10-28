package routes

import (
	"net/http"

	"github.com/SyydMR/To-Do-List/handlers"
	"github.com/SyydMR/To-Do-List/middleware"
	"github.com/gorilla/mux"
)

var RegisterToDoListRoutes = func(router *mux.Router) {
	router.Handle("/items", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetAllItems))).Methods("GET")
	router.Handle("/items", middleware.AuthMiddleware(http.HandlerFunc(handlers.AddItem))).Methods("POST")
	router.Handle("/items/{ItemId}", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateItem))).Methods("PUT")
	router.Handle("/items/{ItemId}", middleware.AuthMiddleware(http.HandlerFunc(handlers.RemoveItem))).Methods("DELETE")
	router.Handle("/items/{ItemId}", middleware.AuthMiddleware(http.HandlerFunc(handlers.CheckItem))).Methods("POST")

	
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{userId}", handlers.GetUserByIDHandler).Methods("GET")

}