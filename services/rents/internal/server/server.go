package server

import (
	"github.com/gorilla/mux"

	"rents/internal/api"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/rent/{scooter}", api.StartScooterRent).Methods("POST")
	router.HandleFunc("/finish-rent/{scooter}", api.FinishScooterRent).Methods("POST")
}
