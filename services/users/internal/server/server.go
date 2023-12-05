package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"scooter/users/internal/api"
)

func RegisterRoutes(router *mux.Router, port int) {

	router.HandleFunc("/user", api.GetUsers).Methods("GET")
	router.HandleFunc("/user/{cpf}", api.GetUser).Methods("GET")
	router.HandleFunc("/user/{cpf}", api.CreateUser).Methods("POST")
	router.HandleFunc("/user/{cpf}", api.UpdateUser).Methods("PATCH")
	router.HandleFunc("/user/{cpf}", api.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
