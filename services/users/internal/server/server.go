package server

import (
	"github.com/gorilla/mux"

	"scooter/users/internal/api"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/user", api.GetUsers).Methods("GET")
	router.HandleFunc("/user", api.CreateUser).Methods("POST")
	router.HandleFunc("/user/{cpf}", api.GetUser).Methods("GET")
	router.HandleFunc("/user/{cpf}", api.UpdateUser).Methods("PATCH")
	router.HandleFunc("/user/{cpf}", api.DeleteUser).Methods("DELETE")
}
